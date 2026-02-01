## Context

The application currently runs via docker-compose with backend (Go) and frontend (Svelte) services. Traefik reverse proxy is already running on the target Linux VM with an external network named `proxy`, handling SSL termination and routing for existing services. The goal is to add automated deployment from GitHub while integrating with this existing infrastructure.

Current docker-compose setup uses an internal network (`rechtebank-network`) and exposes ports directly (8080, 5173). The target VM uses Traefik for routing with labels-based configuration, as seen in the existing Lychee service setup.

## Goals / Non-Goals

**Goals:**
- Automate deployment on push to main branch
- Integrate application containers with existing Traefik proxy
- Manage sensitive environment variables (GEMINI_API_KEY) securely via GitHub Secrets
- Enable SSL/TLS via Traefik's Let's Encrypt integration
- Minimize downtime during deployments

**Non-Goals:**
- Setting up or configuring Traefik itself (already running)
- Multi-environment deployments (staging, production) - main branch only
- Database migrations or data persistence strategies
- Monitoring or logging infrastructure

## Decisions

### 1. GitHub Actions with SSH Action
**Decision**: Use `appleboy/ssh-action` to deploy via SSH  
**Rationale**: Direct SSH deployment keeps the architecture simple and doesn't require self-hosted runners on the VM. The VM becomes the single source of truth for the running application.  
**Alternatives Considered**: Self-hosted GitHub runner would enable Docker commands locally but adds complexity and maintenance overhead.

### 2. Environment Variables Strategy
**Decision**: Write secrets to `.env` file during deployment, not committed to repo  
**Rationale**: GitHub Secrets → SSH command → `.env` file on VM keeps secrets out of version control while making them available to docker-compose. The `.env` file should be listed in `.gitignore`.  
**Alternatives Considered**: Docker secrets or external secret management adds complexity for a single-VM deployment.

### 3. Traefik Network Integration
**Decision**: Remove exposed ports (8080, 5173) and use Traefik's proxy network exclusively for frontend  
**Rationale**: Backend doesn't need direct external access - Traefik routes to frontend, frontend proxies API calls to backend via internal network. This follows the same pattern as the existing Lychee service.  
**Network Strategy**:
- Frontend: joins `proxy` network (external) for Traefik routing
- Backend: joins `rechtebank-network` (internal) only
- Frontend also joins `rechtebank-network` to communicate with backend

### 4. Traefik Labels Configuration
**Decision**: Apply Traefik labels only to frontend service  
**Rationale**: Frontend serves as the public entry point. API calls are proxied through the frontend's nginx config or client-side requests to `/api` paths.  
**Labels Include**:
- `traefik.enable=true`
- Router rule with hostname (e.g., `rechtbank.example.com`)
- TLS with Let's Encrypt certresolver
- Network specification (`traefik.docker.network=proxy`)

### 5. Docker Compose Build Strategy
**Decision**: Use `docker compose up -d --build` on deployment  
**Rationale**: Ensures fresh builds from latest code. The `--build` flag rebuilds images even if they exist, and `-d` runs detached. Docker layer caching minimizes rebuild time for unchanged layers.  
**Alternatives Considered**: Pre-building images in GitHub Actions and pushing to registry adds complexity for a single-VM setup.

### 6. Deployment Working Directory
**Decision**: Clone/deploy to a consistent path on VM (e.g., `/opt/rechtbank` or `/home/user/rechtbank`)  
**Rationale**: The SSH action's `script` section needs to `cd` to the correct directory. Using a consistent path makes the workflow reproducible.

## Risks / Trade-offs

**Risk: SSH Private Key Security**  
→ Mitigation: Store SSH private key in GitHub Secrets (encrypted at rest). Use key-based auth only, disable password auth on VM. Consider using a deploy-specific key with limited permissions.

**Risk: Downtime During Deployment**  
→ Mitigation: `docker compose up -d` performs rolling updates when possible. Consider adding health checks to both services to ensure they're ready before Traefik routes traffic.

**Risk: `.env` File Overwrites**  
→ Mitigation: The deployment script creates `.env` on each run. Ensure all required environment variables are defined in GitHub Secrets. Consider logging which variables are set (not their values) for debugging.

**Trade-off: No Rollback Mechanism**  
→ The workflow does `git pull` which always advances forward. Rolling back requires manual `git reset` or re-running an older commit. Consider tagging releases for easier rollback.

**Risk: Port Conflicts After Traefik Migration**  
→ Mitigation: Remove port mappings from docker-compose.yml to avoid conflicts. Traefik handles all external routing.

**Trade-off: Build Time on VM**  
→ Building on the VM means slower deployments compared to pre-built images. However, for a small app, this is acceptable and simplifies the pipeline.

## Migration Plan

### Pre-deployment Checklist
1. Add GitHub Secrets: `VM_IP`, `VM_USER`, `SSH_PRIVATE_KEY`, `GEMINI_API_KEY`
2. Verify Traefik is running on VM with network named `proxy`
3. Decide on deployment path (e.g., `/opt/rechtbank`)
4. Ensure `.env` is in `.gitignore`
5. Choose domain name for Traefik routing (e.g., `rechtbank.example.com`)

### Deployment Steps
1. Create `.github/workflows/deploy.yml` with SSH action
2. Update `docker-compose.yml`:
   - Remove port mappings
   - Add Traefik labels to frontend
   - Configure networks (proxy + rechtebank-network)
3. Update frontend's PUBLIC_API_URL build arg to use production domain
4. Test workflow on a feature branch first (optional: use a different branch filter)
5. Merge to main to trigger first deployment

### Rollback Strategy
If deployment fails:
1. SSH to VM manually
2. `cd /opt/rechtbank`
3. `git reset --hard <previous-commit-sha>`
4. `docker compose up -d --build`

Alternatively, push a revert commit to main to trigger automatic re-deployment.

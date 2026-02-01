## 1. Pre-Deployment Setup

- [x] 1.1 Add `.env` to `.gitignore` if not already present
- [x] 1.2 Configure GitHub Secrets: VM_IP, VM_USER, SSH_PRIVATE_KEY, GEMINI_API_KEY
- [x] 1.3 Verify Traefik is running on VM with network named `proxy`
- [x] 1.4 Determine deployment path on VM (e.g., `/opt/rechtbank`)
- [x] 1.5 Clone repository to deployment path on VM if not already present
- [x] 1.6 Decide on production domain name (e.g., `rechtbank.example.com`)

## 2. GitHub Actions Workflow

- [x] 2.1 Create `.github/workflows/` directory if not exists
- [x] 2.2 Create `.github/workflows/deploy.yml` workflow file
- [x] 2.3 Configure workflow trigger on push to main branch
- [x] 2.4 Add SSH action step with appleboy/ssh-action@v1.0.3
- [x] 2.5 Configure SSH action with secrets for host, username, and key
- [x] 2.6 Add script to change to deployment directory
- [x] 2.7 Add script to execute `git pull origin main`
- [x] 2.8 Add script to create `.env` file with GEMINI_API_KEY from secrets
- [x] 2.9 Add script to execute `docker compose up -d --build`
- [x] 2.10 Test workflow by pushing to main branch

## 3. Docker Compose Configuration

- [x] 3.1 Remove port mapping `8080:8080` from backend service
- [x] 3.2 Remove port mapping `5173:80` from frontend service
- [x] 3.3 Add `proxy` network to frontend service networks list
- [x] 3.4 Keep `rechtebank-network` on both frontend and backend services
- [x] 3.5 Add `traefik.enable=true` label to frontend service
- [x] 3.6 Add Traefik router rule label with production hostname
- [x] 3.7 Add `traefik.http.routers.<name>.tls=true` label to frontend
- [x] 3.8 Add `traefik.http.routers.<name>.tls.certresolver=lets-encrypt` label
- [x] 3.9 Add `traefik.docker.network=proxy` label to frontend
- [x] 3.10 Define external `proxy` network in networks section
- [x] 3.11 Verify `rechtebank-network` remains as internal network

## 4. Environment Configuration

- [x] 4.1 Update frontend build arg PUBLIC_API_URL for production domain
- [x] 4.2 Ensure backend CORS_ORIGIN accepts production domain
- [x] 4.3 Set backend ENV to `production` via .env file
- [x] 4.4 Verify all required environment variables are documented

## 5. Testing and Validation

- [x] 5.1 Test docker-compose configuration locally with `docker compose config`
- [x] 5.2 Verify workflow syntax with GitHub Actions validator
- [x] 5.3 Perform test deployment to VM
- [x] 5.4 Verify frontend is accessible via production domain with HTTPS
- [x] 5.5 Verify Let's Encrypt certificate was issued successfully
- [x] 5.6 Test API calls from frontend to backend work correctly
- [x] 5.7 Verify backend is not directly accessible from external network
- [x] 5.8 Check deployment logs in GitHub Actions for errors
- [x] 5.9 Test rollback procedure by reverting a commit

## 6. Documentation

- [x] 6.1 Update README with deployment instructions
- [x] 6.2 Document required GitHub Secrets
- [x] 6.3 Document VM setup requirements
- [x] 6.4 Document rollback procedure

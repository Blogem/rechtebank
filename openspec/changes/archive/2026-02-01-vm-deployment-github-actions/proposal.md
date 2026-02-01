## Why

The application needs to be deployed to a production Linux VM environment. Currently there's no automated deployment process, requiring manual deployment steps. This change establishes a CI/CD pipeline using GitHub Actions and integrates with existing Traefik reverse proxy infrastructure for SSL termination and routing.

## What Changes

- Create GitHub Actions workflow for automated deployment via SSH
- Configure docker-compose.yml to integrate with existing Traefik reverse proxy
- Set up environment variable management through GitHub Secrets
- Enable automatic deployment on push to main branch
- Configure proper networking between application containers and Traefik proxy network

## Capabilities

### New Capabilities
- `github-actions-ssh-deployment`: Automated deployment workflow that connects to VM via SSH, pulls latest code, and rebuilds containers
- `traefik-docker-integration`: Docker Compose configuration with Traefik labels for routing, SSL certificates, and network connectivity

### Modified Capabilities

## Impact

- **Infrastructure**: Requires GitHub Secrets for VM credentials (VM_IP, VM_USER, SSH_PRIVATE_KEY, EXTERNAL_API_KEY)
- **Docker Compose**: Updated to use external Traefik proxy network instead of standalone networking
- **CI/CD**: New GitHub Actions workflow added to `.github/workflows/` directory
- **Deployment**: Automated deployment replaces manual deployment process
- **Dependencies**: Assumes Traefik reverse proxy is already running on target VM with network named `proxy`

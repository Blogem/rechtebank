## ADDED Requirements

### Requirement: Workflow triggers on main branch push
The workflow SHALL trigger automatically when code is pushed to the main branch.

#### Scenario: Code pushed to main
- **WHEN** a commit is pushed to the main branch
- **THEN** the deployment workflow starts execution

#### Scenario: Code pushed to other branch
- **WHEN** a commit is pushed to a non-main branch
- **THEN** the deployment workflow does not trigger

### Requirement: SSH connection to target VM
The workflow SHALL establish an SSH connection to the target VM using credentials from GitHub Secrets.

#### Scenario: Successful SSH authentication
- **WHEN** the workflow attempts to connect to the VM
- **THEN** authentication succeeds using VM_IP, VM_USER, and SSH_PRIVATE_KEY secrets

#### Scenario: Missing SSH credentials
- **WHEN** required secrets (VM_IP, VM_USER, or SSH_PRIVATE_KEY) are not configured
- **THEN** the workflow fails with a clear error message

### Requirement: Pull latest code on VM
The workflow SHALL pull the latest code from the main branch on the target VM.

#### Scenario: Successful code pull
- **WHEN** SSH connection is established
- **THEN** the workflow executes `git pull origin main` in the application directory

#### Scenario: Git conflicts on pull
- **WHEN** local changes conflict with remote changes
- **THEN** the workflow fails and reports the conflict

### Requirement: Environment variable injection
The workflow SHALL create an .env file on the VM with secrets from GitHub Secrets.

#### Scenario: API key environment variable
- **WHEN** the deployment script runs
- **THEN** the workflow writes GEMINI_API_KEY from GitHub Secrets to .env file on the VM

#### Scenario: Multiple environment variables
- **WHEN** multiple secrets need to be deployed
- **THEN** the workflow writes all required environment variables to .env file

### Requirement: Docker container rebuild and restart
The workflow SHALL rebuild and restart Docker containers using the latest code.

#### Scenario: Successful container rebuild
- **WHEN** code is pulled and .env is updated
- **THEN** the workflow executes `docker compose up -d --build`

#### Scenario: Container health check
- **WHEN** containers are restarted
- **THEN** the deployment waits for health checks to pass before completing

### Requirement: Deployment working directory
The workflow SHALL execute deployment commands in a consistent working directory on the VM.

#### Scenario: Change to deployment directory
- **WHEN** SSH connection is established
- **THEN** the workflow changes to the configured deployment directory before executing commands

#### Scenario: Deployment directory does not exist
- **WHEN** the deployment directory path is invalid
- **THEN** the workflow fails with a directory not found error

### Requirement: Deployment status reporting
The workflow SHALL report deployment success or failure status in GitHub Actions.

#### Scenario: Successful deployment
- **WHEN** all deployment steps complete successfully
- **THEN** the workflow job status shows success with green checkmark

#### Scenario: Failed deployment
- **WHEN** any deployment step fails
- **THEN** the workflow job status shows failure with error details

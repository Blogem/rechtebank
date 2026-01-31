## ADDED Requirements

### Requirement: Frontend connects to Traefik proxy network
The frontend service SHALL be configured to connect to the external Traefik proxy network.

#### Scenario: Frontend joins proxy network
- **WHEN** docker-compose starts the frontend service
- **THEN** the frontend container is connected to the external network named `proxy`

#### Scenario: Proxy network does not exist
- **WHEN** the external proxy network is not available
- **THEN** docker-compose fails with network not found error

### Requirement: Frontend has Traefik routing labels
The frontend service SHALL include Traefik labels for HTTP routing based on hostname.

#### Scenario: Traefik routing by hostname
- **WHEN** Traefik processes the frontend container labels
- **THEN** the frontend is accessible via the configured hostname (e.g., rechtbank.example.com)

#### Scenario: Multiple hostname routing
- **WHEN** multiple hostnames are configured in the routing rule
- **THEN** the frontend responds to requests from any configured hostname

### Requirement: TLS certificate configuration
The frontend service SHALL be configured with Traefik labels for automatic TLS certificate generation.

#### Scenario: Let's Encrypt certificate issuance
- **WHEN** the frontend service starts with TLS labels
- **THEN** Traefik automatically requests and configures a Let's Encrypt certificate

#### Scenario: HTTPS redirection
- **WHEN** a client accesses the frontend via HTTP
- **THEN** Traefik redirects the request to HTTPS

### Requirement: Traefik network specification
The frontend service SHALL specify which Docker network Traefik should use for routing.

#### Scenario: Explicit network configuration
- **WHEN** the frontend is connected to multiple networks
- **THEN** the `traefik.docker.network` label specifies the `proxy` network

### Requirement: Backend uses internal network only
The backend service SHALL be accessible only through internal Docker networking, not directly from external sources.

#### Scenario: Backend on internal network
- **WHEN** docker-compose starts the backend service
- **THEN** the backend is connected only to the internal `rechtebank-network`

#### Scenario: No backend port exposure
- **WHEN** docker-compose configuration is loaded
- **THEN** the backend service has no port mappings to the host

### Requirement: Frontend communicates with backend via internal network
The frontend service SHALL connect to both the proxy network and internal network to communicate with the backend.

#### Scenario: Frontend multi-network connection
- **WHEN** docker-compose starts the frontend service
- **THEN** the frontend is connected to both `proxy` and `rechtebank-network` networks

#### Scenario: Frontend proxies API requests
- **WHEN** the frontend receives an API request
- **THEN** the frontend forwards the request to the backend via the internal network

### Requirement: Remove direct port exposure
The docker-compose configuration SHALL NOT expose container ports directly to the host.

#### Scenario: No frontend port mapping
- **WHEN** docker-compose configuration is loaded
- **THEN** the frontend service has no port mappings (previously 5173:80)

#### Scenario: No backend port mapping
- **WHEN** docker-compose configuration is loaded
- **THEN** the backend service has no port mappings (previously 8080:8080)

### Requirement: Traefik service enablement
The frontend service SHALL include the Traefik enable label to activate routing.

#### Scenario: Traefik routing activation
- **WHEN** the frontend container starts
- **THEN** the `traefik.enable=true` label allows Traefik to route traffic to the container

### Requirement: External network definition
The docker-compose configuration SHALL reference the proxy network as an external network.

#### Scenario: External network reference
- **WHEN** docker-compose loads the network configuration
- **THEN** the `proxy` network is marked as external with name `proxy`

#### Scenario: Internal network creation
- **WHEN** docker-compose loads the network configuration
- **THEN** the `rechtebank-network` is created as a bridge network if it doesn't exist

# Apache Reverse Proxy Setup Plan

## Overview
This plan outlines the steps to set up an Apache server as a reverse proxy for two backend applications: a FastAPI (Python) app and a basic Golang app. Everything will be containerized using Docker and orchestrated with Docker Compose.

## Architecture

- **Apache Proxy**: Runs on port 80, forwards requests to backend services
- **FastAPI App**: Python web service running on port 8000
- **Golang App**: Go web service running on port 8080
- **Docker Network**: Internal network for service communication

## Step-by-Step Implementation

### 1. Project Structure Setup

```text
apache-proxy/
├── docker-compose.yml
├── apache/
│   ├── Dockerfile
│   └── httpd.conf
├── fastapi-app/
│   ├── Dockerfile
│   ├── requirements.txt
│   └── main.py
├── golang-app/
│   ├── Dockerfile
│   └── main.go
└── plan.md
```

### 2. FastAPI Application (Backend Service 1)
**Purpose**: Simple REST API that returns a greeting message.

**Key Components**:
- FastAPI framework
- Single endpoint: `GET /` returns JSON {"message": "Hello from FastAPI!"}
- Runs on port 8000 inside container
- Minimal dependencies (fastapi, uvicorn)

**Docker Configuration**:
- Base image: python:3.9-slim
- Install requirements
- Expose port 8000
- Start with uvicorn

### 3. Golang Application (Backend Service 2)
**Purpose**: Basic HTTP server that serves a simple HTML page.

**Key Components**:
- Standard Go HTTP server
- Single handler: responds with "Hello from Golang!"
- Runs on port 8080 inside container
- No external dependencies

**Docker Configuration**:
- Multi-stage build
- Base image: golang:1.21-alpine (build stage)
- Final image: alpine:latest (runtime)
- Expose port 8080
- Start the compiled binary

### 4. Apache Reverse Proxy Configuration
**Purpose**: Route incoming requests to appropriate backend services.

**Key Configuration**:
- Enable mod_proxy and mod_proxy_http modules
- Configure virtual host on port 80
- Route `/api/*` to FastAPI service
- Route `/go/*` to Golang service
- Add proxy headers for proper forwarding
- Enable logging for debugging

**Docker Configuration**:
- Base image: httpd:2.4-alpine
- Copy custom httpd.conf
- Expose port 80

### 5. Docker Compose Orchestration
**Services**:
- `apache-proxy`: Apache server, ports 80:80
- `fastapi-app`: FastAPI service, no external ports
- `golang-app`: Golang service, no external ports

**Network**:
- Custom bridge network: `proxy-network`
- All services connected to this network
- Service discovery via container names

**Volumes** (optional):
- Apache logs volume for persistence

### 6. Testing and Verification
1. Start services: `docker-compose up -d`
2. Test FastAPI: `curl http://localhost/api/`
3. Test Golang: `curl http://localhost/go/`
4. Check Apache logs: `docker-compose logs apache-proxy`
5. Verify container networking: `docker network inspect proxy-network`

### 7. Configuration Details

#### Apache httpd.conf additions

```apache
LoadModule proxy_module modules/mod_proxy.so
LoadModule proxy_http_module modules/mod_proxy_http.so

<VirtualHost *:80>
    ProxyPreserveHost On
    
    # Route /api/* to FastAPI
    ProxyPass /api/ http://fastapi-app:8000/
    ProxyPassReverse /api/ http://fastapi-app:8000/
    
    # Route /go/* to Golang
    ProxyPass /go/ http://golang-app:8080/
    ProxyPassReverse /go/ http://golang-app:8080/
</VirtualHost>
```

#### Docker Compose services

```yaml
version: '3.8'
services:
  apache-proxy:
    build: ./apache
    ports:
      - "80:80"
    networks:
      - proxy-network
  
  fastapi-app:
    build: ./fastapi-app
    networks:
      - proxy-network
  
  golang-app:
    build: ./golang-app
    networks:
      - proxy-network

networks:
  proxy-network:
    driver: bridge
```

### 8. Security Considerations

- Run containers as non-root users
- Use specific base images with known vulnerabilities patched
- Implement proper logging and monitoring
- Consider adding SSL/TLS termination at Apache level
- Add rate limiting and basic authentication if needed

### 9. Deployment Steps

1. Create all necessary files as per structure
2. Build and test each service individually
3. Run `docker-compose up` for development
4. Use `docker-compose up -d` for production
5. Monitor logs and resource usage
6. Scale services as needed

### 10. Troubleshooting

- Check container logs: `docker-compose logs [service-name]`
- Verify network connectivity: `docker exec -it [container] ping [service-name]`
- Test internal endpoints: `docker exec -it apache-proxy curl http://fastapi-app:8000/`
- Validate Apache config: `docker exec -it apache-proxy httpd -t`

This plan provides a complete setup for learning Apache reverse proxy concepts with containerized applications.

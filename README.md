# Apache Reverse Proxy Demo

A minimal Apache httpd reverse proxy in front of two containerized backends:

- **FastAPI (Python)** at path prefix `/api/` - Modern Python web framework
- **Go (Gin framework)** at path prefix `/go/` - High-performance Go web framework

All services run via Docker Compose with Basic Authentication and user context passing.

## What’s included

- Apache httpd configured as a reverse proxy (port 80) with Basic Auth
- FastAPI app (Uvicorn) with JSON API and HTML endpoints
- Golang app (Gin framework) with JSON API and HTML endpoints  
- User authentication headers passed from Apache to backend services
- `docker-compose.yml` to orchestrate everything

## Quick start

```powershell
# From the repo root
docker-compose up -d

# Test JSON APIs through Apache proxy (requires auth: demo/demopass)
curl -u demo:demopass http://localhost/api/        # 🐍 FastAPI Python backend (JSON)
curl -u demo:demopass http://localhost/go/         # ⚡ Golang Gin backend (JSON)

# Test HTML interfaces in browser (or curl)
curl -u demo:demopass http://localhost/python-html # 🐍 FastAPI HTML interface
curl -u demo:demopass http://localhost/go-html     # ⚡ Gin HTML interface

# Default route
curl -u demo:demopass http://localhost/            # Defaults to FastAPI
```

## Authentication

The Apache reverse proxy is protected with Basic Authentication:

- **Username:** `demo`
- **Password:** `demopass`

When accessing <http://localhost/> in a browser, you'll be prompted for these credentials. The htpasswd file is automatically created during the Docker build process.

## Web Interfaces

Both backend services provide beautiful HTML interfaces:

- **FastAPI HTML**: <http://localhost/python-html> - Shows service info, authentication details, and HTTP headers
- **Gin HTML**: <http://localhost/go-html> - Modern interface powered by Gin framework with gradient styling

These pages demonstrate:

- Authentication header passing from Apache to backends
- Service identification (Python vs Go, FastAPI vs Gin)
- Complete HTTP header inspection
- User context preservation across the proxy

## Updating containers

After code changes (e.g., editing Dockerfiles or app code):

```powershell
# Rebuild images and restart containers
docker-compose up -d --build

# Or rebuild specific service
docker-compose up -d --build apache-proxy
```

## Project structure

```text
apache/
  Dockerfile
  httpd.conf
fastapi-app/
  Dockerfile
  requirements.txt
  main.py
golang-app/
  Dockerfile
  main.go
  go.mod
  go.sum
  templates/
    page.html
docker-compose.yml
plan.md
README.md
```

## Notes

- Apache proxies `/api/` to FastAPI (8000) and `/go/` to Go (8080)
- HTML interfaces available at `/python-html` and `/go-html`
- Root `/` is routed to FastAPI by default in `httpd.conf`
- Go app uses Gin framework for improved performance and developer experience
- Authentication headers (`X-Remote-User`, `X-Auth-Type`) are passed from Apache to backends
- For HTTPS, add an SSL-terminating proxy (e.g., Apache mod_ssl or a separate reverse proxy like Traefik/Nginx)

# Apache Reverse Proxy Demo

A minimal Apache httpd reverse proxy in front of two containerized backends:

- FastAPI (Python) at path prefix `/api/`
- Go (Golang) at path prefix `/go/`

All services run via Docker Compose.

## What’s included

- Apache httpd configured as a reverse proxy (port 80)
- FastAPI app (Uvicorn) with `/` and `/health`
- Golang app with `/` and `/health`
- `docker-compose.yml` to orchestrate everything

## Quick start

```powershell
# From the repo root
docker-compose up -d

# Test through Apache proxy (requires auth: demo/demopass)
curl -u demo:demopass http://localhost/api/  # 🐍 FastAPI Python backend
curl -u demo:demopass http://localhost/go/   # ⚡ Golang backend
curl -u demo:demopass http://localhost/      # Defaults to FastAPI
```

## Authentication

The Apache reverse proxy is protected with Basic Authentication:

- **Username:** `demo`
- **Password:** `demopass`

When accessing <http://localhost/> in a browser, you'll be prompted for these credentials. The htpasswd file is automatically created during the Docker build process.

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
docker-compose.yml
plan.md
README.md
```

## Notes

- Apache proxies `/api/` to FastAPI (8000) and `/go/` to Go (8080)
- Root `/` is routed to FastAPI by default in `httpd.conf`
- For HTTPS, add an SSL-terminating proxy (e.g., Apache mod_ssl or a separate reverse proxy like Traefik/Nginx)

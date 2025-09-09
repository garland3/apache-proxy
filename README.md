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

# Test through Apache proxy
curl http://localhost/api/
curl http://localhost/go/
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

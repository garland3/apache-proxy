from fastapi import FastAPI, Request

app = FastAPI()

@app.get("/")
async def read_root(request: Request):
    # Read authentication headers passed from Apache
    remote_user = request.headers.get("X-Remote-User", "unknown")
    auth_type = request.headers.get("X-Auth-Type", "none")
    auth_time = request.headers.get("X-Authenticated-Time", "unknown")
    
    return {
        "message": "🐍 Hello from FastAPI Python backend!",
        "service": "FastAPI",
        "language": "Python",
        "authenticated_user": remote_user,
        "auth_method": auth_type,
        "auth_timestamp": auth_time,
        "all_headers": dict(request.headers)
    }

@app.get("/health")
async def health_check():
    return {"status": "healthy"}

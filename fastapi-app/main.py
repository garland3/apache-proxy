from fastapi import FastAPI, Request
from fastapi.responses import HTMLResponse

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

@app.get("/html", response_class=HTMLResponse)
async def read_root_html(request: Request):
    # Read authentication headers passed from Apache
    remote_user = request.headers.get("X-Remote-User", "unknown")
    auth_type = request.headers.get("X-Auth-Type", "none")
    auth_time = request.headers.get("X-Authenticated-Time", "unknown")
    
    headers_html = "".join([f"<tr><td>{k}</td><td>{v}</td></tr>" for k, v in request.headers.items()])
    
    html_content = f"""
    <!DOCTYPE html>
    <html>
    <head>
        <title>FastAPI Python Backend</title>
        <style>
            body {{ font-family: Arial, sans-serif; margin: 40px; background: #f5f5f5; }}
            .container {{ background: white; padding: 30px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }}
            .header {{ color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 10px; }}
            .info {{ background: #ecf0f1; padding: 15px; border-radius: 5px; margin: 20px 0; }}
            .auth-info {{ background: #d5f4e6; padding: 15px; border-radius: 5px; margin: 20px 0; }}
            table {{ width: 100%; border-collapse: collapse; margin-top: 20px; }}
            th, td {{ padding: 8px; text-align: left; border-bottom: 1px solid #ddd; }}
            th {{ background-color: #3498db; color: white; }}
            .emoji {{ font-size: 2em; }}
        </style>
    </head>
    <body>
        <div class="container">
            <h1 class="header"><span class="emoji">🐍</span> FastAPI Python Backend</h1>
            
            <div class="info">
                <h3>Service Information</h3>
                <p><strong>Service:</strong> FastAPI</p>
                <p><strong>Language:</strong> Python</p>
                <p><strong>Message:</strong> Hello from FastAPI Python backend!</p>
            </div>
            
            <div class="auth-info">
                <h3>Authentication Details</h3>
                <p><strong>User:</strong> {remote_user}</p>
                <p><strong>Auth Method:</strong> {auth_type}</p>
                <p><strong>Auth Timestamp:</strong> {auth_time}</p>
            </div>
            
            <h3>All HTTP Headers</h3>
            <table>
                <tr><th>Header</th><th>Value</th></tr>
                {headers_html}
            </table>
            
            <p style="margin-top: 30px; color: #7f8c8d;">
                <strong>Proxied through Apache HTTP Server</strong><br>
                This page demonstrates authentication header passing from Apache to FastAPI.
            </p>
        </div>
    </body>
    </html>
    """
    return html_content

@app.get("/health")
async def health_check():
    return {"status": "healthy"}

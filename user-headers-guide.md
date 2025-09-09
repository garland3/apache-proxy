# Passing User Authentication to Downstream Containers

## How It Works

Apache can pass authentication information to backend containers using HTTP headers. This allows your applications to know who the authenticated user is without handling authentication themselves.

## Headers Added by Apache

| Header | Description | Example |
|--------|-------------|---------|
| `X-Remote-User` | Username of authenticated user | `demo` |
| `X-Auth-Type` | Authentication method used | `Basic`, `Kerberos`, `OAuth2` |
| `X-Authenticated-Time` | When authentication occurred | `2025-09-08_23:15:30` |

## Apache Configuration

```apache
<Location "/">
    AuthType Basic
    AuthName "Restricted Area"
    AuthUserFile "/usr/local/apache2/conf/.htpasswd"
    Require valid-user
    
    # Pass user info to downstream
    RequestHeader set X-Remote-User %{REMOTE_USER}s
    RequestHeader set X-Auth-Type "Basic"
    RequestHeader set X-Authenticated-Time %{TIME_YEAR}s-%{TIME_MON}s-%{TIME_DAY}s_%{TIME_HOUR}s:%{TIME_MIN}s:%{TIME_SEC}s
</Location>
```

## Reading Headers in Applications

### FastAPI (Python)
```python
from fastapi import FastAPI, Request

@app.get("/")
async def read_root(request: Request):
    remote_user = request.headers.get("X-Remote-User", "unknown")
    auth_type = request.headers.get("X-Auth-Type", "none")
    return {"user": remote_user, "auth": auth_type}
```

### Go
```go
func handler(w http.ResponseWriter, r *http.Request) {
    remoteUser := r.Header.Get("X-Remote-User")
    authType := r.Header.Get("X-Auth-Type")
    
    fmt.Fprintf(w, "User: %s, Auth: %s", remoteUser, authType)
}
```

### Node.js/Express
```javascript
app.get('/', (req, res) => {
    const user = req.headers['x-remote-user'] || 'unknown';
    const authType = req.headers['x-auth-type'] || 'none';
    res.json({ user, authType });
});
```

## Security Benefits

1. **Centralized Authentication** - Only Apache handles auth
2. **Zero Trust** - Apps don't need to trust user input
3. **Audit Trail** - All auth decisions logged in Apache
4. **Flexibility** - Change auth method without touching app code

## Available Apache Variables

| Variable | Description |
|----------|-------------|
| `%{REMOTE_USER}s` | Authenticated username |
| `%{AUTH_TYPE}s` | Authentication method |
| `%{TIME_YEAR}s` | Current year |
| `%{TIME_MON}s` | Current month |
| `%{TIME_DAY}s` | Current day |
| `%{REQUEST_URI}s` | Requested URL path |
| `%{HTTP_USER_AGENT}s` | Browser user agent |

## Testing

After rebuilding with these changes:

```powershell
# Test with authentication
Invoke-WebRequest -Uri http://localhost/api/ -Headers @{Authorization = "Basic $([Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes('demo:demopass')))"}

# Response will include:
# {
#   "message": "Hello from FastAPI!",
#   "authenticated_user": "demo",
#   "auth_method": "Basic",
#   "auth_timestamp": "2025-09-08_23:15:30"
# }
```

This pattern works with **any authentication method** (Basic, OAuth2, Kerberos) - just change the `RequestHeader` directives accordingly!

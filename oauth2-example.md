# OAuth2 Setup Example

This would replace Basic Auth with OAuth2 (Google/GitHub/Microsoft).

## Required Changes:

### 1. Update Dockerfile
```dockerfile
FROM httpd:2.4-alpine

# Install mod_auth_openidc (not available in Alpine by default)
RUN apk add --no-cache apache2-mod-auth-openidc

COPY httpd-oauth.conf /usr/local/apache2/conf/httpd.conf
EXPOSE 80
CMD ["httpd-foreground"]
```

### 2. OAuth2 httpd.conf
```apache
LoadModule auth_openidc_module modules/mod_auth_openidc.so
LoadModule proxy_module modules/mod_proxy.so
LoadModule proxy_http_module modules/mod_proxy_http.so

<VirtualHost *:80>
    # OAuth2 Configuration
    OIDCProviderMetadataURL https://accounts.google.com/.well-known/openid_configuration
    OIDCClientID your-google-client-id.googleusercontent.com
    OIDCClientSecret your-client-secret
    OIDCRedirectURI http://localhost/redirect_uri
    OIDCCryptoPassphrase random-secret-phrase

    <Location "/">
        AuthType openid-connect
        Require valid-user
    </Location>

    # Your existing proxy rules...
    ProxyPass /api/ http://fastapi-app:8000/
    ProxyPassReverse /api/ http://fastapi-app:8000/
</VirtualHost>
```

### 3. Google OAuth Setup
1. Go to Google Cloud Console
2. Create OAuth 2.0 credentials
3. Set redirect URI: http://localhost/redirect_uri
4. Get client ID/secret

## User Experience:
- Visit http://localhost/api/
- Redirected to Google login
- Login with Google account
- Redirected back to your app
- Authenticated!

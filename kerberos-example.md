# Kerberos Authentication Setup

This would replace Basic Auth with Kerberos/Active Directory integration.

## Prerequisites

- Active Directory domain
- Apache server joined to domain
- Kerberos keytab file for HTTP service

## Required Changes

### 1. Update Dockerfile

```dockerfile
FROM httpd:2.4-alpine

# Install Kerberos packages
RUN apk add --no-cache krb5 krb5-dev apache2-mod-auth-kerb

# Copy Kerberos config and keytab
COPY krb5.conf /etc/krb5.conf
COPY http.keytab /etc/httpd/conf/http.keytab

COPY httpd-kerberos.conf /usr/local/apache2/conf/httpd.conf
EXPOSE 80
CMD ["httpd-foreground"]
```

### 2. Kerberos httpd.conf

```apache
LoadModule auth_kerb_module modules/mod_auth_kerb.so
LoadModule headers_module modules/mod_headers.so
LoadModule proxy_module modules/mod_proxy.so
LoadModule proxy_http_module modules/mod_proxy_http.so

<VirtualHost *:80>
    # Kerberos Configuration
    <Location "/">
        AuthType Kerberos
        AuthName "Active Directory Login"
        KrbMethodNegotiate On
        KrbMethodK5Passwd On
        KrbAuthRealms EXAMPLE.COM
        Krb5KeyTab /etc/httpd/conf/http.keytab
        KrbServiceName HTTP
        KrbSaveCredentials On
        
        # Pass username to downstream
        RequestHeader set X-Remote-User %{REMOTE_USER}s
        RequestHeader set X-Auth-Type "Kerberos"
        
        Require valid-user
    </Location>

    # Proxy with user headers
    ProxyPreserveHost On
    ProxyRequests Off
    
    ProxyPass /api/ http://fastapi-app:8000/
    ProxyPassReverse /api/ http://fastapi-app:8000/
    
    ProxyPass /go/ http://golang-app:8080/
    ProxyPassReverse /go/ http://golang-app:8080/
</VirtualHost>
```

### 3. Kerberos Configuration Files

**krb5.conf:**
```ini
[libdefaults]
    default_realm = EXAMPLE.COM
    dns_lookup_realm = false
    dns_lookup_kdc = true
    
[realms]
    EXAMPLE.COM = {
        kdc = dc1.example.com
        admin_server = dc1.example.com
    }
    
[domain_realm]
    .example.com = EXAMPLE.COM
    example.com = EXAMPLE.COM
```

**Generate keytab (on domain controller):**
```cmd
ktpass -out http.keytab -princ HTTP/your-server@EXAMPLE.COM -mapuser apache-service -pass YourPassword
```

## Domain Setup Requirements

1. **Create service account** in Active Directory
2. **Generate keytab** for HTTP service principal
3. **Configure DNS** for your Apache server
4. **Join Apache server** to domain (if using Windows)
5. **Set SPNs** for the service account

## User Experience

- **Windows clients**: Automatic SSO (no password prompt)
- **Non-domain clients**: Username/password prompt
- **Works with**: Internet Explorer, Chrome, Firefox (with config)

## Security Benefits

- Single Sign-On for domain users
- No password transmission
- Mutual authentication
- Ticket-based security
- Enterprise audit trails

## Limitations

- Requires Active Directory infrastructure
- Complex setup and maintenance
- Client configuration may be needed
- Firewall considerations (Kerberos ports)

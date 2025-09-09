package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type UserInfo struct {
    Message           string            `json:"message"`
    Service           string            `json:"service"`
    Language          string            `json:"language"`
    AuthenticatedUser string            `json:"authenticated_user"`
    AuthMethod        string            `json:"auth_method"`
    AuthTimestamp     string            `json:"auth_timestamp"`
    AllHeaders        map[string]string `json:"all_headers"`
}

func getUserInfo(c *gin.Context) UserInfo {
    // Read authentication headers passed from Apache
    remoteUser := c.GetHeader("X-Remote-User")
    if remoteUser == "" {
        remoteUser = "unknown"
    }
    
    authType := c.GetHeader("X-Auth-Type")
    if authType == "" {
        authType = "none"
    }
    
    authTime := c.GetHeader("X-Authenticated-Time")
    if authTime == "" {
        authTime = "unknown"
    }
    
    // Collect all headers
    allHeaders := make(map[string]string)
    for key, values := range c.Request.Header {
        if len(values) > 0 {
            allHeaders[key] = values[0]
        }
    }
    
    return UserInfo{
        Message:           "⚡ Hello from Golang backend!",
        Service:           "Go HTTP Server (Gin)",
        Language:          "Go",
        AuthenticatedUser: remoteUser,
        AuthMethod:        authType,
        AuthTimestamp:     authTime,
        AllHeaders:        allHeaders,
    }
}

func jsonHandler(c *gin.Context) {
    userInfo := getUserInfo(c)
    c.JSON(http.StatusOK, userInfo)
}

func htmlHandler(c *gin.Context) {
    userInfo := getUserInfo(c)
    
    c.Header("Content-Type", "text/html")
    c.HTML(http.StatusOK, "page.html", gin.H{
        "userInfo": userInfo,
    })
}

func healthHandler(c *gin.Context) {
    c.String(http.StatusOK, "OK")
}

func main() {
    // Set Gin to release mode (less verbose logging)
    gin.SetMode(gin.ReleaseMode)
    
    r := gin.Default()
    
    // Load HTML template
    r.LoadHTMLGlob("templates/*")
    
    // Routes
    r.GET("/", jsonHandler)
    r.GET("/html", htmlHandler)
    r.GET("/health", healthHandler)
    
    // Start server
    r.Run(":8080")
}

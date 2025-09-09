package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type UserInfo struct {
    Message         string            `json:"message"`
    Service         string            `json:"service"`
    Language        string            `json:"language"`
    AuthenticatedUser string          `json:"authenticated_user"`
    AuthMethod      string            `json:"auth_method"`
    AuthTimestamp   string            `json:"auth_timestamp"`
    AllHeaders      map[string]string `json:"all_headers"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Read authentication headers passed from Apache
    remoteUser := r.Header.Get("X-Remote-User")
    if remoteUser == "" {
        remoteUser = "unknown"
    }
    
    authType := r.Header.Get("X-Auth-Type")
    if authType == "" {
        authType = "none"
    }
    
    authTime := r.Header.Get("X-Authenticated-Time")
    if authTime == "" {
        authTime = "unknown"
    }
    
    // Collect all headers for debugging
    allHeaders := make(map[string]string)
    for name, values := range r.Header {
        if len(values) > 0 {
            allHeaders[name] = values[0]
        }
    }
    
    userInfo := UserInfo{
        Message:         "⚡ Hello from Golang backend!",
        Service:         "Go HTTP Server",
        Language:        "Go",
        AuthenticatedUser: remoteUser,
        AuthMethod:      authType,
        AuthTimestamp:   authTime,
        AllHeaders:      allHeaders,
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(userInfo)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "OK")
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/health", healthHandler)

    fmt.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

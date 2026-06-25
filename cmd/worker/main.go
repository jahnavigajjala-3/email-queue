package main

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
    "github.com/jahnavigajjala/email-queue/internal/job"
    "github.com/redis/go-redis/v9"
)
var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    fmt.Println("Worker started, waiting for jobs...")

    for {
        result, err := rdb.BRPop(ctx, 0*time.Second, "email_queue").Result()
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        var j job.Job
        json.Unmarshal([]byte(result[1]), &j)

        fmt.Printf("Sending email to: %s | Subject: %s | Body: %s\n", j.To, j.Subject, j.Body)
    }
}
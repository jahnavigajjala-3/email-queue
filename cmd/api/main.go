package main
import (
	"encoding/json"
    "fmt"
    "net/http"
    "context"
    "github.com/jahnavigajjala-3/email-queue/internal/job"
    "github.com/redis/go-redis/v9"
)
func main() {
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    http.HandleFunc("/send-emails", handleSendEmails)
    fmt.Println("API server running on port 8080")
    http.ListenAndServe(":8080", nil)
}
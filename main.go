// main.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

// Message contract
type Message struct {
	PostID      string `json:"postId"`
	PostUser    string `json:"postUser"`
	PostMessage string `json:"postMessage"`
}

var ctx = context.Background()

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		// Addr:     "redis-service.go-api-practice.svc.cluster.local:6379",
		Addr: "redis-service:6379",
		Password: "", // No password
		DB:       0,  // Default DB
	})

	// Check connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis:", pong)

	// Populate Redis with 1000 messages
	populateTestData(client)

	// Set up HTTP server
	router := mux.NewRouter()
	router.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		getAllMessages(w, client)
	}).Methods("GET")

	// Start HTTP server
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func populateTestData(client *redis.Client) {
	for i := 1; i <= 1000; i++ {
		message := Message{
			PostID:      fmt.Sprintf("post%d", i),
			PostUser:    fmt.Sprintf("user%d", i),
			PostMessage: fmt.Sprintf("This is message %d", i),
		}

		// Convert message to JSON
		messageJSON, err := json.Marshal(message)
		if err != nil {
			log.Printf("Error marshaling message: %v", err)
			continue
		}

		// Store message in Redis
		err = client.Set(ctx, message.PostID, messageJSON, 0).Err()
		if err != nil {
			log.Printf("Error storing message in Redis: %v", err)
		}
	}
	fmt.Println("Test data populated in Redis.")
}

func getAllMessages(w http.ResponseWriter, client *redis.Client) {
	keys, err := client.Keys(ctx, "*").Result()
	if err != nil {
		http.Error(w, "Error retrieving messages from Redis", http.StatusInternalServerError)
		return
	}

	var messages []Message
	for _, key := range keys {
		val, err := client.Get(ctx, key).Result()
		if err != nil {
			log.Printf("Error retrieving message from Redis: %v", err)
			continue
		}

		var message Message
		err = json.Unmarshal([]byte(val), &message)
		if err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		messages = append(messages, message)
	}

	// Convert messages to JSON
	responseJSON, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, "Error encoding messages to JSON", http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

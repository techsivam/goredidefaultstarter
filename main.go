package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-redis/redis/v8"
	rejson "github.com/nitishm/go-rejson/v4"
)

type pingResponse struct {
	Status string `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)

	/* client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}) */

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Use "redis" instead of "localhost"
		Password: "",
		DB:       0,
	})
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(client)
	switch r.Method {
	case http.MethodGet:
		if path == "/" {
			fmt.Println("SWITCH:", path)
			homepage := "This is a sample JSON homepage for the microservice."
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(homepage)
		} else if path == "/ping" {
			fmt.Println("SWITCH:", path)
			response := pingResponse{Status: "working"}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else if strings.HasPrefix(path, "/tenant/") {
			tenantName := strings.TrimPrefix(path, "/tenant/")
			fmt.Println("SWITCH:", path)

			value, err := rh.JSONGet(tenantName, ".") // Use rh.JSONGet instead of client.Get
			fmt.Println(value)
			if err != nil {
				http.Error(w, "Tenant not found", http.StatusNotFound)
			} else {
				w.Header().Set("Content-Type", "application/json")
				valueBytes, ok := value.([]byte) // Type assertion to []byte
				if !ok {
					http.Error(w, "Error converting JSON data", http.StatusInternalServerError)
					return
				}
				fmt.Fprint(w, string(valueBytes)) // Convert valueBytes to a string
			}
		} else {
			http.Error(w, "Invalid request path", http.StatusBadRequest)
		}

	case http.MethodPost:
		if strings.HasPrefix(path, "/") {
			tenantName := strings.TrimPrefix(path, "/")
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusBadRequest)
				return
			}

			var jsonData map[string]interface{}
			err = json.Unmarshal(body, &jsonData)
			if err != nil {
				http.Error(w, "Error unmarshaling JSON", http.StatusBadRequest)
				return
			}

			_, err = rh.JSONSet(tenantName, ".", jsonData) // Use rh.JSONSet instead of client.Set
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				fmt.Fprintf(w, "POST request for tenant %s, saved JSON data\n", tenantName)
			}
		} else {
			http.Error(w, "Invalid request path", http.StatusBadRequest)
		}

	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

func redisClient(message string) {
	redisURL := os.Getenv("REDIS_URL")
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Set("Message", message, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Key Already set")
}

func main() {
	url := "http://" + os.Getenv("API_JSON") + ":8080"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			log.Fatal(err2)
		}
		bodyString := string(bodyBytes)
		var result map[string]interface{}
		json.Unmarshal([]byte(bodyString), &result)
		fmt.Println("message:", result["message"])
		message := result["message"].(string)
		redisClient(message)
	}
}

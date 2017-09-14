package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

var client *redis.Client

func handler(w http.ResponseWriter, r *http.Request) {
	result := incr()
	fmt.Fprintf(w, "Hi there! %d", result)
}

func handlerICon(w http.ResponseWriter, r *http.Request) {}

func incr() int64 {
	log.Println("Increase")
	result, err := client.Incr("counter").Result()
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	log.Println("http starting on 8080")
	http.HandleFunc("/favicon.ico", handlerICon)
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

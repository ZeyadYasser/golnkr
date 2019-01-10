package main

import (
	"golnkr/handlers"
	"net/http"
	"log"
	"github.com/go-redis/redis"
)

func ConnectDB() *redis.Client{
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	db := ConnectDB()

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
  	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.Handle("/gomake", handlers.GomakeHandler(db))

	mux.Handle("/goget/", http.StripPrefix("/goget/", handlers.GogetHandler(db)))

	mux.Handle("/", http.HandlerFunc(handlers.IndexHandler))

	log.Println("Listening...")
    http.ListenAndServe(":80", mux)
}
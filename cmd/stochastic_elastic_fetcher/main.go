package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	} else {
		log.Print("Yay .env file found")
	}
}

func main() {
	twitterAPIKey, exists := os.LookupEnv("TWITTER_API_KEY")

	if exists {
		fmt.Println(twitterAPIKey)
	}

	twitterAPISecretKey, exists := os.LookupEnv("TWITTER_API_SECRET_KEY")

	if exists {
		fmt.Println(twitterAPISecretKey)
	}

	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, tal paskaro ya king!\n")
	}

	http.HandleFunc("/callback", helloHandler)
	log.Println("Listing for requests at http://localhost:8089/hello")
	log.Fatal(http.ListenAndServe(":8089", nil))
}

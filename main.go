package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func main() {

	if os.Getenv("SERVER_PORT") == "" || os.Getenv("SERVER_PORT") == "" {
		os.Setenv("SERVER_PORT", "8181")
		os.Setenv("MESSAGE", "Hello, world!")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		println(time.Now().Local().String() + " - Message: " + os.Getenv("MESSAGE"))
		json.NewEncoder(w).Encode(os.Getenv("MESSAGE"))
	})

	println("Server running on http://localhost:" + os.Getenv("SERVER_PORT"))
	http.ListenAndServe(":" + os.Getenv("SERVER_PORT"), nil)
}

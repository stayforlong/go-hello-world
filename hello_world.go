package main

import (
	"fmt"
	"net/http"
)

const (
	port = "80"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, "Hello world")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, "I'm running")
}

func main() {
	fmt.Printf("Starting server... at http://localhost:%v.\n", port)
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/health", HealthCheck)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	fmt.Println("Program exit")
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\nis the way\n\n\n%s", r.URL.Path[1:], getVCAP())
}

func main() {
	port := getPort()
	log.Printf("Starting server on port %s\n", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func getPort() string {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		return "8080"
	}

	return port
}

func getVCAP() string {
	return os.Getenv("VCAP_APPLICATION")
}

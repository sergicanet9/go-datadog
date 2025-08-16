package main

import (
	"fmt"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The root route has been reached!")
	fmt.Fprintf(w, "The server is running")
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Generating a report log and sending it to Datadog APM with automatic instrumentation...")

	// Simulate a blocking operation that takes 1 second
	time.Sleep(1 * time.Second)

	fmt.Fprintf(w, "Report generated successfully!")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/report", reportHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

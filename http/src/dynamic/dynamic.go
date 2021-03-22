package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func currentHour(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2021 05:06:05")
	fmt.Fprintf(w, "<h1>Current hour: %s</h1>", s)
}

func main () {
	http.HandleFunc("/", currentHour)

	log.Println("Executing...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
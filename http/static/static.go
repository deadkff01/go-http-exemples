package main

import (
	"log"
	"net/http"
)

func main () {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Executing...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
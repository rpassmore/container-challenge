package main

import (
	"embed"
	"log"
	"net/http"
)

// content holds our static web server content
//go:embed static/*
var content embed.FS

func main() {
	fs := http.FileServer(http.FS(content))
	http.Handle("/static/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

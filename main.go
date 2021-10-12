package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

// content holds our static web server content
//go:embed static/*
var content embed.FS

func main() {
	contentStatic, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(contentStatic)))

	log.Println("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

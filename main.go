package main

import (
	"github.com/Tambarie/online-store/application/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env with godotenv: %s", err)
	}
	server.Start()
}

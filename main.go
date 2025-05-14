package main

import (
	"github.com/joho/godotenv"
	"log"
	"mod/services/telegram"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var tgToken string = mustToken()
	var host string = mustHost()

	tgClient = telegram.NewClient(host, tgToken)
	eventsProcessor := telegram2.NewClient(host, tgToken, files.New('storage'))
	//
	//fetcher = fetcher.New()
	//processor = processor.New()
	//
	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	tgToken, exists := os.LookupEnv("TG_TOKEN")

	if !exists || tgToken == "" {
		log.Fatal("Environment variable TG_TOKEN not set")
	}

	return tgToken
}

func mustHost() string {
	host, exists := os.LookupEnv("HOST")

	if !exists || host == "" {
		log.Fatal("Environment variable HOST not set")
	}

	return host
}

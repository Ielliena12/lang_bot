package main

import (
	"github.com/ielliena/lang_bot/consumer"
	"github.com/ielliena/lang_bot/events/processor"
	"github.com/ielliena/lang_bot/services/telegram"
	"github.com/ielliena/lang_bot/storage/files"
	"github.com/joho/godotenv"
	"log"
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

	tgClient := telegram.NewClient(host, tgToken)
	eventsProcessor := processor.NewProcessor(tgClient, files.NewStorage("storage"))

	if err := consumer.New(eventsProcessor, eventsProcessor, 100).Start(); err != nil {
		log.Fatal(err)
	}
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

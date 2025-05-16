package main

import (
	"github.com/ielliena/lang_bot/config"
	"github.com/ielliena/lang_bot/consumer"
	"github.com/ielliena/lang_bot/events/processor"
	"github.com/ielliena/lang_bot/services/telegram"
	"github.com/ielliena/lang_bot/storage/files"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	tgToken := config.GetToken()
	host := config.GetHost()

	tgClient := telegram.NewClient(host, tgToken)

	eventsProcessor := processor.NewProcessor(tgClient, files.NewStorage("storage"))
	if err := consumer.New(eventsProcessor, eventsProcessor, 100).Start(); err != nil {
		log.Fatal(err)
	}
}

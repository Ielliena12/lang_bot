package config

import (
	"log"
	"os"
)

func GetToken() string {
	tgToken, exists := os.LookupEnv("TG_TOKEN")

	if !exists || tgToken == "" {
		log.Fatal("Environment variable TG_TOKEN not set")
	}

	return tgToken
}

func GetHost() string {
	host, exists := os.LookupEnv("HOST")

	if !exists || host == "" {
		log.Fatal("Environment variable HOST not set")
	}

	return host
}

func GetOwner() string {
	owner, exists := os.LookupEnv("OWNER")

	if !exists || owner == "" {
		log.Fatal("Environment variable OWNER not set")
	}

	return owner
}

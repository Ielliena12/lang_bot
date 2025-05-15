package processor

import (
	"fmt"
	"github.com/ielliena/lang_bot/storage"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	StartCmd = "/start"
	HideCmd  = "/hide"
	GetCmd   = "/get"
)

func (processor *Processor) checkCommand(text string, chatID int) error {
	if strconv.Itoa(chatID) != mustOwner() {
		return nil
	}

	text = strings.TrimSpace(text)

	switch text {
	case StartCmd:
		return processor.tg.SendMessage(chatID, &storage.Message{MessageItem: "Добрый день"})
	case GetCmd:
		return processor.getWord(chatID)
	default:
		if err := processor.saveWord(text); err != nil {
			fmt.Println(err)
			return err
		}
		return processor.tg.SendMessage(chatID, &storage.Message{MessageItem: "Слово сохранено в словарь"})
	}
}

func (processor *Processor) saveWord(word string) error {
	message := &storage.Message{
		MessageItem: word,
	}

	if err := processor.storage.Save(message); err != nil {
		return fmt.Errorf("file does not saved: %w", err)
	}

	return nil
}

func (processor *Processor) getWord(chatID int) error {
	word, err := processor.storage.PickWord()
	if err != nil {
		return fmt.Errorf("word does not pick: %w", err)
	}

	if err := processor.tg.SendMessage(chatID, word); err != nil {
		return fmt.Errorf("message was not sended: %w", err)
	}

	return nil
}

func mustOwner() string {
	owner, exists := os.LookupEnv("OWNER")

	if !exists || owner == "" {
		log.Fatal("Environment variable OWNER not set")
	}

	return owner
}

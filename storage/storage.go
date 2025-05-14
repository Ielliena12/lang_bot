package storage

import (
	"crypto/sha1"
	"fmt"
	"io"
)

type Storage interface {
	Save(message *Message) error
	PickWord() (*Message, error)
}

type Message struct {
	MessageItem string
}

func (message *Message) Hash() (string, error) {
	h := sha1.New()

	if _, err = io.WriteString(h, message.MessageItem); err != nil {
		return "", fmt.Errorf("word was not hashed: %w", err)
	}

	return fmt.Sprintln(h.Sum(nil)), nil
}

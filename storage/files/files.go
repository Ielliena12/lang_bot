package files

import (
	"encoding/gob"
	"fmt"
	"github.com/ielliena/lang_bot/storage"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Storage struct {
	basePath string
}

func NewStorage(basePath string) Storage {
	return Storage{basePath}
}

func (storage Storage) Save(message *storage.Message) (err error) {
	filePath := filepath.Join(storage.basePath, "eng")
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	fileName, err := fileName(message)
	if err != nil {
		return fmt.Errorf("failed to generate filename: %w", err)
	}

	filePath = filepath.Join(filePath, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Проверяем, что MessageItem поддерживает gob-сериализацию
	if err := gob.NewEncoder(file).Encode(message.MessageItem); err != nil {
		return fmt.Errorf("failed to encode message: %w", err)
	}

	return nil
}

func (storage Storage) PickWord() (message *storage.Message, err error) {
	filePath := filepath.Join(storage.basePath, "eng")

	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, fmt.Errorf("word was not picked: %w", err)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("word was not picked: %w", err)
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	return storage.decodeMessage(filepath.Join(filePath, file.Name()))
}

func (s Storage) decodeMessage(filePath string) (*storage.Message, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("file was not opened: %w", err)
	}
	defer file.Close()

	var message string
	if err := gob.NewDecoder(file).Decode(&message); err != nil {
		return nil, fmt.Errorf("file was not decoded: %w", err)
	}

	return &storage.Message{
		MessageItem: message,
	}, nil
}

func fileName(message *storage.Message) (string, error) {
	return message.Hash()
}

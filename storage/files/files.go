package files

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"math/rand"
	"mod/storage"
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
	defer func() { err = fmt.Errorf("word was not saved: %w", err) }()

	filePath := filepath.Join(s.basePath, "eng")
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}

	fileName, err := fileName(message)
	if err != nil {
		return err
	}
	filePath = filepath.Join(filePath, fileName)
	file, err := os.Create(filePath)

	defer file.Close()

	if err := gob.NewEncoder(file).Encode(message); err != nil {
		return err
	}

	return nil
}

func (storage Storage) PickWord() (message *storage.Message, err error) {
	defer func() { err = fmt.Errorf("word was not picked: %w", err) }()

	filePath := filepath.Join(s.basePath, "eng")

	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, nil
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	return storage.decodeMessage(filePath.Join(filePath, file.Name()))
}

func (storage Storage) decodeMessage(filePath string) (*storage.Message, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("file was not opened: %w", err)
	}

	var message = storage.Message
	if err := gob.NewEncoder(file).Decode(&message); err != nil {
		return nil, fmt.Errorf("file was not decoded: %w", err)
	}

	return &message, nil
}

func fileName(message *storage.Message) (string, error) {
	return message.Hash()
}

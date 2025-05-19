package processor

import (
	"fmt"
	"github.com/ielliena/lang_bot/config"
	"github.com/ielliena/lang_bot/events"
	"github.com/ielliena/lang_bot/services/telegram"
	"github.com/ielliena/lang_bot/storage"
	"strconv"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

func NewProcessor(tg *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      tg,
		storage: storage,
	}
}

func (processor *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := processor.tg.GetUpdates(processor.offset, limit)
	if err != nil {
		return nil, fmt.Errorf("event was not geted: %w", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, update := range updates {
		res = append(res, event(update))
	}

	processor.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (processor *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return processor.processMessage(event)
	default:
		return fmt.Errorf("empty message: %w", '-')
	}
}

func (processor *Processor) RemindWord() error {
	chatId, _ := strconv.Atoi(config.GetOwner())
	event := events.Event{
		Type:   events.Message,
		Text:   "/get",
		ChatID: chatId,
	}
	return processor.processMessage(event)
}

func (processor *Processor) processMessage(event events.Event) error {
	if err := processor.checkCommand(event.Text, event.ChatID); err != nil {
		return fmt.Errorf("message was not processed: %w", '-')
	}

	return nil
}

func event(update telegram.Updates) events.Event {
	return events.Event{
		Type:   events.Message,
		Text:   update.Message.Text,
		ChatID: update.Message.Chat.ID,
	}
}

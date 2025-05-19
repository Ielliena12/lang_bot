package consumer

import (
	"fmt"
	"github.com/ielliena/lang_bot/events"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) *Consumer {
	return &Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (consumer Consumer) Start() error {
	for {
		gotEvents, err := consumer.fetcher.Fetch(consumer.batchSize)
		if err != nil {
			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}

		for _, gotEvent := range gotEvents {
			if err := consumer.processor.Process(gotEvent); err != nil {
				continue
			}
		}
	}
}

func (consumer Consumer) RemindWord() error {
	ticker := time.NewTicker(15 * time.Minute)
	for _ = range ticker.C {
		if err := consumer.processor.RemindWord(); err != nil {
			return fmt.Errorf("message was not processed: %w", '-')
		}
	}

	return nil
}

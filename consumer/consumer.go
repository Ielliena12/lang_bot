package consumer

import (
	"mod/events"
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

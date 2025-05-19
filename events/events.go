package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(event Event) error
	RemindWord() error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type   Type
	Text   string
	ChatID int
}

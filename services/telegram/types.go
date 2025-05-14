package telegram

type UpdatesResponse struct {
	Ok     bool      `json:"ok"`
	Result []Updates `json:"result"`
}

type Updates struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	ID int `json:"id"`
}

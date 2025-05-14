package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func NewClient(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: "bot" + token,
		client:   http.Client{},
	}
}

func (client *Client) GetUpdates(offset int, limit int) ([]Updates, error) {
	query := url.Values{}
	query.Add("offset", strconv.Itoa(offset))
	query.Add("limit", strconv.Itoa(limit))

	data, err := client.sendRequest("getUpdates", query)
	if err != nil {
		return nil, err
	}

	var resp UpdatesResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

func (client *Client) SendMessage(chatID int, message string) error {
	query := url.Values{}
	query.Add("chatID", strconv.Itoa(chatID))
	query.Add("message", message)

	_, err := client.sendRequest("sendMessage", query)
	if err != nil {
		return fmt.Errorf("message was not sended: %w", err)
	}

	return nil
}

func (client *Client) sendRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   client.host,
		Path:   path.Join(client.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("request was not sended: %w", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request was not sended: %w", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	return body, nil
}

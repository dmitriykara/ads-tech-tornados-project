package events

import (
	"bytes"
	"encoding/json"
	"net/http"
	"ttb-service/internal/utils"
)

// EventPublisher интерфейс для отправки событий
type EventPublisher interface {
	Publish(event Event) error
}

// HTTPPublisher отправляет события через HTTP
type HTTPPublisher struct {
	url    string
	client *http.Client
	logger *utils.Logger
}

// NewHTTPPublisher создает новый HTTP-паблишер
func NewHTTPPublisher(url string, client *http.Client, logger *utils.Logger) *HTTPPublisher {
	return &HTTPPublisher{
		url:    url,
		client: client,
		logger: logger,
	}
}

// Publish отправляет событие через HTTP POST
func (hp *HTTPPublisher) Publish(event Event) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		hp.logger.Error("Error marshaling event: " + err.Error())
		return err
	}

	// Оборачиваем []byte в io.Reader
	reqBody := bytes.NewReader(jsonData)

	req, err := http.NewRequest("POST", hp.url, reqBody)
	if err != nil {
		hp.logger.Error("Error creating HTTP request: " + err.Error())
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := hp.client.Do(req)
	if err != nil {
		hp.logger.Error("Error sending HTTP request: " + err.Error())
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		hp.logger.Error("Received non-OK response from server: " + resp.Status)
		return nil
	}

	hp.logger.Info("Event successfully published.")
	return nil
}

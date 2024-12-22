package events

import "time"

// Event структура для базовых данных события
type Event struct {
	ID        string    `json:"id"`
	EventType string    `json:"event_type"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

// NewEvent создает новое событие с текущим временем
func NewEvent(eventType, payload string) Event {
	return Event{
		ID:        generateEventID(),
		EventType: eventType,
		Payload:   payload,
		Timestamp: time.Now(),
	}
}

// Генерация уникального идентификатора события (можно заменить на более сложную логику)
func generateEventID() string {
	return time.Now().Format("20060102150405") // Пример простого ID
}

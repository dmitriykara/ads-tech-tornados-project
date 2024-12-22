package events

import (
	"ttb-service/internal/utils"
)

// EventHandler интерфейс для обработки событий
type EventHandler interface {
	Handle(event Event) error
}

// EventProcessor структура, которая управляет обработчиками событий
type EventProcessor struct {
	handlers map[string]EventHandler
	logger   *utils.Logger
}

// NewEventProcessor создает новый процессор для событий с переданным логгером
func NewEventProcessor(logger *utils.Logger) *EventProcessor {
	return &EventProcessor{
		handlers: make(map[string]EventHandler),
		logger:   logger,
	}
}

// RegisterHandler регистрирует обработчик для конкретного типа события
func (ep *EventProcessor) RegisterHandler(eventType string, handler EventHandler) {
	ep.handlers[eventType] = handler
}

// ProcessEvent обрабатывает событие, направляя его в соответствующий обработчик
func (ep *EventProcessor) ProcessEvent(event Event) error {
	handler, exists := ep.handlers[event.EventType]
	if !exists {
		ep.logger.Error("No handler registered for event type: " + event.EventType)
		return nil
	}

	if err := handler.Handle(event); err != nil {
		ep.logger.Error("Error handling event: " + err.Error())
		return err
	}

	ep.logger.Info("Event processed successfully: " + event.EventType)
	return nil
}

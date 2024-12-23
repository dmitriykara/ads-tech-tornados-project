package events

// PolicyEvent represents an event triggered by a policy update or application.
type PolicyEvent struct {
	EventType string `json:"eventType"` // Type of event (e.g., PolicyApplied, PolicyUpdated)
	Details   string `json:"details"`   // Details or additional information about the event
}

func (pe PolicyEvent) ToEvent() Event {
	return Event{
		EventType: pe.EventType,
		Payload:   pe.Details,
	}
}

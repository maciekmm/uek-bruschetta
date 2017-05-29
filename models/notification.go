package models

type NotificationPriority int

const (
	LowNotificationPriority NotificationPriority = iota
	MediumNotificationPriority
	HighNotificationPriority
)

type Notification struct {
	ID       uint64               `json:"id,omitempty" db:"id"`
	EventID  uint64               `json:"event_id,omitempty" db:"event_id"`
	Channel  ChannelType          `json:"channel,omitempty" db:"channel"`
	Metadata string               `json:"metadata,omitempty" db:"metadata"`
	Priority NotificationPriority `json:"priority,omitempty" db:"priority"`
}

package httpserver

import "time"

type AddRequest struct {
	UUID               string        `json:"uuid"`
	Header             string        `json:"header"`
	DateTime           time.Time     `json:"dt"`
	Duration           time.Duration `json:"duration"`
	Description        string        `json:"description"`
	UserID             string        `json:"user_id"`
	NotificationBefore time.Duration `json:"notification_before"`
}

type ModifyRequest struct {
	Header             string        `json:"header"`
	DateTime           time.Time     `json:"dt"`
	Duration           time.Duration `json:"duration"`
	Description        string        `json:"description"`
	UserID             string        `json:"user_id"`
	NotificationBefore time.Duration `json:"notification_before"`
}

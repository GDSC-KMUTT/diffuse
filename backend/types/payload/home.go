package payload

import "time"

type HomeNotification struct {
	Type     *string        `json:"type"`
	Href     *string        `json:"href"`
	Date     *time.Time     `json:"date"`
	Messages []*HomeMessage `json:"messages"`
}

type HomeMessage struct {
	Type *string `json:"type"`
	Text *string `json:"text"`
	Href *string `json:"href"`
}

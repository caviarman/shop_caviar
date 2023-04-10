package entity

import "time"

type User struct {
	ID               int       `json:"id"`
	Name             string    `json:"name,omitempty"`
	Email            string    `json:"email,omitempty"`
	Phone            string    `json:"phone,omitempty"`
	TelegramUsername string    `json:"telegram_username,omitempty"`
	TelegramID       int       `json:"telegram_id,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

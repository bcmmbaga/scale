package models

import "time"

type MessagesWallet struct {
	ID                      string            `json:"id"`
	Count                   int               `json:"count"`
	Price                   int               `json:"price"`
	SenderID                string            `json:"sender_id"`
	SenderVerificationState VerificationState `json:"sender_verification_state"`
	CreatedAt               time.Time         `json:"created_at"`
	UpdatedAt               time.Time         `json:"updated_at"`
}

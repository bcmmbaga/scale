package model

import (
	"time"
)

type Account struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Email             string            `json:"email"`
	Phone             string            `json:"phone"`
	Password          string            `json:"password"`
	Type              AccountType       `json:"type"`
	VerificationState VerificationState `json:"verification_state"`
	MessageWallet     *MessagesWallet   `json:"message_wallet"`
	OrganizationInfo  *OrganizationInfo `json:"organization_info"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
}

type OrganizationInfo struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Phone         string          `json:"phone"`
	MessageWallet *MessagesWallet `json:"message_wallet"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type AccountType string

const (
	AccountTypePersonal     AccountType = "PERSONAL"
	AccountTypeOrganization AccountType = "ORGANIZATION"
)

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Account struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Email             string            `json:"email"`
	Phone             string            `json:"phone"`
	Password          string            `json:"password"`
	Type              AccountType       `json:"type"`
	VerificationState VerificationState `json:"verificationState"`
	MessageWallet     *MessagesWallet   `json:"messageWallet"`
	OrganizationInfo  *OrganizationInfo `json:"organizationInfo"`
}

type MessagesWallet struct {
	ID                      string            `json:"id"`
	Count                   int               `json:"count"`
	Price                   int               `json:"price"`
	SenderID                string            `json:"senderID"`
	SenderVerificationState VerificationState `json:"senderVerificationState"`
}

type OrganizationInfo struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Phone         string          `json:"phone"`
	MessageWallet *MessagesWallet `json:"messageWallet"`
}

type AccountType string

const (
	AccountTypePersonal      AccountType = "PERSONAL"
	AccountTypeOrganinzation AccountType = "ORGANINZATION"
)

var AllAccountType = []AccountType{
	AccountTypePersonal,
	AccountTypeOrganinzation,
}

func (e AccountType) IsValid() bool {
	switch e {
	case AccountTypePersonal, AccountTypeOrganinzation:
		return true
	}
	return false
}

func (e AccountType) String() string {
	return string(e)
}

func (e *AccountType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccountType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccountType", str)
	}
	return nil
}

func (e AccountType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VerificationState string

const (
	VerificationStateVerified              VerificationState = "VERIFIED"
	VerificationStateUnverified            VerificationState = "UNVERIFIED"
	VerificationStateVerificationRequested VerificationState = "VERIFICATION_REQUESTED"
)

var AllVerificationState = []VerificationState{
	VerificationStateVerified,
	VerificationStateUnverified,
	VerificationStateVerificationRequested,
}

func (e VerificationState) IsValid() bool {
	switch e {
	case VerificationStateVerified, VerificationStateUnverified, VerificationStateVerificationRequested:
		return true
	}
	return false
}

func (e VerificationState) String() string {
	return string(e)
}

func (e *VerificationState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VerificationState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VerificationState", str)
	}
	return nil
}

func (e VerificationState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

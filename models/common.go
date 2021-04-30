package models

type VerificationState string

const (
	VerificationStateVerified              VerificationState = "VERIFIED"
	VerificationStateUnverified            VerificationState = "UNVERIFIED"
	VerificationStateVerificationRequested VerificationState = "VERIFICATION_REQUESTED"
)

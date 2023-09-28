package token

import "time"

// Maker is an interface to managing tokens
type Maker interface {
	// CreateToken create a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken check if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}

package token

import "time"

type Meker interface {
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

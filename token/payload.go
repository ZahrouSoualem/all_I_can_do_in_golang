package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	ErrExpired = "expired token"
	ErrInvalid = "invalid token"
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"name"`
	Create_at time.Time `json:"createat"`
	Expire_at time.Time `json:"expireat"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		return nil, errors.New("we couldbn't create the id of the payload")
	}

	payload := &Payload{
		Id:        id,
		Username:  username,
		Create_at: time.Now(),
		Expire_at: time.Now().Add(duration),
	}

	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.Expire_at) {
		return errors.New(ErrExpired)
	}
	return nil
}

package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// the secret key shouldn't be too short
const minSecretKey = 32

type JWTMaker struct {
	secretkey string
}

func NewJwtMaker(secretkey string) (*JWTMaker, error) {

	if len(secretkey) < minSecretKey {
		return nil, errors.New(secretkey)
	}

	return &JWTMaker{secretkey}, nil
}

func (m *JWTMaker) CreateToken(username string, dur time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, dur)

	if err != nil {
		return "", &Payload{}, errors.New("we couldn't create paylaod")
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	if err != nil {
		return "", &Payload{}, errors.New("we couldn't create token")
	}
	// if you want to sign sth there must always be a key
	token, err := jwtToken.SignedString([]byte(m.secretkey))

	if err != nil {
		return "", &Payload{}, errors.New("we couldn't create token")
	}

	return token, payload, nil
}

func (m *JWTMaker) VerifyToken(token string) (*Payload, error) {
	errFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New(ErrInvalid)
		}
		return []byte(m.secretkey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, errFunc)

	if err != nil {
		return nil, errors.New("we couldn't extract the token")
	}

	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, errors.New("we couldn't extract the payload")
	}

	return payload, nil
}

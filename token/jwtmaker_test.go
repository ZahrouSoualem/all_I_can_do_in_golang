package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tester/util"
)

func CreateMaker(t *testing.T) *JWTMaker {
	s := util.RandomString(34)

	maker, err := NewJwtMaker(s)
	require.NoError(t, err)
	require.NotEmpty(t, maker)
	require.Equal(t, maker.secretkey, s)

	return maker
}

func TestCreateMaker(t *testing.T) {
	maker := CreateMaker(t)
	require.NotEmpty(t, maker)

	user := util.RandomString(6)

	token, payload, err := maker.CreateToken(user, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, payload.Username, user)
	require.NotEmpty(t, token)

	payload2, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload2)

	issuedat := time.Now()
	expiredat := time.Now().Add(time.Minute)

	require.Equal(t, payload.Id, payload2.Id)
	require.Equal(t, payload.Username, payload2.Username)
	require.WithinDuration(t, payload.Create_at, issuedat, time.Second)
	require.WithinDuration(t, payload.Expire_at, expiredat, time.Second)

}

func TestCreateToken(t *testing.T) {
	maker := CreateMaker(t)
	require.NotEmpty(t, maker)
}

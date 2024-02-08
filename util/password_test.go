package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashpassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashpassword)

	err = CheckPassword(hashpassword, password)
	require.NoError(t, err)

	password2 := RandomString(7)

	hashpassword, err = HashPassword(password2)
	require.NoError(t, err)

	err = CheckPassword(hashpassword, password)
	require.Error(t, err)
}

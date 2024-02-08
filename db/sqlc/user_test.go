package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tester/util"
)

func TestCreateUse(t *testing.T) {
	CreateRandomUser(t)
}

func CreateRandomUser(t *testing.T) User {
	password := util.RandomString(6)
	Hashpassword, err := util.HashPassword(password)
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:     util.RandomString(6),
		HashPassword: Hashpassword,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.HashPassword, Hashpassword)
	require.NotZero(t, user.ID)

	return user
}

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.ID, user2.ID)
	require.Equal(t, user.Username, user2.Username)

}

func TestGetUserByName(t *testing.T) {
	user := CreateRandomUser(t)

	user2, err := testQueries.GetUserByName(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.ID, user2.ID)
	require.Equal(t, user.Username, user2.Username)

}

func TestDeleteUser(t *testing.T) {
	user := CreateRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
	user2, err := testQueries.GetUser(context.Background(), user.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestUpdateUser(t *testing.T) {
	user := CreateRandomUser(t)

	arg := UpdateUserParams{
		ID:       user.ID,
		Username: util.RandomString(6),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.ID, user2.ID)
	require.NotEqual(t, user.Username, user2.Username)

}

package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassowrd1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassowrd1)

	err = CheckPassword(password, hashedPassowrd1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPassowrd1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassowrd2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassowrd2)
	require.NotEqual(t, hashedPassowrd1, hashedPassowrd2)
}

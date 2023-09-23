package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	HashPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, HashPassword1)

	err = CheckPassword(password, HashPassword1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, HashPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	HashPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, HashPassword2)
	require.NotEqual(t, HashPassword1, HashPassword2)
}

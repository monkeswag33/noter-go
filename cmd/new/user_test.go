package new

import (
	"testing"

	"github.com/monkeswag33/noter-go/argon2"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/errordef"
	"github.com/stretchr/testify/assert"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

type userValidationTester struct {
	username       string
	password       string
	validator      func(string) error
	expectedResult error
}

func TestCreateUser(t *testing.T) {
	var username string = "testcreateuser"
	user, err := createUser([]string{username}, password)
	assert.NoError(t, err)
	assert.NoError(t, insertUser(user))
	assert.Equal(t, user.Username, username)
	matches, err := argon2.VerifyPass(password, user.Password)
	assert.NoError(t, err)
	assert.True(t, matches)
}

func TestUserValidator(t *testing.T) {
	var username string = "testusernamevalidator"
	var user db.User = db.User{
		Username: username,
		Password: password,
	}
	assert.NoError(t, database.CreateUser(&user))
	var testCases []userValidationTester = []userValidationTester{
		// Usernames
		{
			username:       "hi", // Too short
			validator:      newUserValidateUsername,
			expectedResult: errordef.ErrUsernameTooShort,
		},
		{
			username:       username, // Already exists
			validator:      newUserValidateUsername,
			expectedResult: errordef.ErrUserAlreadyExists,
		},
		{
			username:       "test$*!#", // Must contain only alphanumeric characters
			validator:      newUserValidateUsername,
			expectedResult: errordef.ErrUsernameMustContainAlphaNumeric,
		},
		{
			username:       "validusername123",
			validator:      newUserValidateUsername,
			expectedResult: nil,
		},
		// Passwords
		{
			password:       "short", // Too short
			validator:      newUserValidatePassword,
			expectedResult: errordef.ErrPasswordTooShort,
		},
		{
			password:       "password", // insecure
			validator:      newUserValidatePassword,
			expectedResult: passwordvalidator.Validate("password", minEntropyBits),
		},
		{
			password:       password, // Valid password
			validator:      newUserValidatePassword,
			expectedResult: nil,
		},
	}
	for _, testCase := range testCases {
		var argument string = testCase.username
		if len(testCase.password) != 0 {
			argument = testCase.password
		}
		assert.Equal(t, testCase.validator(argument), testCase.expectedResult)
	}
}

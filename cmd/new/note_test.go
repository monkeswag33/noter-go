package new

import (
	"testing"

	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/errordef"
	"github.com/stretchr/testify/assert"
)

type noteValidatationTester struct {
	noteName       string
	username       string
	validator      func(string) error
	expectedResult error
}

func TestCreateNote(t *testing.T) {
	// Create user
	noteName, username := "Testing Create Note", "testcreatenote"
	var user db.User = db.User{
		Username: username,
		Password: password,
	}
	assert.NoError(t, database.CreateUser(&user))
	exists, err := database.CheckUserExists(db.User{
		Username: username,
	})
	assert.NoError(t, err)
	assert.True(t, exists)
	note, err := createNote([]string{noteName}, username, noteBody)
	assert.NoError(t, err)
	assert.NoError(t, insertNote(note))
	assert.Equal(t, username, note.User.Username)
	assert.Equal(t, noteName, note.Name)
	assert.Equal(t, noteBody, note.Body)
}

func TestNoteValidator(t *testing.T) {
	noteName, username := "Testing Note Name Validator", "testnotenamevalidator"
	var user db.User = db.User{
		Username: username,
		Password: password,
	}
	assert.NoError(t, database.CreateUser(&user))
	var note db.Note = db.Note{
		Name: noteName,
		User: db.User{
			Username: username,
		},
	}
	assert.NoError(t, database.CreateNote(&note))
	var testCases []noteValidatationTester = []noteValidatationTester{
		{
			noteName:       "Hi", // Too short
			validator:      newNoteValidateNoteName,
			expectedResult: errordef.ErrNoteNameTooShort,
		},
		{
			noteName:       noteName, // Note already exists
			validator:      newNoteValidateNoteName,
			expectedResult: errordef.ErrNoteAlreadyExists,
		},
		{
			username:       "userthatdoesntexist", // User doesn't exist
			validator:      newNoteValidateUsername,
			expectedResult: errordef.ErrUserDoesntExist,
		},
		{
			noteName:       "Valid Note Name",
			validator:      newNoteValidateNoteName,
			expectedResult: nil,
		},
		{
			username:       username,
			validator:      newNoteValidateUsername,
			expectedResult: nil,
		},
	}
	for _, testCase := range testCases {
		var argument string = testCase.noteName
		if len(testCase.username) != 0 {
			argument = testCase.username
		} else {
			continue // Something is wrong with this test case as no argument is specified
		}
		assert.Equal(t, testCase.validator(argument), testCase.expectedResult)
	}
}

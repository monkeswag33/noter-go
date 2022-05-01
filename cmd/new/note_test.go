package new

import (
	"testing"
	"time"

	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/global"
	"github.com/stretchr/testify/assert"
)

type noteValidatationTester struct {
	NoteName       string
	Username       string
	Validator      func(string) error
	ExpectedResult error
}

func TestCreateNote(t *testing.T) {
	global.InitTesterDB()
	defer global.ShutdownDB()
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
	assert.NoError(t, insertNote(&note))
	assert.Equal(t, username, note.User.Username)
	assert.Equal(t, noteName, note.Name)
	assert.Equal(t, noteBody, note.Body)
}

func TestNoteValidator(t *testing.T) {
	global.InitTesterDB()
	time.Sleep(2 * time.Second)
	exists, err := database.CheckUserExists(db.User{
		Username: "testcreatenote",
	})
	assert.NoError(t, err)
	assert.True(t, exists)
	// noteName, username := "Testing Note Name Validator", "testnotenamevalidator"
	// db.CreateUser(username, password)
	// if _, err := createNote([]string{noteName}, username, noteBody); err != nil {
	// 	t.Fatal(err)
	// }
	// var testCases []noteValidatationTester = []noteValidatationTester{
	// 	{
	// 		NoteName:       "Hi", // Too short
	// 		Validator:      newNoteValidateNoteName,
	// 		ExpectedResult: errordef.ErrNoteNameTooShort,
	// 	},
	// 	{
	// 		NoteName:       noteName, // Note already exists
	// 		Validator:      newNoteValidateNoteName,
	// 		ExpectedResult: errordef.ErrNoteAlreadyExists,
	// 	},
	// 	{
	// 		Username:       "userthatdoesntexist", // User doesn't exist
	// 		Validator:      newNoteValidateUsername,
	// 		ExpectedResult: errordef.ErrUserDoesntExist,
	// 	},
	// 	{
	// 		NoteName:       "Valid Note Name",
	// 		Validator:      newNoteValidateNoteName,
	// 		ExpectedResult: nil,
	// 	},
	// 	{
	// 		Username:       username,
	// 		Validator:      newNoteValidateUsername,
	// 		ExpectedResult: nil,
	// 	},
	// }
	// for _, testCase := range testCases {
	// 	var argument string
	// 	if len(testCase.NoteName) != 0 {
	// 		argument = testCase.NoteName
	// 	} else if len(testCase.Username) != 0 {
	// 		argument = testCase.Username
	// 	} else {
	// 		continue // Something is wrong with this test case as no argument is specified
	// 	}
	// 	assert.Equal(t, testCase.Validator(argument), testCase.ExpectedResult)
	// }
	// db.DeleteUser(username)
}

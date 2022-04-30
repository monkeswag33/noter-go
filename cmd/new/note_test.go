package new

import (
	"testing"
)

type noteValidatationTester struct {
	NoteName       string
	Username       string
	Validator      func(string) error
	ExpectedResult error
}

func TestCreateNote(t *testing.T) {
	// Create user
	// _, username := "Testing Create Note", "testcreatenote"
	// db.CreateUser(username, password)
	// fmt.Println(db.CheckUserExists(username))
	// note, err := createNote([]string{noteName}, username, noteBody)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// assert.Equal(t, username, note.User.Username)
	// assert.Equal(t, noteName, note.Name)
	// assert.Equal(t, noteBody, note.Body)
	// db.DeleteUser(username)
	// assert.False(t, db.CheckUserExists("ishank"))
}

func TestNoteValidator(t *testing.T) {
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

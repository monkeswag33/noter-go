package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var user User = User{
	Username: "user",
	Password: "password",
}

var note Note = Note{
	Name: "test",
	Body: "test",
	User: User{
		Username: "user",
	},
}

func TestGetUsers(t *testing.T) {
	database, err := InitTesterDB()
	assert.NoError(t, err)
	assert.NoError(t, database.DB.Create(&user).Error)
	getUsers, err := database.GetUsers(User{})
	assert.NoError(t, err)
	var getUser User = getUsers[0]
	assert.EqualValues(t, user, getUser)
}

func TestGetNotes(t *testing.T) {
	database, err := InitTesterDB()
	assert.NoError(t, err)
	assert.NoError(t, database.DB.Create(&user).Error)
	assert.NoError(t, database.DB.Create(&note).Error)
	getNotes, err := database.GetNotes(Note{})
	assert.NoError(t, err)
	var getNote Note = getNotes[0]
	fmt.Println(getNote.Name)
}

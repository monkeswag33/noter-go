package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	var clonedUser User = user
	database, err := InitTesterDB()
	assert.NoError(t, err)
	assert.NoError(t, database.DB.Create(&clonedUser).Error)
	getUsers, err := database.GetUsers(User{})
	assert.NoError(t, err)
	var getUser User = getUsers[0]
	assert.EqualValues(t, clonedUser, getUser)
}

func TestGetNotes(t *testing.T) {
	var clonedUser User = user
	var clonedNote Note = note
	database, err := InitTesterDB()
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.NoError(t, database.DB.Create(&clonedUser).Error)
	clonedNote.User = clonedUser
	assert.NoError(t, database.DB.Create(&clonedNote).Error)
	getNotes, err := database.GetNotes(Note{})
	assert.NoError(t, err)
	assert.Len(t, getNotes, 1)
	assert.EqualValues(t, clonedNote, getNotes[0])
}

package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	var clonedUser User = user
	database, err := InitTesterDB()
	assert.NoError(t, err)
	assert.NoError(t, database.DB.Create(&clonedUser).Error)
	var users []User
	database.DB.Find(&users)
	assert.Len(t, users, 1)
	var insertedUser User = users[0]
	assert.EqualValues(t, clonedUser, insertedUser)
}

func TestCreateNote(t *testing.T) {
	database, err := InitTesterDB()
	assert.NoError(t, err)
	var notes []Note
	var clonedUser User = user
	var clonedNote Note = note
	assert.NoError(t, database.DB.Create(&clonedUser).Error)
	clonedNote.User = clonedUser
	assert.NoError(t, database.CreateNote(&clonedNote))
	assert.NoError(t, database.DB.Preload("User").Find(&notes).Error)
	assert.Len(t, notes, 1)
	assert.EqualValues(t, clonedNote, notes[0])
}

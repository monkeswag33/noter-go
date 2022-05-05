package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	database, err := InitTesterDB()
	assert.NoError(t, err)
	var clonedUser User = user
	assert.NoError(t, database.DB.Create(&clonedUser).Error)
	var users []User
	assert.NoError(t, database.DB.Find(&users).Error)
	assert.Len(t, users, 1)
	assert.NoError(t, database.DeleteUser(clonedUser))
	assert.NoError(t, database.DB.Find(&users).Error)
	assert.Len(t, users, 0)
}

func TestDeleteNote(t *testing.T) {
	database, err := InitTesterDB()
	assert.NoError(t, err)
	var clonedUser User = user
	var clonedNote Note = note
	assert.NoError(t, database.DB.Create(&clonedUser).Error)
	clonedNote.User = clonedUser
	assert.NoError(t, database.DB.Create(&clonedNote).Error)
	var notes []Note
	assert.NoError(t, database.DB.Find(&notes).Error)
	assert.Len(t, notes, 1)
	assert.NoError(t, database.DeleteNote(clonedNote))
	assert.NoError(t, database.DB.Find(&notes).Error)
	assert.Len(t, notes, 0)
}

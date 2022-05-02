package global

import (
	"testing"

	"github.com/monkeswag33/noter-go/db"
	"github.com/stretchr/testify/assert"
)

func TestTesterDB(t *testing.T) {
	var database *db.DB = InitTesterDB()
	var user db.User = db.User{
		Username: "user",
		Password: "password",
	}
	assert.NoError(t, database.CreateUser(&user))

	var note db.Note = db.Note{
		Name: "Test Note",
		Body: "Test Body",
		User: db.User{
			Username: "user",
		},
	}
	assert.NoError(t, database.CreateNote(&note))
}

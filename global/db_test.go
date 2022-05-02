package global

import (
	"testing"

	"github.com/monkeswag33/noter-go/db"
	"github.com/stretchr/testify/assert"
)

func TestTesterDB(t *testing.T) {
	database := InitTesterDB()
	assert.IsType(t, &db.DB{}, database)
}

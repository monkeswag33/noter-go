package new

import (
	"os"
	"testing"

	"github.com/monkeswag33/noter-go/db"
	"github.com/sirupsen/logrus"
)

const (
	noteBody string = "Test Body"
	password string = "vEry$Ecur3pA$$w0rD4%23"
)

func TestMain(m *testing.M) {
	var err error
	database, err = db.InitTesterDB()
	if err != nil {
		logrus.Fatal(err)
	}
	if err := database.Init(); err != nil {
		logrus.Fatal(err)
	}
	code := m.Run()
	if err := database.Close(); err != nil {
		logrus.Fatal(err)
	}
	os.Exit(code)
}

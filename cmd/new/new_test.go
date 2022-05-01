package new

import (
	"os"
	"testing"

	"github.com/monkeswag33/noter-go/global"
)

const (
	noteBody string = "Test Body"
	password string = "vEry$Ecur3pA$$w0rD4%23"
)

func TestMain(m *testing.M) {
	database = global.InitTesterDB()
	database.Init()
	code := m.Run()
	database.Close()
	os.Exit(code)
}

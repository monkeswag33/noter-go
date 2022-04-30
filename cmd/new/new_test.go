package new

import (
	"os"
	"testing"
)

const (
	noteBody string = "Test Body"
	password string = "verysecurepassword123"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

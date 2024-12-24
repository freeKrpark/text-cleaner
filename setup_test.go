package cleaner

import (
	"os"
	"testing"
)

var cleaner Cleaner

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

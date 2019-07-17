package filesystem

import (
	"testing"

	"github.com/kisrobot/oss/tests"
)

func TestAll(t *testing.T) {
	fileSystem := New("/tmp")
	tests.TestAll(fileSystem, t)
}

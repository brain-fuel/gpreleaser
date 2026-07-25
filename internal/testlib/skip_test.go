package testlib

import (
	"testing"

	"goforge.dev/gpreleaser/internal/pipe"
)

func TestAssertSkipped(t *testing.T) {
	AssertSkipped(t, pipe.Skip("skip"))
}

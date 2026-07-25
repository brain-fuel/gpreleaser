package cmd

import (
	"testing"

	"goforge.dev/gpreleaser/internal/skips"
	"goforge.dev/gpreleaser/pkg/context"
	"github.com/stretchr/testify/require"
)

func requireAll(tb testing.TB, ctx *context.Context, keys ...skips.Key) {
	tb.Helper()
	for _, key := range keys {
		require.True(tb, ctx.Skips[string(key)], "expected %q to be true, but was false", key)
	}
}

package healthcheck

import (
	"testing"

	"goforge.dev/gpreleaser/internal/testctx"
	"goforge.dev/gpreleaser/pkg/config"
	"github.com/stretchr/testify/require"

	// langs to init.
	_ "goforge.dev/gpreleaser/internal/builders/bun"
	_ "goforge.dev/gpreleaser/internal/builders/deno"
	_ "goforge.dev/gpreleaser/internal/builders/golang"
	_ "goforge.dev/gpreleaser/internal/builders/node"
	_ "goforge.dev/gpreleaser/internal/builders/rust"
	_ "goforge.dev/gpreleaser/internal/builders/zig"
)

func TestSystemDependencies(t *testing.T) {
	ctx := testctx.Wrap(t.Context())
	require.Equal(t, []string{"git"}, system{}.Dependencies(ctx))
}

func TestSystemStringer(t *testing.T) {
	require.NotEmpty(t, system{}.String())
}

func TestBuildDependencies(t *testing.T) {
	ctx := testctx.WrapWithCfg(t.Context(), config.Project{
		Builds: []config.Build{
			{Builder: "bun"},
			{Builder: "deno"},
			{Builder: "go"},
			{Builder: "rust"},
			{Builder: "zig"},
			{Builder: "node"},
		},
	})
	require.Equal(t, []string{
		"bun",
		"deno",
		"go",
		"cargo",
		"rustup",
		"cargo-zigbuild",
		"zig",
		"zig", // dedup happens later on
		"node",
	}, builds{}.Dependencies(ctx))
}

func TestBuildStringer(t *testing.T) {
	require.NotEmpty(t, builds{}.String())
}

func TestHealthCheckers(t *testing.T) {
	require.NotEmpty(t, HealthCheckers)
}

func TestDependencyCheckers(t *testing.T) {
	require.NotEmpty(t, DependencyCheckers)
}

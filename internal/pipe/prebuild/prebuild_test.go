package prebuild

import (
	"testing"

	"goforge.dev/gpreleaser/internal/testctx"
	"goforge.dev/gpreleaser/internal/testlib"
	"goforge.dev/gpreleaser/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		ctx := testctx.WrapWithCfg(t.Context(), config.Project{
			Env:    []string{"FOO=bar"},
			Builds: []config.Build{{Main: "{{ .Env.FOO }}"}},
		})
		require.NoError(t, Pipe{}.Run(ctx))
		require.Equal(t, "bar", ctx.Config.Builds[0].Main)
	})

	t.Run("empty", func(t *testing.T) {
		ctx := testctx.WrapWithCfg(t.Context(), config.Project{
			Env:    []string{"FOO="},
			Builds: []config.Build{{Main: "{{ .Env.FOO }}"}},
		})
		require.NoError(t, Pipe{}.Run(ctx))
		require.Equal(t, ".", ctx.Config.Builds[0].Main)
	})

	t.Run("bad", func(t *testing.T) {
		ctx := testctx.WrapWithCfg(t.Context(), config.Project{
			Builds: []config.Build{{Main: "{{ .Env.FOO }}"}},
		})
		testlib.RequireTemplateError(t, Pipe{}.Run(ctx))
	})
}

func TestString(t *testing.T) {
	require.NotEmpty(t, Pipe{}.String())
}

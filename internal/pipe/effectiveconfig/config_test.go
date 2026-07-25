package effectiveconfig

import (
	"os"
	"path/filepath"
	"testing"

	"goforge.dev/gpreleaser/internal/testctx"
	"goforge.dev/gpreleaser/internal/testlib"
	"goforge.dev/gpreleaser/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestPipeDescription(t *testing.T) {
	require.Empty(t, Pipe{}.String())
}

func TestRun(t *testing.T) {
	folder := testlib.Mktmp(t)
	dist := filepath.Join(folder, "dist")
	require.NoError(t, os.Mkdir(dist, 0o755))
	ctx := testctx.WrapWithCfg(t.Context(), config.Project{
		Dist: dist,
	})

	require.NoError(t, Pipe{}.Run(ctx))
	bts, err := os.ReadFile(filepath.Join(dist, "config.yaml"))
	require.NoError(t, err)
	require.NotEmpty(t, string(bts))
}

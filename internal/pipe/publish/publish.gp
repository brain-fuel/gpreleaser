// Package publish contains the publishing pipe.
package publish

import (
	"fmt"

	"goforge.dev/gpreleaser/internal/middleware/errhandler"
	"goforge.dev/gpreleaser/internal/middleware/logging"
	"goforge.dev/gpreleaser/internal/middleware/skip"
	"goforge.dev/gpreleaser/internal/pipe/artifactory"
	"goforge.dev/gpreleaser/internal/pipe/aur"
	"goforge.dev/gpreleaser/internal/pipe/aursources"
	"goforge.dev/gpreleaser/internal/pipe/blob"
	"goforge.dev/gpreleaser/internal/pipe/brew"
	"goforge.dev/gpreleaser/internal/pipe/cask"
	"goforge.dev/gpreleaser/internal/pipe/chocolatey"
	"goforge.dev/gpreleaser/internal/pipe/custompublishers"
	"goforge.dev/gpreleaser/internal/pipe/docker"
	dockerv2 "goforge.dev/gpreleaser/internal/pipe/docker/v2"
	"goforge.dev/gpreleaser/internal/pipe/dockerdigest"
	"goforge.dev/gpreleaser/internal/pipe/ko"
	"goforge.dev/gpreleaser/internal/pipe/krew"
	"goforge.dev/gpreleaser/internal/pipe/mcp"
	"goforge.dev/gpreleaser/internal/pipe/milestone"
	"goforge.dev/gpreleaser/internal/pipe/nix"
	"goforge.dev/gpreleaser/internal/pipe/release"
	"goforge.dev/gpreleaser/internal/pipe/scoop"
	"goforge.dev/gpreleaser/internal/pipe/sign"
	"goforge.dev/gpreleaser/internal/pipe/snapcraft"
	"goforge.dev/gpreleaser/internal/pipe/upload"
	"goforge.dev/gpreleaser/internal/pipe/winget"
	"goforge.dev/gpreleaser/internal/skips"
	"goforge.dev/gpreleaser/pkg/context"
)

// Publisher should be implemented by pipes that want to publish artifacts.
type Publisher interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Publish(ctx *context.Context) error
}

// New publish pipeline.
func New() Pipe {
	return Pipe{
		pipeline: []Publisher{
			blob.Pipe{},
			upload.Pipe{},
			artifactory.Pipe{},
			docker.Pipe{},
			docker.ManifestPipe{},
			dockerv2.Publish{},
			dockerdigest.Pipe{},
			ko.Pipe{},
			sign.DockerPipe{},
			snapcraft.Pipe{},
			// This should be one of the last steps
			release.Pipe{},
			// brew et al use the release URL, so, they should be last
			nix.New(),
			winget.Pipe{},
			brew.Pipe{},
			cask.Pipe{},
			aur.Pipe{},
			aursources.Pipe{},
			krew.Pipe{},
			scoop.Pipe{},
			chocolatey.Pipe{},
			mcp.New(),
			milestone.Pipe{},
			custompublishers.Pipe{},
		},
	}
}

// Pipe that publishes artifacts.
type Pipe struct {
	pipeline []Publisher
}

func (Pipe) String() string                 { return "publishing" }
func (Pipe) Skip(ctx *context.Context) bool { return skips.Any(ctx, skips.Publish) }

func (p Pipe) Run(ctx *context.Context) error {
	memo := errhandler.Memo{}
	for _, publisher := range p.pipeline {
		if err := skip.Maybe(
			publisher,
			logging.PadLog(
				publisher.String(),
				errhandler.Handle(publisher.Publish),
			),
		)(ctx); err != nil {
			if ig, ok := publisher.(Continuable); ok && ig.ContinueOnError() && !ctx.FailFast {
				memo.Memorize(fmt.Errorf("%s: %w", publisher.String(), err))
				continue
			}
			return fmt.Errorf("%s: failed to publish artifacts: %w", publisher.String(), err)
		}
	}
	return memo.Error()
}

type Continuable interface {
	ContinueOnError() bool
}

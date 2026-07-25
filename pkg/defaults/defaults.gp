// Package defaults make the list of Defaulter implementations available
// so projects extending GoReleaser are able to use it, namely, GoDownloader.
package defaults

import (
	"fmt"

	"goforge.dev/gpreleaser/internal/pipe/archive"
	"goforge.dev/gpreleaser/internal/pipe/artifactory"
	"goforge.dev/gpreleaser/internal/pipe/aur"
	"goforge.dev/gpreleaser/internal/pipe/aursources"
	"goforge.dev/gpreleaser/internal/pipe/blob"
	"goforge.dev/gpreleaser/internal/pipe/bluesky"
	"goforge.dev/gpreleaser/internal/pipe/brew"
	"goforge.dev/gpreleaser/internal/pipe/build"
	"goforge.dev/gpreleaser/internal/pipe/cask"
	"goforge.dev/gpreleaser/internal/pipe/changelog"
	"goforge.dev/gpreleaser/internal/pipe/checksums"
	"goforge.dev/gpreleaser/internal/pipe/chocolatey"
	"goforge.dev/gpreleaser/internal/pipe/discord"
	"goforge.dev/gpreleaser/internal/pipe/discourse"
	"goforge.dev/gpreleaser/internal/pipe/dist"
	"goforge.dev/gpreleaser/internal/pipe/docker"
	dockerv2 "goforge.dev/gpreleaser/internal/pipe/docker/v2"
	"goforge.dev/gpreleaser/internal/pipe/dockerdigest"
	"goforge.dev/gpreleaser/internal/pipe/flatpak"
	"goforge.dev/gpreleaser/internal/pipe/gomod"
	"goforge.dev/gpreleaser/internal/pipe/ko"
	"goforge.dev/gpreleaser/internal/pipe/krew"
	"goforge.dev/gpreleaser/internal/pipe/linkedin"
	"goforge.dev/gpreleaser/internal/pipe/makeself"
	"goforge.dev/gpreleaser/internal/pipe/mastodon"
	"goforge.dev/gpreleaser/internal/pipe/mattermost"
	"goforge.dev/gpreleaser/internal/pipe/mcp"
	"goforge.dev/gpreleaser/internal/pipe/milestone"
	"goforge.dev/gpreleaser/internal/pipe/nfpm"
	"goforge.dev/gpreleaser/internal/pipe/nix"
	"goforge.dev/gpreleaser/internal/pipe/notary"
	"goforge.dev/gpreleaser/internal/pipe/opencollective"
	"goforge.dev/gpreleaser/internal/pipe/project"
	"goforge.dev/gpreleaser/internal/pipe/reddit"
	"goforge.dev/gpreleaser/internal/pipe/release"
	"goforge.dev/gpreleaser/internal/pipe/sbom"
	"goforge.dev/gpreleaser/internal/pipe/scoop"
	"goforge.dev/gpreleaser/internal/pipe/sign"
	"goforge.dev/gpreleaser/internal/pipe/slack"
	"goforge.dev/gpreleaser/internal/pipe/smtp"
	"goforge.dev/gpreleaser/internal/pipe/snapcraft"
	"goforge.dev/gpreleaser/internal/pipe/snapshot"
	"goforge.dev/gpreleaser/internal/pipe/sourcearchive"
	"goforge.dev/gpreleaser/internal/pipe/srpm"
	"goforge.dev/gpreleaser/internal/pipe/teams"
	"goforge.dev/gpreleaser/internal/pipe/telegram"
	"goforge.dev/gpreleaser/internal/pipe/twitter"
	"goforge.dev/gpreleaser/internal/pipe/universalbinary"
	"goforge.dev/gpreleaser/internal/pipe/upload"
	"goforge.dev/gpreleaser/internal/pipe/upx"
	"goforge.dev/gpreleaser/internal/pipe/webhook"
	"goforge.dev/gpreleaser/internal/pipe/winget"
	"goforge.dev/gpreleaser/pkg/context"
)

// Defaulter can be implemented by a Piper to set default values for its
// configuration.
type Defaulter interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Default(ctx *context.Context) error
}

// Defaulters is the list of defaulters.
//
//nolint:gochecknoglobals
var Defaulters = []Defaulter{
	dist.Pipe{},
	snapshot.Pipe{},
	release.Pipe{},
	project.Pipe{},
	changelog.Pipe{},
	gomod.Pipe{},
	build.Pipe{},
	universalbinary.Pipe{},
	upx.Pipe{},
	sign.BinaryPipe{},
	notary.MacOS{},
	sourcearchive.Pipe{},
	archive.Pipe{},
	makeself.Pipe{},
	nfpm.Pipe{},
	srpm.Pipe{},
	snapcraft.Pipe{},
	flatpak.Pipe{},
	checksums.Pipe{},
	sign.Pipe{},
	sign.DockerPipe{},
	sbom.Pipe{},
	docker.Pipe{},
	dockerv2.Base{},
	docker.ManifestPipe{},
	dockerdigest.Pipe{},
	artifactory.Pipe{},
	blob.Pipe{},
	upload.Pipe{},
	aur.Pipe{},
	aursources.Pipe{},
	nix.Pipe{},
	winget.Pipe{},
	brew.Pipe{},
	cask.Pipe{},
	krew.Pipe{},
	ko.Pipe{},
	scoop.Pipe{},
	mcp.Pipe{},
	discord.Pipe{},
	reddit.Pipe{},
	slack.Pipe{},
	teams.Pipe{},
	twitter.Pipe{},
	smtp.Pipe{},
	mastodon.Pipe{},
	mattermost.Pipe{},
	milestone.Pipe{},
	linkedin.Pipe{},
	telegram.Pipe{},
	webhook.Pipe{},
	chocolatey.Pipe{},
	opencollective.Pipe{},
	bluesky.Pipe{},
	discourse.Pipe{},
}

// Package pipeline provides generic errors for pipes to use.
package pipeline

import (
	"fmt"

	"goforge.dev/gpreleaser/internal/pipe/announce"
	"goforge.dev/gpreleaser/internal/pipe/archive"
	"goforge.dev/gpreleaser/internal/pipe/aur"
	"goforge.dev/gpreleaser/internal/pipe/aursources"
	"goforge.dev/gpreleaser/internal/pipe/before"
	"goforge.dev/gpreleaser/internal/pipe/brew"
	"goforge.dev/gpreleaser/internal/pipe/build"
	"goforge.dev/gpreleaser/internal/pipe/cask"
	"goforge.dev/gpreleaser/internal/pipe/changelog"
	"goforge.dev/gpreleaser/internal/pipe/checksums"
	"goforge.dev/gpreleaser/internal/pipe/chocolatey"
	"goforge.dev/gpreleaser/internal/pipe/defaults"
	"goforge.dev/gpreleaser/internal/pipe/dist"
	"goforge.dev/gpreleaser/internal/pipe/docker"
	dockerv2 "goforge.dev/gpreleaser/internal/pipe/docker/v2"
	"goforge.dev/gpreleaser/internal/pipe/effectiveconfig"
	"goforge.dev/gpreleaser/internal/pipe/env"
	"goforge.dev/gpreleaser/internal/pipe/flatpak"
	"goforge.dev/gpreleaser/internal/pipe/git"
	"goforge.dev/gpreleaser/internal/pipe/gomod"
	"goforge.dev/gpreleaser/internal/pipe/ko"
	"goforge.dev/gpreleaser/internal/pipe/krew"
	"goforge.dev/gpreleaser/internal/pipe/makeself"
	"goforge.dev/gpreleaser/internal/pipe/metadata"
	"goforge.dev/gpreleaser/internal/pipe/nfpm"
	"goforge.dev/gpreleaser/internal/pipe/nix"
	"goforge.dev/gpreleaser/internal/pipe/notary"
	"goforge.dev/gpreleaser/internal/pipe/partial"
	"goforge.dev/gpreleaser/internal/pipe/prebuild"
	"goforge.dev/gpreleaser/internal/pipe/publish"
	"goforge.dev/gpreleaser/internal/pipe/reportsizes"
	"goforge.dev/gpreleaser/internal/pipe/sbom"
	"goforge.dev/gpreleaser/internal/pipe/scoop"
	"goforge.dev/gpreleaser/internal/pipe/semver"
	"goforge.dev/gpreleaser/internal/pipe/sign"
	"goforge.dev/gpreleaser/internal/pipe/snapcraft"
	"goforge.dev/gpreleaser/internal/pipe/snapshot"
	"goforge.dev/gpreleaser/internal/pipe/sourcearchive"
	"goforge.dev/gpreleaser/internal/pipe/srpm"
	"goforge.dev/gpreleaser/internal/pipe/universalbinary"
	"goforge.dev/gpreleaser/internal/pipe/upx"
	"goforge.dev/gpreleaser/internal/pipe/winget"
	"goforge.dev/gpreleaser/pkg/context"
)

// Piper defines a pipe, which can be part of a pipeline (a series of pipes).
type Piper interface {
	fmt.Stringer

	// Run the pipe
	Run(ctx *context.Context) error
}

// BuildPipeline contains all build-related pipe implementations in order.
//
//nolint:gochecknoglobals
var BuildPipeline = []Piper{
	// set default dist folder and remove it if `--clean` is set
	dist.CleanPipe{},
	// load and validate environment variables
	env.Pipe{},
	// get and validate git repo state
	git.Pipe{},
	// parse current tag to a semver
	semver.Pipe{},
	// load default configs
	defaults.Pipe{},
	// setup things for partial builds/releases
	partial.Pipe{},
	// snapshot version handling
	snapshot.Pipe{},
	// run global hooks before build
	before.Pipe{},
	// ensure ./dist exists and is empty
	dist.Pipe{},
	// setup metadata options
	metadata.Pipe{},
	// creates a metadata.json files in the dist directory
	metadata.MetaPipe{},
	// setup gomod-related stuff
	gomod.Pipe{},
	// run prebuild stuff
	prebuild.Pipe{},
	// proxy gomod if needed
	gomod.CheckGoModPipe{},
	// proxy gomod if needed
	gomod.ProxyPipe{},
	// writes the actual config (with defaults et al set) to dist
	effectiveconfig.Pipe{},
	// build
	build.Pipe{},
	// universal binary handling
	universalbinary.Pipe{},
	// upx
	upx.Pipe{},
	// sign binaries
	sign.BinaryPipe{},
	// notarize macos apps
	notary.MacOS{},
}

// BuildCmdPipeline is the pipeline run by goreleaser build.
//
//nolint:gochecknoglobals
var BuildCmdPipeline = append(
	BuildPipeline,
	reportsizes.Pipe{},
	metadata.ArtifactsPipe{},
)

// Pipeline contains all pipe implementations in order.
//
//nolint:gochecknoglobals
var Pipeline = append(
	BuildPipeline,
	// builds the release changelog
	changelog.Pipe{},
	// archive in tar.gz, zip or binary (which does no archiving at all)
	archive.Pipe{},
	// archive the source code using git-archive
	sourcearchive.Pipe{},
	// archive via fpm (deb, rpm) using "native" go impl
	nfpm.Pipe{},
	// create source RPMs
	srpm.Pipe{},
	// create makeself self-extracting archives
	makeself.Pipe{},
	// archive via snapcraft (snap)
	snapcraft.Pipe{},
	// create flatpak bundles
	flatpak.Pipe{},
	// create SBOMs of artifacts
	sbom.Pipe{},
	// checksums of the files
	checksums.Pipe{},
	// sign artifacts
	sign.Pipe{},
	// create arch linux aur pkgbuild
	aur.Pipe{},
	// create arch linux aur pkgbuild (sources)
	aursources.Pipe{},
	// create nixpkgs
	nix.New(),
	// winget installers
	winget.Pipe{},
	// homebrew formula
	brew.Pipe{},
	// homebrew cask
	cask.Pipe{},
	// krew plugins
	krew.Pipe{},
	// create scoop buckets
	scoop.Pipe{},
	// create chocolatey pkg and publish
	chocolatey.Pipe{},
	// reports artifacts sizes to the log and to artifacts.json
	reportsizes.Pipe{},
	// create and push docker images
	docker.Pipe{},
	dockerv2.Snapshot{},
	// create and push docker images using ko
	ko.Pipe{},
	// publishes artifacts
	publish.New(),
	// creates a artifacts.json files in the dist directory
	metadata.ArtifactsPipe{},
	// announce releases
	announce.Pipe{},
)

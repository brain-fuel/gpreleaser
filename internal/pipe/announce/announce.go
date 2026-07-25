// Package announce contains the announcing pipe.
package announce

import (
	"fmt"

	"goforge.dev/gpreleaser/internal/middleware/errhandler"
	"goforge.dev/gpreleaser/internal/middleware/logging"
	"goforge.dev/gpreleaser/internal/middleware/skip"
	"goforge.dev/gpreleaser/internal/pipe/bluesky"
	"goforge.dev/gpreleaser/internal/pipe/discord"
	"goforge.dev/gpreleaser/internal/pipe/discourse"
	"goforge.dev/gpreleaser/internal/pipe/linkedin"
	"goforge.dev/gpreleaser/internal/pipe/mastodon"
	"goforge.dev/gpreleaser/internal/pipe/mattermost"
	"goforge.dev/gpreleaser/internal/pipe/opencollective"
	"goforge.dev/gpreleaser/internal/pipe/reddit"
	"goforge.dev/gpreleaser/internal/pipe/slack"
	"goforge.dev/gpreleaser/internal/pipe/smtp"
	"goforge.dev/gpreleaser/internal/pipe/teams"
	"goforge.dev/gpreleaser/internal/pipe/telegram"
	"goforge.dev/gpreleaser/internal/pipe/twitter"
	"goforge.dev/gpreleaser/internal/pipe/webhook"
	"goforge.dev/gpreleaser/internal/skips"
	"goforge.dev/gpreleaser/internal/tmpl"
	"goforge.dev/gpreleaser/pkg/context"
)

// Announcer should be implemented by pipes that want to announce releases.
type Announcer interface {
	fmt.Stringer
	Announce(ctx *context.Context) error
}

//nolint:gochecknoglobals
var announcers = []Announcer{
	// XXX: keep asc sorting
	bluesky.New(),
	discord.Pipe{},
	discourse.Pipe{},
	linkedin.Pipe{},
	mastodon.Pipe{},
	mattermost.Pipe{},
	opencollective.Pipe{},
	reddit.Pipe{},
	slack.Pipe{},
	smtp.Pipe{},
	teams.Pipe{},
	telegram.Pipe{},
	twitter.Pipe{},
	webhook.Pipe{},
}

// Pipe that announces releases.
type Pipe struct{}

func (Pipe) String() string { return "announcing" }

func (Pipe) Skip(ctx *context.Context) (bool, error) {
	if skips.Any(ctx, skips.Announce) {
		return true, nil
	}
	return tmpl.New(ctx).Bool(ctx.Config.Announce.Skip)
}

// Run the pipe.
func (Pipe) Run(ctx *context.Context) error {
	memo := errhandler.Memo{}
	for _, announcer := range announcers {
		if err := skip.Maybe(
			announcer,
			logging.PadLog(announcer.String(), errhandler.Handle(announcer.Announce)),
		)(ctx); err != nil {
			memo.Memorize(fmt.Errorf("%s: %w", announcer.String(), err))
		}
	}
	return memo.Error()
}

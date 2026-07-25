# gpreleaser — parity ledger

Go+ rewrite of [`goreleaser/goreleaser`](https://github.com/goreleaser/goreleaser),
pinned to **v2.17.0** (MIT, © 2016-2026 Carlos Alexandro Becker). Module:
`goforge.dev/gpreleaser`. Wave 2, candidate 5. Binary stays `goreleaser`
(`go install goforge.dev/gpreleaser/cmd/goreleaser@v1.0.0`).

## Status

**Phase 1 — parity foundation: DONE.** goreleaser's own source (~34k LOC —
internal/pipe release stages, internal/builders, pkg, cmd) vendored under
`goforge.dev/gpreleaser` via import-surgery (the root `main.go` relocated to
`cmd/goreleaser/` so the binary stays `goreleaser`; all deps reused). Builds;
`goreleaser check` is byte-identical to upstream; `build --snapshot` works;
**the test suite has zero forge-only failures** — the only failing package
(`internal/pipe/docker`) fails identically on upstream (needs a docker daemon).
Two fixture fixes for the rename: import-surgery reverted inside `testdata/`,
and the `release` SCM `no repo` tests expect the auto-detected repo name (now
`gpreleaser`) instead of the hardcoded `goreleaser`.

**Phase 2 — Go+ authorship (next).** Convert the hand-authored source to Go+
(`.gp`): internal/{pipe, builders, pipeline}, pkg, cmd. Enum idiomata only for
genuine scalar sums with wire-format preservation.

**Phase 3 — release** to `git@github.com:brain-fuel/gpreleaser.git`.

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

**Phase 2 — Go+ authorship: DONE.** 204 hand-authored .go sources converted to Go+ (internal/pipe all ~60 stages, internal/builders, internal/pipeline, internal/*, pkg, cmd). 11 .go kept (9 //go:embed, 1 generated licenses.go, cmd/goreleaser/main.go embed). ZERO enums — artifact.Type is JSON-serialized (dist/artifacts.json), context.Action's zero value is load-bearing, config format/type/mode/provider strings are YAML-coupled. Zero goplus codegen errors under released v0.139.0. gen-check clean, deterministic; build/vet green; 120 packages pass (only internal/pipe/docker fails — needs a docker daemon, identical on upstream); `check` byte-identical across 3 configs; binary `goreleaser`.

**Phase 2 (old).** Convert the hand-authored source to Go+
(`.gp`): internal/{pipe, builders, pipeline}, pkg, cmd. Enum idiomata only for
genuine scalar sums with wire-format preservation.

**Phase 3 — release** to `git@github.com:brain-fuel/gpreleaser.git`.

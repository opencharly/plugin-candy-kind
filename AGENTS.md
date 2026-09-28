# AGENTS.md — plugin-candy-kind

Standalone plugin repo serving the `candy` box⊻layer factory kind
(`kind:candy`, compiled-in). The plugin is a Go module at
`candy/plugin-candy-kind/` (module path
`github.com/opencharly/plugin-candy-kind/candy/plugin-candy-kind`); the root
`charly.yml` only declares `discover: candy` so the repo is a project and its
candy is scanned.

Canonical files:

- `candy/plugin-candy-kind/charly.yml` — the `plugin-candy-kind:` candy entity
  (`plugin:` block, `plan:` check).
- `candy/plugin-candy-kind/plugin.go` — the kind provider (`Invoke(OpLoad)`
  echoes the host-threaded box/layer node) and `NewProvider()`/`NewMeta()`.
- `candy/plugin-candy-kind/schema/candykind.cue` — the self-contained
  `#CandyKindPlugin`.
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.
- `README.md` — user overview only; never agent guidance.

## Load these skills first (R0)

- `/charly-image:image` — the schema for `kind: candy` (charly.yml authoring —
  `base:`/`from:` makes an image; neither makes a layer). This candy carries no
  `skill:` entity of its own; the gap is recorded against
  [opencharly/opencharly#291](https://github.com/opencharly/opencharly/issues/291).
- `/charly-internals:plugin` — the plugin authoring reference: the `plugin:`
  block, the `kind` provider class, the per-plugin CUE-schema contract.
- `/charly-internals:git-workflow` — before any git/PR action.

## Build / validate / test

- `go build ./...` in `candy/plugin-candy-kind/` — compile the plugin module.
- `go test ./...` in `candy/plugin-candy-kind/` — the plugin's Go tests.
- `charly box validate` at the repo root — the structural check (the candy +
  `plugin:` block, CUE schema).
- The merge gate is the **org-wide** `charly/pr-validator` (required check
  `validate / validate`, defined in `opencharly/.github`); this repo has **no**
  per-repo candy gate.
- The changed path is exercised by EVERY box and layer decode across the
  ecosystem (every image and composed layer routes through this kind).

## Modify this repo

- Edit the `plugin-candy-kind:` candy entity, the Go source, and
  `schema/candykind.cue` **together** — the schema is the single source for the
  kind's served declaration surface.
- The rich `candy:` value is validated host-side against `#CandyValue`; the
  plugin's `OpLoad` only ECHOES the host-pre-decoded node. Keep the two shapes
  (`candy-image` → box, `candy-layer` → layer) in step with the host's fold.

## Landing

- PR-only. Every change lands through a pull request; the org-required
  `charly/pr-validator` validates the diff and body and arms native auto-merge on
  PASS. Direct pushes to `main` are blocked.
- History lives in `CHANGELOG/` (written by `tag-on-merge` at merge time); the PR
  body IS the changelog.
- The authoritative rulebook is the umbrella `AGENTS.md` in
  `opencharly/opencharly` and `charly/AGENTS.md` in the charly repo. Do not
  restate its rules here.

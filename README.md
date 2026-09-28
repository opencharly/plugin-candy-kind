# plugin-candy-kind

The `candy` box⊻layer factory kind for OpenCharly — the LAST structural kind,
relocated into a plugin (`kind:candy`).

A `candy:` node is EITHER a full IMAGE (`base:`/`from:` → a box) OR a LAYER
fragment. The rich `candy:` value is core-referencing, so the host pre-decodes
the canonical node (via `candyIsImage` + `buildCandy`) and threads the result to
this plugin, whose `OpLoad` echoes it back for the host to fold into a box or a
layer — byte-equivalent to the former in-process decode.

The plugin is **compiled-in**: a compiled-in plugin registers at `init()`, before
any load, so it has no fetch/build cycle — which dissolves the historical
bootstrap-cycle blocker (a `candy:` plugin would otherwise need to discover the
very candies it lives among).

## What it provides

| Capability | Surface |
|---|---|
| `kind:candy` | the `candy:` box⊻layer factory kind — a full image or a layer fragment |

The rich `candy:` value is validated HOST-SIDE against the kept `#CandyValue`
def, not by this plugin's served schema; the self-contained
`#CandyKindPlugin` (`schema/candykind.cue`) documents the kind's surface.

## How to use it

Compose the plugin candy in a project's `candy:` list:

```yaml
- '@github.com/opencharly/plugin-candy-kind/candy/plugin-candy-kind:<tag>'
```

Then author a `candy:` node — with `base:`/`from:` it is an image, without it a
layer.

## Layout

- `candy/plugin-candy-kind/` — the plugin module: `plugin.go`,
  `schema/candykind.cue`, `cmd/serve/main.go`.
- `charly.yml` — the root project manifest (`discover: candy`).
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.

## Related

- Owning skill: `/charly-image:image` — the schema for `kind: candy`. This candy
  carries no `skill:` entity of its own; the gap is tracked in
  [opencharly/opencharly#291](https://github.com/opencharly/opencharly/issues/291).
- `/charly-internals:plugin` — the plugin/provider model, including the `kind`
  provider class.
- [`opencharly/charly`](https://github.com/opencharly/charly) — the charly CLI.

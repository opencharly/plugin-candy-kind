// plugin-candy-kind's OWN self-contained CUE schema — the SINGLE SOURCE for this
// plugin's declaration surface. There is NO schema-less plugin: every plugin
// ships a non-empty, self-contained schema, served over Describe (the SDK splices
// `base ++ plugin` at the load gate), and this one DOCUMENTS the kind's surface.
//
// SELF-CONTAINED: it references NO base def, so it compiles STANDALONE — the exact
// property `cue exp gengotypes` needs to generate Go params, AND the property that
// lets the SDK compile it serve-side.
//
// The `candy:` value itself is RICH + core-referencing (#Candy/#Box), so it is
// validated HOST-SIDE against the kept #CandyValue def and threaded in op.Env —
// this schema documents the KIND, it does not decode the value.
#CandyKindPlugin: {
	// The structural kind word this plugin serves.
	kind: "candy"

	// The two shapes the host pre-decodes and threads in op.Env.Standalone.Shape:
	// a full IMAGE (base:/from: → uf.Box) or a LAYER fragment (→ uf.Candy).
	shapes: ["candy-image", "candy-layer"]

	// What the plugin does, in one line (the public-docs surface).
	contract: string & !=""
}

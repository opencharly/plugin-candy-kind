package candykind

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/opencharly/sdk"
	pb "github.com/opencharly/spec/proto"
	"github.com/opencharly/spec/spec"
)

// envJSON marshals the host's load envelope the way the host does before Invoke.
func envJSON(t *testing.T, env *spec.StructuralKindLoadEnv) []byte {
	t.Helper()
	out, err := json.Marshal(env)
	if err != nil {
		t.Fatalf("marshal env: %v", err)
	}
	return out
}

// TestInvokeEchoesTheOpaqueBody pins the cutover behaviour (opencharly/spec#322):
// every candy shape rides ONE bytes carrier and the plugin echoes it verbatim. It
// fails to compile against the PRE-cutover envelope — the one carrying the typed
// Box/Candy arms and no Body field — and it fails at run time if the arm stops echoing.
func TestInvokeEchoesTheOpaqueBody(t *testing.T) {
	for _, shape := range []string{"candy-image", "candy-layer"} {
		body := []byte(`{"name":"demo","kind":"candy"}`)
		env := &spec.StructuralKindLoadEnv{
			Standalone: &spec.StandaloneLoad{Shape: shape, Body: spec.RawBody(body)},
		}
		reply, err := provider{}.Invoke(context.Background(), &pb.InvokeRequest{
			Op: sdk.OpLoad, EnvJson: envJSON(t, env),
		})
		if err != nil {
			t.Fatalf("%s: Invoke: %v", shape, err)
		}
		if string(reply.ResultJson) != string(body) {
			t.Errorf("%s: echoed %q, want the body verbatim %q", shape, reply.ResultJson, body)
		}
	}
}

// TestInvokeRejectsEmptyBody pins the guard the opaque carrier needs: an empty body is
// not "an empty node", it is a host bug, and the arm must say so rather than echo "".
func TestInvokeRejectsEmptyBody(t *testing.T) {
	env := &spec.StructuralKindLoadEnv{
		Standalone: &spec.StandaloneLoad{Shape: "candy-layer"},
	}
	_, err := provider{}.Invoke(context.Background(), &pb.InvokeRequest{
		Op: sdk.OpLoad, EnvJson: envJSON(t, env),
	})
	if err == nil {
		t.Fatal("an empty body was accepted; the guard is gone")
	}
}

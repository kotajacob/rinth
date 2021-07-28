package rinth

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetMod_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mod/gvQqBUqZ", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"slug": "lithium"}`)
	})

	ctx := context.Background()
	got, _, err := client.GetMod(ctx, "gvQqBUqZ")
	if err != nil {
		t.Errorf("rinth GetMod() returned error: %v", err)
	}

	want := &Mod{Slug: String("lithium")}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("rinth GetMod() (-want +got):\n%s", diff)
	}
}

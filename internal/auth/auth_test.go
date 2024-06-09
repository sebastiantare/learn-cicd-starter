package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
  hh := http.Header{"Authorization": {"test"}}
	got, err := GetAPIKey(hh)
  print(err)
	want := []string{"a"}
	if reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

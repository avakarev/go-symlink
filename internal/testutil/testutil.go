package testutil

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// root returns path to project's root directory.
func root() string {
	_, file, _, _ := runtime.Caller(0)
	dir, err := filepath.Abs(filepath.Join(filepath.Dir(file), "..", ".."))
	if err != nil {
		panic(err)
	}
	return dir
}

func caller() string {
	_, abs, no, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	rel, _ := filepath.Rel(root(), abs)
	return fmt.Sprintf("\nFailed at %s:%d\n", rel, no)
}

// Diff fail the test if `want` differs from `got`, and prints human-readable error
func Diff(want interface{}, got interface{}, t *testing.T) {
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("%sGot unexpected result (-want +got):\n%s", caller(), diff)
	}
}

// NoErr fail the test if `err` is not nil
func NoErr(err error, t *testing.T) {
	if err != nil {
		t.Errorf("%sGot unexpected error: %v", caller(), err)
	}
}

// FixturePath returns absolute path to the given fixture
func FixturePath(names ...string) string {
	return filepath.Join(root(), "test", "fixtures", filepath.Join(names...))
}

package symlink_test

import (
	"testing"

	"github.com/avakarev/go-symlink"
	"github.com/avakarev/go-symlink/internal/testutil"
)

func TestNew(t *testing.T) {
	source := testutil.FixturePath("home", "dotfiles", "rc")
	target := testutil.FixturePath("home", ".rc")
	sym := symlink.New(source, target)

	testutil.Diff(false, sym.Source.Exists, t)
	testutil.Diff(source, sym.Source.Path, t)
	testutil.Diff(false, sym.Target.Exists, t)
	testutil.Diff(target, sym.Target.Path, t)
	testutil.Diff(false, sym.IsLinked(), t)
}

func TestReadOnSucess(t *testing.T) {
	source := testutil.FixturePath("home", "dotfiles", "rc")
	target := testutil.FixturePath("home", ".rc")
	sym := symlink.New(source, target)

	err := sym.Read()

	testutil.NoErr(err, t)
	testutil.Diff(true, sym.Source.Exists, t)
	testutil.Diff(source, sym.Source.Path, t)
	testutil.Diff(true, sym.Target.Exists, t)
	testutil.Diff(target, sym.Target.Path, t)
	testutil.Diff(true, sym.IsLinked(), t)
}

func TestReadWhenSourceNotExist(t *testing.T) {
	source := testutil.FixturePath("home", "dotfiles", "not.exist")
	target := testutil.FixturePath("home", ".rc")
	sym := symlink.New(source, target)

	err := sym.Read()
	if err == nil {
		t.Error("Got nil, but expected error")
	}

	testutil.Diff(false, symlink.IsTargetErr(err), t)
	testutil.Diff(true, symlink.IsSourceErr(err), t)
	testutil.Diff("source does not exist", err.Error(), t)
	testutil.Diff(false, sym.IsLinked(), t)
}

func TestReadWhenTargetNotExist(t *testing.T) {
	source := testutil.FixturePath("home", "dotfiles", "rc")
	target := testutil.FixturePath("home", ".not.exist")
	sym := symlink.New(source, target)

	err := sym.Read()
	if err == nil {
		t.Error("Got nil, but expected error")
	}

	testutil.Diff(false, symlink.IsSourceErr(err), t)
	testutil.Diff(true, symlink.IsTargetErr(err), t)
	testutil.Diff("target does not exist", err.Error(), t)
	testutil.Diff(false, sym.IsLinked(), t)
}

func TestReadWhenTargetNotLink(t *testing.T) {
	source := testutil.FixturePath("home", "dotfiles", "rc")
	target := testutil.FixturePath("home", ".rc.file")
	sym := symlink.New(source, target)

	err := sym.Read()
	if err == nil {
		t.Error("Got nil, but expected error")
	}

	testutil.Diff(false, symlink.IsSourceErr(err), t)
	testutil.Diff(true, symlink.IsTargetErr(err), t)
	testutil.Diff("target is not a link", err.Error(), t)
	testutil.Diff(false, sym.IsLinked(), t)
}

func TestReadWhenTargetSourceMismatch(t *testing.T) {
	source := testutil.FixturePath("home", "dotfiles", "rc")
	target := testutil.FixturePath("home", ".rc2")
	sym := symlink.New(source, target)

	err := sym.Read()
	if err == nil {
		t.Error("Got nil, but expected error")
	}

	testutil.Diff(false, symlink.IsSourceErr(err), t)
	testutil.Diff(true, symlink.IsTargetErr(err), t)
	testutil.Diff("target mismatch", err.Error(), t)
	testutil.Diff(false, sym.IsLinked(), t)
}

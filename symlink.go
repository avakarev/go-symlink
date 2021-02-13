package symlink

import (
	"errors"
	"os"
)

// Symlink represents file of directory symlink
type Symlink struct {
	Source *Source
	Target *Target
	read   bool
}

// IsLinked check whether target linked to the given source
func (sym Symlink) IsLinked() bool {
	return sym.Target.exists && sym.Target.link == sym.Source.path
}

func (sym *Symlink) Read() error {
	sym.read = true
	if err := sym.Source.Read(); err != nil {
		return err
	}
	if err := sym.Target.Read(); err != nil {
		return err
	}
	return sym.Validate()
}

// Validate check whether target linked to the given source
func (sym *Symlink) Validate() error {
	if !sym.read {
		if err := sym.Read(); err != nil {
			return err
		}
	}
	if sym.Source.exists && sym.Target.exists && sym.Target.link != sym.Source.path {
		return ErrTargetMismatch
	}
	return nil
}

// Link creates symlink
func (sym *Symlink) Link() error {
	if err := sym.Validate(); err != nil {
		if !errors.Is(err, ErrTargetNotExist) {
			return err
		}
	}

	if err := os.Symlink(sym.Source.path, sym.Target.path); err != nil {
		if errors.Is(err, os.ErrExist) {
			return ErrTargetExist
		}
		return err
	}

	sym.Target.exists = true
	sym.Target.link = sym.Source.path

	return nil
}

// Unlink deletes symlink (only target, source file/dir stays)
func (sym *Symlink) Unlink() error {
	if err := sym.Validate(); err != nil {
		return err
	}
	if err := os.Remove(sym.Target.path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrTargetNotExist
		}
		return err
	}

	sym.Target.exists = false
	sym.Target.link = ""

	return nil
}

// New returns new Symlink value
func New(s string, t string) *Symlink {
	return &Symlink{
		Source: NewSource(s),
		Target: NewTarget(t),
	}
}

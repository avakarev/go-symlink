package symlink

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	// ErrTargetNotExist represents an error when symlink's target doesn't exist
	ErrTargetNotExist = errors.New("target does not exist")

	// ErrTargetExist represents an error when symlink's target already exists
	ErrTargetExist = errors.New("target already exist")

	// ErrTargetNotLink represents an error when symlink's target is supposed to be a link but it's not
	ErrTargetNotLink = errors.New("target is not a link")

	// ErrTargetMismatch represents an error when symlink's actual target is not what is expected
	ErrTargetMismatch = errors.New("target mismatch")
)

// Target represents Symlink's target
type Target struct {
	path   string
	exists bool
	link   string
}

// Path returns target path
func (t *Target) Path() string {
	return t.path
}

// Exists returns target existance flag
func (t *Target) Exists() bool {
	return t.exists
}

// Link returns target link
func (t *Target) Link() string {
	return t.link
}

// Read reads the actual file attributes from the file system
func (t *Target) Read() error {
	t.exists = false
	t.link = ""

	source, err := filepath.EvalSymlinks(t.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrTargetNotExist
		}
		return err
	}

	t.exists = true

	if t.path == source {
		return ErrTargetNotLink
	}

	t.link = source
	return nil
}

// NewTarget returns new Target value
func NewTarget(path string) *Target {
	return &Target{
		path: path,
	}
}

// IsTargetErr checks whether error is target error
func IsTargetErr(err error) bool {
	return errors.Is(err, ErrTargetNotExist) ||
		errors.Is(err, ErrTargetExist) ||
		errors.Is(err, ErrTargetNotLink) ||
		errors.Is(err, ErrTargetMismatch)
}

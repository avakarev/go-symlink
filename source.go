package symlink

import (
	"errors"
	"os"
)

var (
	// ErrSourceNotExist represents an error when symlink's source doesn't exist
	ErrSourceNotExist = errors.New("source does not exist")
)

// Source represents Symlink's source
type Source struct {
	path   string
	exists bool
}

// Path returns source path
func (s *Source) Path() string {
	return s.path
}

// Exists returns source existence flag
func (s *Source) Exists() bool {
	return s.exists
}

// Read reads the actual file attributes from the file system
func (s *Source) Read() error {
	s.exists = false

	if _, err := os.Lstat(s.path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrSourceNotExist
		}
		return err
	}

	s.exists = true
	return nil
}

// NewSource returns new Source value
func NewSource(path string) *Source {
	return &Source{
		path: path,
	}
}

// IsSourceErr checks whether error is target error
func IsSourceErr(err error) bool {
	return errors.Is(err, ErrSourceNotExist)
}

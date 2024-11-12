package stackerr

import (
	"fmt"
)

// Frame is a single step in stack trace.
type Frame struct {
	// Path contains a file path.
	Path string
	// Line contains a line number.
	Line int
}

// String formats Frame to string.
func (f Frame) String() string {
	return fmt.Sprintf("%s:%d", f.Path, f.Line)
}

package stackerr

import (
	"fmt"
	"runtime"
)

// DefaultCap is a default cap for frames array.
// It can be changed to number of expected frames
// for purpose of performance optimisation.
var DefaultCap = 15

func trace(err error, skip int) Stakerr {
	frames := make([]Frame, 0, DefaultCap)
	for {
		_, path, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		frame := Frame{
			Line: line,
			Path: path,
		}
		frames = append(frames, frame)
		skip++
	}
	return &stakerr{
		err:    err,
		frames: frames,
	}
}

// New creates new error with stacktrace.
func New(message string) Stakerr {
	return trace(fmt.Errorf(message), 2)
}

// Wrap adds stacktrace to existing error.
func Wrap(err error, message string) Stakerr {
	if err == nil {
		return nil
	}
	stakerr, ok := err.(*stakerr)
	if ok {
		if len(message) != 0 {
			stakerr.err = fmt.Errorf("%s -> %w", message, err)
		}
		return stakerr
	}
	return trace(err, 2)
}

// StackTrace returns stack trace of an error.
// It will be empty if err is not of type Error.
func StackTrace(err error) []Frame {
	e, ok := err.(Stakerr)
	if !ok {
		return nil
	}
	return e.StackTrace()
}

// Stakerr is an error with stack trace.
type Stakerr interface {
	Error() string
	StackTrace() []Frame
	Unwrap() error
}

type stakerr struct {
	// err contains original error.
	err error
	// frames contains stack trace of an error.
	frames []Frame
}

// Error returns error message.
func (e *stakerr) Error() string {
	return e.err.Error()
}

// StackTrace returns stack trace of an error.
func (e *stakerr) StackTrace() []Frame {
	return e.frames
}

// Unwrap returns the original error.
func (e *stakerr) Unwrap() error {
	return e.err
}

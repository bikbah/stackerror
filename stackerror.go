package stackerror

import (
	"errors"
	"fmt"
	"runtime"
)

type StackError struct {
	Err   error
	Stack Stack
}

var _ error = &StackError{}

func (se StackError) Error() string {
	return fmt.Sprintf("%v\n%s", se.Err, se.Stack)
}

func (se *StackError) Unwrap() error {
	return se.Err
}

func New(err error) error {
	if err == nil {
		return nil
	}

	var serr *StackError
	if errors.As(err, &serr) {
		return err
	}

	return &StackError{
		Err:   err,
		Stack: StackTrace(3),
	}
}

// Stack represents a stacktrace as a slice of Frames.
type Stack []Frame

func StackTrace(skip int) Stack {
	pc := make([]uintptr, 50)
	n := runtime.Callers(skip, pc)
	pc = pc[:n]

	frames := runtime.CallersFrames(pc)
	stack := make(Stack, 0, n)
	for {
		frame, more := frames.Next()
		stack = append(stack, Frame{
			Filename: frame.File,
			Method:   frame.Function,
			Line:     frame.Line,
		})

		if !more {
			break
		}
	}

	return stack
}

func (s Stack) String() string {
	if len(s) == 0 {
		return ""
	}

	var res string
	for _, f := range s {
		res += f.String() + "\n"
	}

	return res
}

type Frame struct {
	Filename string
	Method   string
	Line     int
}

func (f Frame) String() string {
	return fmt.Sprintf("%s:%d[%s]", f.Filename, f.Line, f.Method)
}

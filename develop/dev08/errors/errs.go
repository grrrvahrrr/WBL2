package errors

import "errors"

// ErrNoPath is returned when 'cd' was called without a second argument.
var ErrNoPath = errors.New("path required")
var ErrTooManyArgs = errors.New("too many arguments")
var ErrNoEcho = errors.New("Nothing to echo")
var ErrNoProcessToKill = errors.New("Nothing to kill")
var ErrInvalidPid = errors.New("Invalid Pid")

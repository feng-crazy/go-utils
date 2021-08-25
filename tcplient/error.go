package tcplient

import (
	"runtime"
	"syscall"
)

// ----------------------------------------------------------------------------

// Error is the error type of the GAS package.
//
// It implements the error interface.
type Error int

const (
	// ErrMaxRetries is returned when the called function failed after the
	// maximum number of allowed tries.
	ErrMaxRetries Error = 0x01
)

// ----------------------------------------------------------------------------

// Error returns the error as a string.
func (e Error) Error() string {
	switch e {
	case ErrMaxRetries:
		return "ErrMaxRetries"
	default:
		return "unknown error"
	}
}

// ----------------------------------------------------------------------------
const (
	WSAECONNREFUSED syscall.Errno = 10061
)

func isConnResetErrorWin(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == syscall.WSAECONNRESET || se == syscall.WSAECONNABORTED
	}
	return false
}

func isConnRefusedErrorWin(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == WSAECONNREFUSED
	}
	return false
}

func isConnResetErrorNix(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == syscall.ECONNRESET || se == syscall.EPIPE
	}
	return false
}

func isConnRefusedErrorNix(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == syscall.ECONNREFUSED
	}
	return false
}

func isConnResetError(err error) bool {
	if runtime.GOOS == "windows" {
		return isConnResetErrorWin(err)
	} else {
		return isConnResetErrorNix(err)
	}
}

func isConnRefusedError(err error) bool {
	if runtime.GOOS == "windows" {
		return isConnRefusedErrorWin(err)
	} else {
		return isConnRefusedErrorNix(err)
	}
}

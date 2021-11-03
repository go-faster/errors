package errors

import (
	"sync/atomic"
)

var isTrace int64

const (
	traceEnabled  = 0
	traceDisabled = 1
)

// SetTrace sets tracing flag that controls capturing caller frames.
func SetTrace(trace bool) {
	if trace {
		atomic.StoreInt64(&isTrace, traceEnabled)
	} else {
		atomic.StoreInt64(&isTrace, traceDisabled)
	}
}

// EnableTrace enables capturing caller frames.
func EnableTrace() { SetTrace(true) }

// DisableTrace disables capturing caller frames.
func DisableTrace() { SetTrace(false) }

// Trace reports whether caller stack capture is enabled.
func Trace() bool {
	return atomic.LoadInt64(&isTrace) == traceEnabled
}

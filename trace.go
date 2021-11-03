package errors

import (
	"sync/atomic"
)

var isTrace int64

const (
	traceEnabled  int64 = 1
	traceDisabled int64 = 0
)

// SetTrace sets tracing flag that controls capturing caller frames.
func SetTrace(trace bool) {
	v := traceEnabled
	if !trace {
		v = traceDisabled
	}
	atomic.StoreInt64(&isTrace, v)
}

// EnableTrace enables capturing caller frames.
func EnableTrace() { SetTrace(true) }

// DisableTrace disables capturing caller frames.
func DisableTrace() { SetTrace(false) }

// Trace reports whether caller stack capture is enabled.
func Trace() bool {
	return atomic.LoadInt64(&isTrace) == traceEnabled
}

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/ogen-go/errors"
	"github.com/ogen-go/errors/internal"
)

type myType struct{}

func (myType) Format(s fmt.State, v rune) {
	s.Write(bytes.Repeat([]byte("Hi! "), 10))
}

func BenchmarkErrorf(b *testing.B) {
	err := errors.New("foo")
	// pi := big.NewFloat(3.14) // Something expensive.
	num := big.NewInt(5)
	args := func(a ...interface{}) []interface{} { return a }
	benchCases := []struct {
		name   string
		format string
		args   []interface{}
	}{
		{"no_format", "msg: %v", args(err)},
		{"with_format", "failed %d times: %v", args(5, err)},
		{"method: mytype", "pi: %v", args("myfile.go", myType{}, err)},
		{"method: number", "pi: %v", args("myfile.go", num, err)},
	}
	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			b.Run("ExpWithTrace", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					errors.Errorf(bc.format, bc.args...)
				}
			})
			b.Run("ExpNoTrace", func(b *testing.B) {
				internal.EnableTrace = false
				defer func() { internal.EnableTrace = true }()

				for i := 0; i < b.N; i++ {
					errors.Errorf(bc.format, bc.args...)
				}
			})
			b.Run("Core", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					fmt.Errorf(bc.format, bc.args...)
				}
			})
		})
	}
}

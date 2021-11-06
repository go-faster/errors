// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/scanner"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-faster/errors"
)

var exitCode = 0

func report(err error) {
	scanner.PrintError(os.Stderr, err)
	exitCode = 2
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gowrapper [path ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func isGoFile(f os.FileInfo) bool {
	// ignore non-Go files
	name := f.Name()
	return !f.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
}

func process(src []byte) []byte {
	result := new(bytes.Buffer)
	s := bufio.NewScanner(bytes.NewReader(src))
	var inImports bool

	const target = `"github.com/go-faster/errors"`
	replacer := strings.NewReplacer(
		`"github.com/pkg/errors"`, target,
		`"golang.org/x/xerrors"`, target,
		`"errors"`, target,
	)

	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "import") {
			inImports = true
		}

		// Multi-line import.
		if inImports && strings.HasPrefix(line, ")") {
			inImports = false
		}

		if inImports {
			line = replacer.Replace(line)
		} else {
			line = replace(line)
		}

		result.WriteString(line)
		result.WriteByte('\n')

		if strings.HasPrefix(line, "import") && !strings.Contains(line, "(") {
			// Was single line import.
			inImports = false
		}
	}
	return result.Bytes()
}

func processFile(filename string) error {
	src, err := os.ReadFile(filename) // #nosec: G304 // expected file via var
	if err != nil {
		return errors.Wrap(err, "read")
	}
	res := process(src)
	if bytes.Equal(src, res) {
		return nil
	}

	// On Windows, we need to re-set the permissions from the file. See golang/go#38225.
	var perms os.FileMode
	if fi, err := os.Stat(filename); err == nil {
		perms = fi.Mode() & os.ModePerm
	}
	if err := os.WriteFile(filename, res, perms); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func visitFile(path string, f os.FileInfo, err error) error {
	if err == nil && isGoFile(f) {
		err = processFile(path)
	}
	if err != nil {
		report(err)
	}
	return nil
}

func walkDir(path string) {
	_ = filepath.Walk(path, visitFile)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		report(errors.New("error: no path list provided"))
		usage()
	}
	for _, path := range args {
		switch dir, err := os.Stat(path); {
		case err != nil:
			report(err)
		case dir.IsDir():
			walkDir(path)
		default:
			if err := processFile(path); err != nil {
				report(err)
			}
		}
	}
	os.Exit(exitCode)
}

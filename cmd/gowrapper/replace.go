package main

import (
	"fmt"
	"strings"
)

func replace(input string) string {
	if strings.Contains(input, "xerrors.New") {
		input = strings.ReplaceAll(input, "xerrors.New", "errors.New")
	}
	if strings.Contains(input, "xerrors.Is") {
		input = strings.ReplaceAll(input, "xerrors.Is", "errors.Is")
	}
	const marker = `xerrors.Errorf("`
	if start := strings.Index(input, marker); start > 0 {
		const endMarker = `: %w", err)`
		end := strings.Index(input[start:], endMarker)
		if end > 0 {
			toReplace := input[start : start+end+len(endMarker)]
			target := fmt.Sprintf(`errors.Wrap(err, %q)`, input[start+len(marker):start+end])
			input = strings.ReplaceAll(input, toReplace, target)
		} else {
			input = strings.ReplaceAll(input, "xerrors.Errorf", "errors.Errorf")
		}
	}
	return input
}

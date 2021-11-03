package main

import (
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
	start := strings.Index(input, marker)
	if start < 0 {
		return input
	}
	const endMarker = `)`
	end := strings.LastIndex(input[start:], endMarker)
	if !strings.Contains(input, "%") {
		// No format.
		return strings.ReplaceAll(input, "xerrors.Errorf", "errors.New")
	}
	if end == 0 || !strings.Contains(input, `%w`) {
		// No wrapping, just new error.
		return strings.ReplaceAll(input, "xerrors.Errorf", "errors.Errorf")
	}

	// xerrors.Errorf("MESSAGE: %w", a1, a2, a3, a4, err)
	// fmtStart------^ fmtEnd-^    ^--firstComma   ^---lastComma
	call := input[start : start+end+len(endMarker)]
	lastComma := strings.LastIndex(call, ",")
	fmtEnd := strings.Index(call, `: %w"`)
	fmtStart := strings.Index(call, `"`)
	for _, idx := range [...]int{
		lastComma,
		fmtEnd,
		fmtStart,
	} {
		if idx > 0 {
			continue
		}
		return input
	}

	firstComma := strings.Index(call[fmtEnd:], ",")
	if firstComma < 0 {
		return input
	}

	firstComma += fmtEnd + 1
	var formatArgs string
	if firstComma < lastComma {
		formatArgs = strings.TrimSpace(call[firstComma:lastComma])
	}
	wrapTarget := strings.TrimSpace(call[lastComma+1 : len(call)-1])

	var b strings.Builder
	b.WriteString("errors.")
	if formatArgs != "" {
		b.WriteString("Wrapf")
	} else {
		b.WriteString("Wrap")
	}
	b.WriteByte('(')
	b.WriteString(wrapTarget)
	b.WriteString(`, "`)
	b.WriteString(call[fmtStart+1 : fmtEnd])
	b.WriteString(`"`)
	if formatArgs != "" {
		b.WriteString(", ")
		b.WriteString(formatArgs)
	}
	b.WriteString(")")

	return strings.ReplaceAll(input, call, b.String())
}

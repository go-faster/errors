package main

import (
	"testing"
)

func TestReplace(t *testing.T) {
	for _, tc := range []struct {
		Name, Input, Result string
	}{
		{Name: "Blank"},
		{
			Name:   "Format Wrap",
			Input:  `xerrors.Errorf("resolve '%s' reference: %w", ref, err)`,
			Result: `errors.Wrapf(err, "resolve '%s' reference", ref)`,
		},
		{
			Name:   "Wrap",
			Input:  `xerrors.Errorf("foo: %w", err)`,
			Result: `errors.Wrap(err, "foo")`,
		},
		{
			Name:   "No wrap",
			Input:  `xerrors.Errorf("bad string %d: %s", v, s)`,
			Result: `errors.Errorf("bad string %d: %s", v, s)`,
		},
		{
			Name:   "No format",
			Input:  `xerrors.Errorf("oops")`,
			Result: `errors.New("oops")`,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			if got := replace(tc.Input); got != tc.Result {
				t.Errorf("Mismatch:\n%s\n\t(got) != (result)\n%s", got, tc.Result)
			}
		})
	}
}

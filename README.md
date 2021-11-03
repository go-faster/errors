# errors [![Go Reference](https://img.shields.io/badge/go-pkg-00ADD8)](https://pkg.go.dev/github.com/ogen-go/errors#section-documentation) [![codecov](https://img.shields.io/codecov/c/github/ogen-go/errors?label=cover)](https://codecov.io/gh/ogen-go/errors)

Fork of [xerrrors](https://pkg.go.dev/golang.org/x/xerrors) with `Wrap` and `Wrapf` instead of `%w` parsing.

```
go get github.com/ogen-go/errors
```

```go
if err != nil {
	return errors.Wrap(err, "something went wrong")
}
```

## Why
* Using `Wrap` is the most explicit way to wrap errors
* Wrapping with `fmt.Errorf("foo: %w", err)` is implicit, redundant and error-prone
* Parsing `"foo: %w"` is implicit, redundant and slow
* The [pkg/errors](https://github.com/pkg/errors) and [xerrrors](https://pkg.go.dev/golang.org/x/xerrors) are not maintainted
* The [cockroachdb/errors](https://github.com/cockroachdb/errors) is too big
* The `errors` has no caller stack trace

## Don't need traces?
Call `errors.DisableTrace` or use build tag `noerrtrace`.

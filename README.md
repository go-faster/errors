# errors

Fork of [xerrrors](https://pkg.go.dev/golang.org/x/xerrors) with `Wrap` and `Wrapf` instead of `%w` parsing.

# Why
* Using `Wrap` is the most explicit way to do error wrapping
* Writing`fmt.Errorf("foo: %w", err)` implicit, redundant and error-prone
* Parsing of `"foo: %w"` is slow, implicit and redundant
* The [pkg/errors](https://github.com/pkg/errors) and [xerrrors](https://pkg.go.dev/golang.org/x/xerrors) are not maintainted
* The [cockroachdb/errors](https://github.com/cockroachdb/errors) is too big

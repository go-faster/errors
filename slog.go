package errors

import (
	"fmt"
	"log/slog"
)

func SlogString(err error) slog.Attr {
	return slog.String("err", fmt.Sprintf("%+v", err))
}

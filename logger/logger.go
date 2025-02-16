package logger

import (
	"log/slog"
	"os"
)

var Log = slog.New(slog.NewTextHandler(os.Stdout, nil))

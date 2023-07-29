package utils

import (
	"io"
	"os"

	"golang.org/x/exp/slog"
)

// functions to create, and define loggers
func NewDefaultLogger(w io.Writer) error {
	h := slog.NewTextHandler(w, nil)
	// Set new default logger
	l := slog.New(h)
	slog.SetDefault(l)
	return nil
}
func NewLogger() (*slog.Logger, error) {
	h := slog.NewTextHandler(os.Stdout, nil)
	// Return new logget
	l := slog.New(h)
	return l, nil
}

// function to write logs into a file

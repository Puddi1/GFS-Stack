package utils

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/exp/slog"
)

// Init_LoggerGFS will handle the initialization of the logger and set it as the default
func Init_LoggerGFS(logFilePath string, handlerOpts ...HandlerOptionsFunc) error {
	wr := NewWriterGFS(logFilePath)
	err := NewDefaultLogger(wr, handlerOpts...)
	if err != nil {
		return err
	}
	return nil
}

// LoggerGFS is a custom GFS logger
type WriterGFS struct {
	filePath string
}

// NewWriterGFS returns a newly initialized WriterGFS struct pointers
func NewWriterGFS(logFilePath string) *WriterGFS {
	return &WriterGFS{
		filePath: logFilePath,
	}
}

// WriterGFS implementation of the io.Write interface.
// It will log in the console and write the log on a file
func (l WriterGFS) Write(b []byte) (n int, err error) {
	// First writes out to the console
	n, errPrintLn := fmt.Println(string(b))
	if errPrintLn != nil {
		return n, errPrintLn
	}

	// Then writes down to the log file
	logFile, err := os.OpenFile(l.filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer logFile.Close()

	n, errLogFile := logFile.Write(b)
	if errLogFile != nil {
		return n, errLogFile
	}

	return n, nil
}

// HandlerOptionsFunc is the type used to define optional handler options
type HandlerOptionsFunc func(*slog.HandlerOptions)

// HandlerWithoutAddSource is the type used to set AddSource to false
func HandlerWithoutAddSource(opts *slog.HandlerOptions) {
	opts.AddSource = false
}

// HandlerWithLevel is the type used to add an optional Leveler
func HandlerWithLevel(l slog.Leveler) HandlerOptionsFunc {
	return func(opts *slog.HandlerOptions) {
		opts.Level = l
	}
}

// HandlerWithLevel is the type used to add an optional ReplaceAttr
func HandlerWithReplaceAttr(att func(groups []string, a slog.Attr) slog.Attr) HandlerOptionsFunc {
	return func(opts *slog.HandlerOptions) {
		opts.ReplaceAttr = att
	}
}

// NewDefaultLogger to create, and define a new default logger
func NewDefaultLogger(w io.Writer, handlerOpts ...HandlerOptionsFunc) error {
	opts := &slog.HandlerOptions{}
	for _, fn := range handlerOpts {
		fn(opts)
	}
	h := slog.NewTextHandler(w, opts)
	// Set new default logger
	l := slog.New(h)
	slog.SetDefault(l)
	return nil
}

// NewDefaultLogger to create, and define a new returned logger
func NewLogger(w io.Writer, handlerOpts ...HandlerOptionsFunc) (*slog.Logger, error) {
	opts := &slog.HandlerOptions{}
	for _, fn := range handlerOpts {
		fn(opts)
	}
	h := slog.NewTextHandler(w, opts)
	// Return new logget
	l := slog.New(h)
	return l, nil
}

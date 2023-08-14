package logger

import (
	"fmt"
	"io"
	"log"
)

type writer struct{}

func (w *writer) Write(p []byte) (int, error) {
	Info(string(p))
	return len(p), nil
}

func Writer() io.Writer {
	return &writer{}
}

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func Info(message string) {
	log.Printf("[INFO]: %v\r\n", message)
}

func InfoFormat(format string, v ...any) {
	Info(fmt.Sprintf(format, v...))
}

func Warn(message string) {
	log.Printf("[WARN]: %v\r\n", message)
}

func WarnFormat(format string, v ...any) {
	Warn(fmt.Sprintf(format, v...))
}

func Error(message string) {
	log.Printf("[ERROR]: %v\r\n", message)
}

func ErrorObj(err error) {
	ErrorFormat("%v", err)
}

func ErrorFormat(format string, v ...any) {
	Error(fmt.Sprintf(format, v...))
}

func WrapError(message string, err error) error {
	wrap := fmt.Errorf("failed to "+message+": %w", err) //nolint:goerr113
	ErrorObj(wrap)

	return wrap
}

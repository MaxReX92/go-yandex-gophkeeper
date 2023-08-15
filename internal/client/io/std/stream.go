package std

import (
	"bufio"
	"io"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type stream struct {
	input  *bufio.Scanner
	output *bufio.Writer
}

// NewIOStream creates a new instance of client std io command stream.
func NewIOStream(input io.Reader, output io.Writer) *stream {
	return &stream{
		input:  bufio.NewScanner(input),
		output: bufio.NewWriter(output),
	}
}

func (s *stream) Read() <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)
		for s.input.Scan() {
			result <- s.input.Text()
		}
	}()

	return result
}

func (s *stream) Write(message string) {
	_, err := s.output.WriteString(message)
	if err != nil {
		logger.ErrorFormat("failed to write output message: %v", err)
	}

	s.output.Flush()
}

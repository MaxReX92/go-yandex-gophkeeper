package io

// CommandStream represent a client command stream.
type CommandStream interface {
	// Read provide client command stream channel.
	Read() <-chan string
	// Write sends message to client io stream.
	Write(string)
}

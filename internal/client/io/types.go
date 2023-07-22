package io

type CommandStream interface {
	Read() <-chan string
	Write(string)
}

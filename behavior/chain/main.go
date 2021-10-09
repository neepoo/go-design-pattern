package main

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string2 string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)
	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (f *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)
		if f.NextChain != nil {
			f.NextChain.Next(s)
		}
		return
	}
	fmt.Printf("Finishing in second logging\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (f *WriterLogger) Next(s string) {
	if f.Writer != nil {
		f.Writer.Write([]byte("WriterLogger: " + s))
	}
	if f.NextChain != nil {
		f.NextChain.Next(s)
	}

}

type ClosureChain struct {
	NextChain ChainLogger
	Closure   func(string)
}

func (c *ClosureChain) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}
	if c.NextChain != nil {
		c.Next(s)
	}
}

func main() {

}

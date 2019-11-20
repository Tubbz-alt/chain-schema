package codegen

import (
	"bytes"
	"fmt"
)

type Writer interface {
	fmt.Stringer
	Write(x string)
}

type writer struct {
	*bytes.Buffer
}

func (w writer) Write(x string) {
	_, _ = w.WriteString(x)
}

func NewWriter() Writer {
	return &writer{&bytes.Buffer{}}
}




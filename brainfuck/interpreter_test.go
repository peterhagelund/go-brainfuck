package brainfuck

import (
	"bytes"
	"testing"
)

func TestNewInterpreter(t *testing.T) {
	interpreter := NewInterpreter(256)
	if interpreter == nil {
		t.Fatal("nil interpreter returned")
	}
}

func TestRun(t *testing.T) {
	const helloWorld = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
	interpreter := NewInterpreter(256)
	writer := bytes.NewBuffer([]byte{})
	reader := bytes.NewReader([]byte{})
	err := interpreter.Run(helloWorld, reader, writer)
	if err != nil {
		t.Fatal(err)
	}
	output := writer.String()
	if output != "Hello World!\n" {
		t.Fatalf("expected 'Hello World!' but got '%s'", output)
	}
}

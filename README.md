# Brainfuck

[Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) interpreter for Go.

## Copyright and Licensing

Copyright (c) 2024 Peter Hagelund

This software is licensed under the [MIT License](https://en.wikipedia.org/wiki/MIT_License)

See `LICENSE.txt`

## Installing

```bash
go get github.com/peterhagelund/go-brainfuck
```

## Using

```go
package main

import (
	"log"
	"os"

	"github.com/peterhagelund/go-brainfuck/brainfuck"
)

const helloWorld = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

func main() {
	interpreter := brainfuck.NewInterpreter(256)
	if err := interpreter.Run(helloWorld, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
```

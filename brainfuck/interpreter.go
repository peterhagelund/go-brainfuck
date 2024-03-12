package brainfuck

import (
	"fmt"
	"io"
)

// Interpreter defines the Brainfuck interpreter interface.
type Interpreter interface {
	// Run runs the specified code.
	Run(code string, reader io.Reader, writer io.Writer) error
}

type interpreter struct {
	ip    int
	dp    int
	size  int
	cells []byte
}

func (i *interpreter) Run(code string, reader io.Reader, writer io.Writer) error {
	i.cells = make([]byte, i.size)
	i.ip = 0
	i.dp = 0
	commands := []rune(code)
	for {
		if i.ip >= len(commands) {
			break
		}
		command := commands[i.ip]
		switch command {
		case '>':
			i.dp++
			if i.dp >= i.size {
				return fmt.Errorf("at %d: data pointer (%d) incremented beyond available cells (%d)", i.ip, i.dp, i.size)
			}
		case '<':
			i.dp--
			if i.dp < 0 {
				return fmt.Errorf("at %d: data pointer (%d) decremented below 0", i.ip, i.dp)
			}
		case '-':
			i.cells[i.dp]--
		case '+':
			i.cells[i.dp]++
		case '.':
			n, err := writer.Write(i.cells[i.dp : i.dp+1])
			if err != nil {
				return err
			}
			if n == 0 {
				return fmt.Errorf("at %d: no data written", i.ip)
			}
		case ',':
			buffer := []byte{0}
			n, err := reader.Read(buffer)
			if err != nil {
				return err
			}
			if n == 0 {
				return fmt.Errorf("at %d: no data available to read", i.ip)
			}
			i.cells[i.dp] = buffer[0]
		case '[':
			if i.cells[i.dp] == 0 {
				ip, level := 0, 0
				for ip = i.ip; ip < len(commands); ip++ {
					if commands[ip] == '[' {
						level++
					}
					if commands[ip] == ']' {
						level--
						if level == 0 {
							break
						}
					}
				}
				if commands[ip] != ']' || level != 0 {
					return fmt.Errorf("at %d: no matching ']' command found", i.ip)
				}
				i.ip = ip
			}
		case ']':
			if i.cells[i.dp] != 0 {
				ip, level := 0, 0
				for ip = i.ip; ip >= 0; ip-- {
					if commands[ip] == ']' {
						level++
					}
					if commands[ip] == '[' {
						level--
						if level == 0 {
							break
						}
					}
				}
				if commands[ip] != '[' || level != 0 {
					return fmt.Errorf("at %d: no matching '[' command found", i.ip)
				}
				i.ip = ip
			}
		}
		i.ip++
	}
	return nil
}

// NewInterpreter creates and returns a new interpreter.
func NewInterpreter(size int) Interpreter {
	return &interpreter{
		ip:   0,
		dp:   0,
		size: size,
	}
}

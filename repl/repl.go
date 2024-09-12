package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"kisumu/interpreter"
	"kisumu/lexer"
	"kisumu/object"
	"kisumu/parser"
)

const PROMPT = ">> "

// Start initializes and starts the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	var input strings.Builder

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		input.WriteString(line + "\n")

		if isCompleteStatement(line) {
			l := lexer.New(input.String())
			p := parser.New(l)

			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				printParserErrors(out, p.Errors())
				// Reset the input after reporting errors
				input.Reset()
				continue
			}

			evaluated := interpreter.Eval(program, env)
			if evaluated != nil {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}

			// Reset input after processing
			input.Reset()
		}
	}
}

// isCompleteStatement checks if the input line indicates a complete statement
func isCompleteStatement(line string) bool {
	// Heuristic: complete if it ends with a closing brace or semicolon
	return strings.HasSuffix(strings.TrimSpace(line), "}") || strings.HasSuffix(strings.TrimSpace(line), ";")
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

// Run reads input from a file if provided or starts the REPL
func Run() {
	if len(os.Args) > 1 {
		filename := os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()
		Start(file, os.Stdout)
	} else {
		Start(os.Stdin, os.Stdout)
	}
}

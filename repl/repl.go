package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kevinglasson/monkey/lexer"
	"github.com/kevinglasson/monkey/token"
)

// PROMPT is the REPL prompt constant.
const PROMPT = ">> "

// Start creates the REPL.
func Start(in io.Reader, out io.Writer) {
	// Create a scanner to read from the reader.
	scanner := bufio.NewScanner(in)

	// Loop forever.
	for {
		fmt.Fprintf(out, PROMPT)

		// Scan until the end.
		scanned := scanner.Scan()
		// Return if there was nothing scanned.
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// Loop => init tok to the first token, while not EOF and call
		// `NextToken` on each iteration
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dlmiddlecote/monkeylang/lexer"
	"github.com/dlmiddlecote/monkeylang/token"
)

// PROMPT is the prompt of the repl
const PROMPT = "🐵> "

// Start starts the REPL loop
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

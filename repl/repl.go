package repl

import (
	"bufio"
	"fmt"
	"io"
	"turkey-lang/evaluator"
	"turkey-lang/lexer"
	"turkey-lang/parser"
)

const PROMPT = ">> "

const TURKEY_ART = `
            .--.
           /} p \             /}
          '~)-) /           /' }
           ( / /          /'}.' }
            / / .-'""-.  / ' }-'}
           / (.'       \/ '.'}_.}
          |            '}   .}._}
          |     .-=-';   } ' }_.}
          \    '.-=-;'  } '.}.-}
           '.   -=-'    ;,}._.}
             '-,_  __.'' '-._}
                 '|||
                .=='=,
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	Logo(out)
	io.WriteString(out, "Gobble Gobble! Got some dressing to do!\n")
	io.WriteString(out, " parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func Logo(out io.Writer) {
	io.WriteString(out, TURKEY_ART)
	io.WriteString(out, "\n\n")
}

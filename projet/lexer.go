// lexer.go
package main

import (
	"fmt"
	"os"
	"github.com/antlr4-go/antlr/v4"
	// dépend de l'install antlr4 ? (debian 4.7.2-5)
	//"github.com/antlr4-go/antlr"  /v4 ?
)

func lexical_error(s string) {
	fmt.Println("Lexical error: " + s)
	os.Exit(2)
}

func run_lexer(file string) bool {
	log("Showing lexer output")
	// Setup the input
	is, err := antlr.NewFileStream(file)
	if err != nil { error("can't open file: " + file) }

	// Create the Lexer
	lexerErrors := &CustomErrorListener{}
	lexer := NewPcgLexer(is)
	lexer.AddErrorListener(lexerErrors)
	
	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q) position %d:%d %d-%d\n",
			lexer.SymbolicNames[t.GetTokenType()],
			t.GetText(),
			t.GetLine(), t.GetColumn(),
			t.GetStart(), t.GetStop())
	}
	return (lexerErrors.count > 0)
}

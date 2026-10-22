// parser.go
package main

import (
	"github.com/antlr4-go/antlr/v4"
	// dépend de l'install antlr4 ? (debian 4.7.2-5)
	//"github.com/antlr4-go/antlr"
)


// we add a CustomErrorListener just to count the number of errors

type CustomErrorListener struct {
    *antlr.DefaultErrorListener // Embed default which ensures we fit the interface
    count int
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
    c.count = c.count + 1
}

func run_parser(file string) (bool,bool) {
	log("Lexing and parsing")
	// Setup the input
	is, err := antlr.NewFileStream(file)
	if err != nil { error("can't open file: " + file) }

	// Create the Lexer and the stream of tokens
	lexerErrors := &CustomErrorListener{}
	l := NewPcgLexer(is)
	l.AddErrorListener(lexerErrors)
	stream := antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel)

	// Create the Parser
	parserErrors := &CustomErrorListener{}
	p := NewPcgParser(stream)
	p.AddErrorListener(parserErrors)

	// run it
	p.Prog()

	return lexerErrors.count>0, parserErrors.count>0
}

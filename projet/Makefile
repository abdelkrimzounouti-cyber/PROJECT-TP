build: pcg_lexer.go pcg_parser.go kind_string.go type_string.go
	go build

pcg_lexer.go pcg_parser.go: Pcg.g4
	antlr4 -Dlanguage=Go -package main -no-listener Pcg.g4

kind_string.go: tree.go
	stringer -type=Kind

type_string.go: typing.go
	stringer -type=Type

clean:
	rm -f *.interp *.tokens pcg_*.go *_string.go pcg

test:
	../run-test.sh Tests/*.pcg

btest:
	../run-test.sh Tests/*.pcg 2> /dev/null | grep ^--



package main

import (
	"os"
	"flag"
	"fmt"
)

var root *Tree
var verbose *bool

func main () {

	verbose = flag.Bool("v", false, "be verbose")
	showlex := flag.Bool("l", false, "show lexer output and stop")
	showparse := flag.Bool("p", false, "show parser output")
	showtree := flag.String("d", "", "dump tree created by parser in file")
	showttree := flag.String("t", "", "dump typed tree in file")
	showtypes := flag.Bool("s", false, "show symbol table after typing")
	gen   := flag.Bool("g", false, "generate code")
	inter := flag.Bool("i", false, "interpret code")
	flag.Parse()

	rest := flag.Args()
	if len(rest)<1 {
		println("Need a source file as argument!")
		os.Exit(1)
	}

	// run lexical analysis only
	if *showlex {
		if run_lexer(rest[0]) { os.Exit(2) } else { os.Exit(0) }
	}

	// run full analysis
	el, er := run_parser(rest[0])
	if el { os.Exit(2) }
	if er { os.Exit(3) }
	
	if *showparse {
		if root.s1 != nil {
			fmt.Printf("declare: %v\ninstructions: %v\n", root.s1.ss, root.s2.ss)
		} else {
			fmt.Printf("declare: %v\ninstructions: %v\n", "<none>", root.s2.ss)
		}
	}
	if *showtree != "" { tree_dump(root, false, *showtree) }

	if typing(root) {
		if *showtypes { print_st(root) }
		if *showttree != "" { tree_dump(root, true, *showttree) }
		if *gen {
			codegen(root)
		}
		if *inter {
			interpret(root)
		}
	} else {
		fmt.Println("Typing failed")
		if *showtypes { print_st(root) }
		if *showttree != "" { tree_dump(root, true, *showttree) }
		os.Exit(4)
	}
}

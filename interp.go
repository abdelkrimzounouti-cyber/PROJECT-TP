package main

import ("strconv"; "fmt"; "os")

// our language supports integer, boolean, and string values
// we'll use that as a sort of union type
type value = struct { i int; b bool; s string }

// mapping variable names to their current value
var state map[string]value

func int_expr(n *Tree) int {
	var v = interpret_expr(n)
	return v.i
}


func ierror(s string) {
	fmt.Println("Error! " + s)
	os.Exit(5)
}

func interpret_expr(n *Tree) value {
	var v value

	switch n.Label {
	case Eint: v.i, _ = strconv.Atoi(n.attr)
	default: error("interpret_expr missing case!")
	}
	return v
}

func interpret_inst(n * Tree) {
	//fmt.Printf("state: %v\n", state)
	switch n.Label {
	case Iwrite:
        fmt.Print(int_expr(n.s1))
	default: error("interpret_inst missing case!")
	}
}

// TODO: needed?
func interpret_decls(ns []*Tree) {
	for _,d := range(ns) {
		switch d.Label {
		case Dvar:
		}
	}
}

func interpret_insts(ns []*Tree) {
	log("Interpreting")
	for _,i := range(ns) {
		interpret_inst(i)
	}
}

func interpret(r *Tree) {
	state = make(map[string]value)
	if r.s1 != nil { interpret_decls(r.s1.ss) }
	interpret_insts(r.s2.ss)
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

// our language supports integer, boolean, and string values
// we'll use that as a sort of union type
type value = struct {
	i int
	b bool
	s string
}

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
	case Eint:
		x, err := strconv.Atoi(n.attr)
		if err != nil {
			ierror("Invalid integer: " + n.attr)
		}
		v.i = x
	case Ebool:
		v.b = (n.attr == "true")
	case Eadd:
		v.i = int_expr(n.s1) + int_expr(n.s2)
	case Esub:
		v.i = int_expr(n.s1) - int_expr(n.s2)
	case Emul:
		v.i = int_expr(n.s1) * int_expr(n.s2)
	case Ediv:
		right := int_expr(n.s2)
		if right == 0 {
			ierror("Division by zero")
		}
		v.i = int_expr(n.s1) / right
	case Eeq:
		v.b = (interpret_expr(n.s1) == interpret_expr(n.s2))
	case Elt:
		v.b = (int_expr(n.s1) < int_expr(n.s2))
	case Egt:
		v.b = (int_expr(n.s1) > int_expr(n.s2))
	case Ele:
		v.b = (int_expr(n.s1) <= int_expr(n.s2))
	case Ege:
		v.b = (int_expr(n.s1) >= int_expr(n.s2))
	case Ene:
		v.b = (interpret_expr(n.s1) != interpret_expr(n.s2))
	case Eand:
		v.b = (interpret_expr(n.s1).b && interpret_expr(n.s2).b)
	case Eor:
		v.b = (interpret_expr(n.s1).b || interpret_expr(n.s2).b)
	case Eneg:
		v.i = -int_expr(n.s1)
	case Enot:
		v.b = !interpret_expr(n.s1).b
	case Did:
		var ok bool
		v, ok = state[n.attr]
		if !ok {
			ierror("Undefined variable: " + n.attr)
		}
	default:
		ierror("interpret_expr missing case!")
	}
	return v
}

func interpret_inst(n *Tree) {
	//fmt.Printf("state: %v\n", state)
	switch n.Label {

	case Iblock:
		for _, inst := range n.ss {
			interpret_inst(inst)
		}
	case Iwrite:
		fmt.Print(int_expr(n.s1))
	case Iwriteln:
		fmt.Println(int_expr(n.s1))
	case Iassign:
		state[n.s1.attr] = interpret_expr(n.s2)
	case Iread:
    var input string
    _, err := fmt.Scan(&input)
    if err != nil {
        ierror("Invalid input")
    }
    x, err := strconv.ParseInt(input, 10, 32)
    if err != nil {
        ierror("Invalid input: " + input)
    }
    state[n.s1.attr] = value{i: int(x)}
	case Iif:
		if interpret_expr(n.s1).b {
			interpret_inst(n.s2)
		}
	case Iifelse:
		if interpret_expr(n.s1).b {
			interpret_inst(n.s2)
		} else {
			interpret_inst(n.s3)
		}
	case Iwhile:
		for interpret_expr(n.s1).b {
			interpret_inst(n.s2)
		}
	default:
		ierror("interpret_inst missing case!")
	}
}

// TODO: needed?
func interpret_decls(ns []*Tree) {
	for _, d := range ns {
		switch d.Label {
		case Dvar:
			state[d.s1.attr] = value{}
		case Dblock:
			interpret_decls(d.ss)
		default:
			ierror("interpret_decls missing case!")
		}
	}
}

func interpret_insts(ns []*Tree) {
	log("Interpreting")
	for _, i := range ns {
		interpret_inst(i)
	}
}

func interpret(r *Tree) {
	state = make(map[string]value)
	if r.s1 != nil {
		interpret_decls(r.s1.ss)
	}
	interpret_insts(r.s2.ss)
}

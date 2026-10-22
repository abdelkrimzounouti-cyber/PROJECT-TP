package main

import (
	"fmt"
)

// Our types

type Type int
const (
	None Type = iota
	Integer
	Void
	TypeError
)

func read_type(s string) Type {
	switch s {
	case "integer": return Integer
	}
	return TypeError
}

// Symbol table

type STEntry struct {  // we could add more fields later...
	Type Type
	addr int32
}

var st map[string]STEntry

var type_success = true

func typing(r *Tree) bool {
	log("Performing typechecking")

	st = make(map[string]STEntry)
	
	if r.s1 != nil {
		for _,d := range(r.s1.ss) {
			typecheck_decl(d)
		}
	}
	for _,i := range(r.s2.ss) {
		typecheck_inst(i)
	}
	return type_success
}

func print_st(r *Tree) {
	fmt.Printf("Symbol table: %v\n", st)
}

func type_error(m string, l int) Type {
	fmt.Printf("line %d -- type error: %s\n", l, m)
	type_success = false
	return TypeError
}

func typecheck_decl(d *Tree) {
}

func get_id_type(i *Tree) Type {
	e, ok := st[i.attr]
	if ok { return e.Type } else {
		return type_error("unknown id: " + i.attr, i.line)
	}
}

func typecheck_inst(i *Tree) Type {
	var ok bool
	var err string

	switch i.Label {
	case Iwrite:
		te := typecheck_expr(i.s1)
	        ok = te == Integer
	        err = "Write only accepts an integer expression"
	default: error("typecheck_inst missing case!")
	}
	if ok { i.t = Void; return Void
	} else { i.t = TypeError; return type_error(err, i.line) }
}

func typecheck_expr(e *Tree) Type {
	var t Type
	switch e.Label {
	case Eint: t = Integer
	default: error("typecheck_expr missing case!")
	}
	e.t = t
	return t
}

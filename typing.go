package main

import (
	"fmt"
)

// Our types

type Type int

const (
	None Type = iota
	Integer
	// j'ai ajouter cette type apres j'ai fait make build pour regenerer le fichier type_string.go
	Boolean
	Void
	TypeError
)

func read_type(s string) Type {
	switch s {
	case "integer":
		return Integer
	case "boolean":
		return Boolean
	}
	return TypeError
}

// Symbol table

type STEntry struct { // we could add more fields later...
	Type Type
	addr int32
}

var st map[string]STEntry

var type_success = true

func typing(r *Tree) bool {
	log("Performing typechecking")

	st = make(map[string]STEntry)

	if r.s1 != nil {
		for _, d := range r.s1.ss {
			typecheck_decl(d)
		}
	}
	for _, i := range r.s2.ss {
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

// cette fonction is for tommorrow not now :/
func typecheck_decl(d *Tree) {
	var err string

	if d.Label != Dvar {
		error("typecheck_decl missing case!")
	}
	t := read_type(d.s1.attr)
	if t == TypeError {
		err = "Unknow type: " + d.s1.attr
		type_error(err, d.line)
	}
	st[d.s2.attr] = STEntry{Type: t, addr: 0}
}

func get_id_type(i *Tree) Type {
	e, ok := st[i.attr]
	if ok {
		return e.Type
	} else {
		return type_error("unknown id: "+i.attr, i.line)
	}
}

// i think j'ai fini cette fonction
func typecheck_inst(i *Tree) Type {
	var ok bool
	var err string

	switch i.Label {
	case Iwrite:
		te := typecheck_expr(i.s1)
		ok = te == Integer
		err = "Write only accepts an integer expression"
	case Iwriteln:
		te := typecheck_expr(i.s1)
		ok = te == Integer
		err = "Writeln only accepts an integer expression"
	case Iread:
		te := get_id_type(i.s1)
		ok = te == Integer
		err = "Read only accepts an integer variable"
	case Iassign:
		t1 := get_id_type(i.s1)
		t2 := typecheck_expr(i.s2)
		ok = t1 == t2 && (t1 == Integer || t1 == Boolean)
		err = "Assignment only accepts an integer or boolean variable and an integer or boolean expression"
	case Iif:
		t1 := typecheck_expr(i.s1)
		ok = t1 == Boolean
		err = "If only accepts a boolean expression"
		t2 := typecheck_inst(i.s2)
		if t2 == TypeError {
			ok = false
		}
	case Iifelse:
		t1 := typecheck_expr(i.s1)
		ok = t1 == Boolean
		err = "If-Else only accepts a boolean expression"
		t2 := typecheck_inst(i.s2)
		t3 := typecheck_inst(i.s3)
		if t2 == TypeError || t3 == TypeError {
			ok = false
		}
	case Iwhile:
		t1 := typecheck_expr(i.s1)
		ok = t1 == Boolean
		err = "While only accepts a boolean expression"
		t2 := typecheck_inst(i.s2)
		if t2 == TypeError {
			ok = false
		}
	// il reste encore le cas de Iblock mais Ana 3yite db

	default:
		error("typecheck_inst missing case!")
	}
	if ok {
		i.t = Void
		return Void
	} else {
		i.t = TypeError
		return type_error(err, i.line)
	}
}

// cette fonction est nickel
func typecheck_expr(e *Tree) Type {
	var t Type
	switch e.Label {
	case Eint:
		t = Integer
	case Ebool:
		t = Boolean
	case Eadd, Esub, Emul, Ediv:
		t1 := typecheck_expr(e.s1)
		t2 := typecheck_expr(e.s2)
		if t1 == Integer && t2 == Integer {
			t = Integer
		} else {
			t = TypeError
		}
	case Elt, Ele, Egt, Ege:
		t1 := typecheck_expr(e.s1)
		t2 := typecheck_expr(e.s2)
		if t1 == Integer && t2 == Integer {
			t = Boolean
		} else {
			t = TypeError
		}
	case Eeq, Ene:
		t1 := typecheck_expr(e.s1)
		t2 := typecheck_expr(e.s2)
		if t1 == t2 && (t1 == Integer || t1 == Boolean) {
			t = Boolean
		} else {
			t = TypeError
		}
	case Eand, Eor:
		t1 := typecheck_expr(e.s1)
		t2 := typecheck_expr(e.s2)
		if t1 == Boolean && t2 == Boolean {
			t = Boolean
		} else {
			t = TypeError
		}
	case Eneg:
		t1 := typecheck_expr(e.s1)
		if t1 == Integer {
			t = Integer
		} else {
			t = TypeError
		}
	case Enot:
		t1 := typecheck_expr(e.s1)
		if t1 == Boolean {
			t = Boolean
		} else {
			t = TypeError
		}
	default:
		error("typecheck_expr missing case!")
	}
	e.t = t
	return t
}

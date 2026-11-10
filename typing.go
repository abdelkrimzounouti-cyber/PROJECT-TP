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
	type_success = true
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

// j'ai changer juste l'ordre des enfants dans l'arbre :/
func typecheck_decl(d *Tree) {
	var err string

	if d.Label != Dvar {
		error("typecheck_decl missing case!")
	}
	t := read_type(d.s2.attr)
	if t == TypeError {
		err = "Unknow type: " + d.s2.attr
		type_error(err, d.line)
	}
	d.t = t
    d.s1.t = t
	st[d.s1.attr] = STEntry{Type: t, addr: 0}
}

func get_id_type(i *Tree) Type {
    e, ok := st[i.attr]
	if ok {
        i.t = e.Type
        return e.Type
    } else {
        i.t = TypeError
        return type_error("unknown id: "+i.attr, i.line)
    }
}

// je pense que j'ai fini cette fonction :)
func typecheck_inst(i *Tree) Type {
    var ok bool
    var err string

    switch i.Label {
    case Iwrite:
        te := typecheck_expr(i.s1)
        if te == TypeError {
            i.t = TypeError
            return TypeError
        }
        ok = te == Integer
        err = "Write only accepts an integer expression"
	case Iwriteln:
        te := typecheck_expr(i.s1)
        if te == TypeError {
            i.t = TypeError
            return TypeError
        }
        ok = te == Integer
        err = "Writeln only accepts an integer expression"
	case Iread:
        te := get_id_type(i.s1)
        if te == TypeError {
            i.t = TypeError
            return TypeError
        }
        ok = te == Integer
        err = "Read only accepts an integer variable"
	case Iassign:
        t1 := get_id_type(i.s1)
        t2 := typecheck_expr(i.s2)
		if t1 == TypeError || t2 == TypeError {
            i.t = TypeError
            return TypeError
        }
		ok = t1 == t2 && (t1 == Integer || t1 == Boolean)
        err = "Assignment only accepts an integer or boolean variable and an integer or boolean expression"
	case Iif:
        t1 := typecheck_expr(i.s1)
        t2 := typecheck_inst(i.s2)
		if t1 == TypeError || t2 == TypeError {
            i.t = TypeError
            return TypeError
        }
		ok = t1 == Boolean
        err = "If only accepts a boolean expression"
	case Iifelse:
        t1 := typecheck_expr(i.s1)
        t2 := typecheck_inst(i.s2)
        t3 := typecheck_inst(i.s3)
		if t1 == TypeError || t2 == TypeError || t3 == TypeError {
            i.t = TypeError
            return TypeError
        }
		ok = t1 == Boolean
        err = "If-Else only accepts a boolean expression"
	case Iwhile:
        t1 := typecheck_expr(i.s1)
        t2 := typecheck_inst(i.s2)
		if t1 == TypeError || t2 == TypeError {
            i.t = TypeError
            return TypeError
        }
		ok = t1 == Boolean
        err = "While only accepts a boolean expression"
	case Iblock:
    ok = true
    for _, inst := range i.ss {
        if typecheck_inst(inst) == TypeError {
            ok = false
        }
    }
	if !ok {
        i.t = TypeError
        return TypeError
    }
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


// cette fonction est bon :)
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
		if (t1 == TypeError || t2 == TypeError ) {
   			 t = TypeError
		} else if (t1 == Integer && t2 == Integer ){
			t = Integer
		} else {
            t = type_error("Arithmetic operators require integer operands", e.line)
        }
	case Elt, Ele, Egt, Ege:
		t1 := typecheck_expr(e.s1)
		t2 := typecheck_expr(e.s2)
		
		if t1 == TypeError || t2 == TypeError {
    		t = TypeError
		} else if t1 == Integer && t2 == Integer {
			t = Boolean
		} else {
            t = type_error("Comparison operators require integer operands", e.line)
        }
	case Eeq, Ene:
		t1 := typecheck_expr(e.s1)
		t2 := typecheck_expr(e.s2)
		if (t1 == TypeError || t2 == TypeError ) {
   			 t = TypeError
		} else if (t1 == t2 && (t1 == Integer || t1 == Boolean)) {
			t = Boolean
		} else {
            t = type_error("Equality operators require operands of the same type", e.line)
        }
	case Eand, Eor:
		t1 := typecheck_expr(e.s1)
		t2 := typecheck_expr(e.s2)
		if (t1 == TypeError || t2 == TypeError ) {
          t = TypeError
		} else if t1 == Boolean && t2 == Boolean {
			t = Boolean
		} else {
            t = type_error("Logical operators require boolean operands", e.line)
        }
	case Eneg:
		t1 := typecheck_expr(e.s1)
		if (t1 == TypeError) {
			t = TypeError
		} else if t1 == Integer {
			t = Integer
		} else {
            t = type_error("Unary minus requires an integer operand", e.line)
        }
	case Enot:
		t1 := typecheck_expr(e.s1)
		if (t1 == TypeError) {
			t = TypeError
		} else if t1 == Boolean {
			t = Boolean
		} else {
            t = type_error("Logical NOT requires a boolean operand", e.line)
        }
	case Did:
    	t = get_id_type(e)
	default:
		error("typecheck_expr missing case!")
	}
	e.t = t
	return t
}

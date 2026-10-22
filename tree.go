package main

import (
	"fmt"
	"os"
)

// Labels for Tree nodes
type Kind int

const (
	Pgm Kind = iota
	Dblock
	Dvar
	Did

	Eint
	Ebool
	Eadd
	Esub
	Emul
	Ediv
	Elt
	Ele
	Egt
	Ege
	Eeq
	Ene
	Eand
	Eor
	Eneg
	Enot

	Iblock
	Iassign
	Iwrite
	Iwriteln
	Iread
	Iifelse
	Iif
	Iwhile
)

type Tree struct {
	Label      Kind
	s1, s2, s3 *Tree   // if a node has <= 3 sons, use s1, s2 and s3
	ss         []*Tree // if it has more use ss (sons)
	t          Type    // will be set by typing
	attr       string  // attribute
	line       int     // source line
}

// Constructors for node with 0 (leaf), 1, 2, 3 or more sons
// p is the source position (line number)

func Leaf(k Kind, a string, p int) *Tree {
	return &Tree{k, nil, nil, nil, nil, 0, a, p}
}

func Node1(k Kind, s *Tree, p int) *Tree {
	n := new(Tree)
	n.s1 = s
	n.Label = k
	n.line = p
	return n
}

func Node2(k Kind, l, r *Tree, p int) *Tree {
	n := new(Tree)
	n.s1 = l
	n.s2 = r
	n.Label = k
	n.line = p
	return n
}

func Node3(k Kind, s1, s2, s3 *Tree, p int) *Tree {
	n := new(Tree)
	n.s1 = s1
	n.s2 = s2
	n.s3 = s3
	n.Label = k
	n.line = p
	return n
}

func NodeN(k Kind, ss []*Tree, p int) *Tree {
	n := new(Tree)
	n.ss = ss
	n.Label = k
	n.line = p
	return n
}

// make Tree an instance of the Stringer interface

func (t Tree) String() string {
	var r string
	s1 := t.s1
	s2 := t.s2
	s3 := t.s3
	n := t.Label
	switch {
	case s3 != nil:
		r = fmt.Sprintf("(%v %v %v %s)", n, s1.String(), s2.String(), s3.String())
	case s2 != nil:
		r = fmt.Sprintf("(%v %v %v)", n, s1.String(), s2.String())
	case s1 != nil:
		r = fmt.Sprintf("%s %s", n, s1.String())
	case t.ss != nil:
		r = fmt.Sprintf("%s %v", n, t.ss)
	case true:
		r = fmt.Sprintf("%s [%s]", n.String(), t.attr)
	}
	return r
}

func id_of_assign(n *Tree) string {
	return (n.s1.attr)
}

// dump tree as indented text file

var dump_indent uint = 0

func tree_dump1(f *os.File, t *Tree, showt bool) {
	var s string

	if t == nil {
		return
	}

	// warning: root node can have s1 nil (no decls) and s2 w/ insts
	if t.s1 == nil && t.s2 == nil && t.ss == nil {
		s = t.attr
	} else {
		s = fmt.Sprintf("%v", t.Label)
	}
	fmt.Fprintf(f, "%[1]*s%s\n", dump_indent, "", s)
	fmt.Fprintf(f, "%[1]*s:l %v\n", dump_indent, "", t.line)
	if showt {
		fmt.Fprintf(f, "%[1]*s:t %v\n", dump_indent, "", t.t)
	}

	dump_indent += 2
	if t.ss == nil {
		tree_dump1(f, t.s1, showt)
		tree_dump1(f, t.s2, showt)
		tree_dump1(f, t.s3, showt)
	} else {
		for _, n := range t.ss {
			tree_dump1(f, n, showt)
		}
	}
	dump_indent -= 2
}

func tree_dump(r *Tree, showt bool, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	logs("Dumping tree in file " + filename)
	tree_dump1(f, r, showt)
	f.Close()
}

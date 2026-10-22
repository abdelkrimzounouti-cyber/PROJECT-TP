package main

import ("strconv"; "fmt"; "os";
	. "codeberg.org/deleuzec/stackm")

func write_int32(addr int, v int32) {
        prog[addr]   = Inst(v>>24)
	prog[addr+1] = Inst(v & 0xFF0000 >> 16)
	prog[addr+2] = Inst(v & 0xFF00 >> 8)
	prog[addr+3] = Inst(v & 0xFF)
}

var prog [4096]Inst
var pc = 0


var addr int32 = 4096  // for variables

func add(i Inst) {
	//fmt.Printf(" %04x %v\n", pc, i)
	prog[pc] = i
	pc++
}

func add_push(v int32) {
	//fmt.Printf(" %04x Push %d\n", pc, v)
	prog[pc] = Push
	write_int32(pc+1, v)
	pc += 5
}

func save(file string, n int) {
	bin := make([]byte, n)
	for i:=0;i<n;i++ {
		bin[i] = byte(prog[i])
	}
	
	err := os.WriteFile(file, bin, 0666)
	if err != nil {
		fmt.Println("error saving file")
	}
	fmt.Printf("Saved %d bytes to %v\n", n, file)
}


func codegen_expr(e *Tree) {
	switch e.Label {
	case Eint:
		v, _ := strconv.Atoi(e.attr)
		add_push(int32(v))
	default: error("codegen_expr missing case!")
	}
}

func codegen_inst(i *Tree) {
	switch i.Label {
	case Iwrite:
		codegen_expr(i.s1)
		add(WriteI)
	default: error("codegen_inst missing case!")
	}
}

func codegen_decl(d *Tree) {
}

func codegen(r *Tree) {
	log("Generating stack machine code")
	if r.s1 != nil {
		for _,d := range(r.s1.ss) {
			codegen_decl(d)
		}
	}
	logs(fmt.Sprintf("Symbol table: %v\n", st))

	for _,i := range(r.s2.ss) {
		codegen_inst(i)
	}
	add(Halt)
	if *verbose { Dump(prog[:pc]) }
	save("/tmp/output.b", pc)

}

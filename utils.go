package main

import ("fmt" ; "os")

func log(s string) {
	if *verbose { fmt.Println("\n-- " + s + "...") }
}

func logs(s string) {
	if *verbose { fmt.Println(s) }
}

func error(s string) {
	fmt.Println(s)
	os.Exit(1)
}

func reverse (s []*Tree) []*Tree {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

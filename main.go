package main

import "fmt"

func main() {
	n := 6
	m := n ^ 1
	fmt.Printf("result: %+v \n", m)

	// перевести n в двоичную систему
	dn := fmt.Sprintf("%b", n)
	fmt.Printf("result: n=%s m=%s \n", dn, fmt.Sprintf("%b", m))
}

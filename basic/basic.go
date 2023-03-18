package main

import "fmt"

func enum() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
	return

}

func main() {
	enum()
}

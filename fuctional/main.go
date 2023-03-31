package main

import (
	"bufio"
	"fmt"
	"io"
	"learngo/fuctional/fib"
)

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	g := fib.Fibonacci()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(g())
	//}
	printFileContents(g)
}

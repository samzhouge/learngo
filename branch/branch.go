package main

import (
	"fmt"
	"os"
)

func grade(score int) string {
	var g string
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score"))
	case score < 60:
		g = "F"
	case score < 70:
		g = "E"
	case score < 80:
		g = "D"
	case score < 90:
		g = "C"
	case score < 100:
		g = "B"
	case score == 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(70),
		grade(80),
		grade(90),
		grade(100),
		//grade(101),
	)
}

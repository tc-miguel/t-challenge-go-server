package main

import "fmt"

func vals() (int, int) {
	return 3, 4
}

func main12() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

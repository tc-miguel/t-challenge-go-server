package main

import (
	"fmt"
	"slices"
)

func main16() {
	var s []string
	fmt.Println("Uninit: ", s, s == nil, len(s) == 0)
	//Console result [] true true

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set values for 0,1,2 index:", s)
	fmt.Println("get position 2:", s[2])
	fmt.Println("len:", len(s))

	//Add elements to slide
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//Copy array into other
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy c array:", s)

	//Copy specific range
	l := s[2:5]
	fmt.Println("Array specific copy range:", l)

	l = s[2:]
	fmt.Println("Array range from position to end:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("T array definition:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t equals t2")
	}

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}

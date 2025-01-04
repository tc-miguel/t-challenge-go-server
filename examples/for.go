package main

import "fmt"

func main6() {
	/*i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}*/

	/*for j := 0; j < 3; j++ {
		fmt.Println(j)
	}*/

	/*for i := range 3 {
		fmt.Println("range", i)
	}*/

	/*for {
		fmt.Println("Loop")
		break
	}*/

	for n := range 6 {
		fmt.Println("Value to process: ", n)
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

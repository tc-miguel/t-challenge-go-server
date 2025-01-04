package main

import "fmt"

func main1() {

	//Basic array
	var a [5]int
	fmt.Println("empty array:", a)

	a[4] = 100
	fmt.Println("set 100 by 4 index :", a)
	fmt.Println("get index 4:", a[4])

	fmt.Println("Length (len function):", len(a))

	//Set array values for initialization with specific length
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//Set array values for initialization without specific length
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//Initiliazation
	//First position = 100
	//After 3 position 400 and 500
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	//2D array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}

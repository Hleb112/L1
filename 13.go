package main

import "fmt"

func main() {
	a := 4
	b := 6
	a = a + b //в а сумма. в б старое знач
	b = a - b //10 - 6 = 4
	a = a - b //10 - 4 = 6
	fmt.Println("a = ", a, "b = ", b)

	x := 10 //спопсоб 2
	z := 0
	fmt.Println(x, z)
	z, x = x, z
	fmt.Println(x, z)
}

package main

import "fmt"

func F(a int) int {

	if a < 2 {
		return a
	}
	fmt.Println("Current value:", a)
	return F(a-2) + F(a-1)

}

func main() {
	fmt.Println(F(10))
}

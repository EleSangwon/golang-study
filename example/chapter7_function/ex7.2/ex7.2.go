package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, " Score is :", avg)
}
func main() {
	PrintAvgScore("A", 80, 74, 95)
	PrintAvgScore("B", 90, 64, 85)
	PrintAvgScore("C", 70, 54, 75)

}

package main

import "fmt"

func main() {
	fmt.Println(mySqrt(20))
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	length := x/2
	for i := 1; i<= length;i++ {
		if i*i <= x && (i+1)*(i+1) > x{
			return i
		}
	}
	return 1
}

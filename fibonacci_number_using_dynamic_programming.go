//fibonacci number in Golang using dynamic programming
package main

import "fmt"

func fibonacci(n int) int {

	first := 0
	second := 1
	temp := 0
	if n == 0 {

		return first
	} else if n == 1 {

		return second
	}
	for i := 2; i <= n; i++ {

		temp = first + second
		first = second
		second = temp
	}
	return temp
}

func main() {
	fmt.Println(fibonacci(20))
}

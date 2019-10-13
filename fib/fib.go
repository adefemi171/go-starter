package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
func main() {
	fmt.Println("Program to Calculate Fibonnaci of Value ")
	fmt.Print("Enter Your End Value ")
	var userinput int
	fmt.Scanln(&userinput)
	var useranswer int
	useranswer = fib(userinput)
	fmt.Println("Fib of ", userinput, " is ", useranswer)
}

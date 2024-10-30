package utils

import "fmt"

func PrintGreeting(message string) {
	fmt.Println(message)
}

func CountNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("Count: %d\n", i)
	}
}

var globalVar = "I am a global variable"

func DemonstrateScope() {
	localVar := "I am a local variable"
	fmt.Println(globalVar)
	fmt.Println(localVar)
}
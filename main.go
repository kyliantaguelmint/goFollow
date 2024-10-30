package main

import (
	"fmt"
	"myproject/models"
	"myproject/utils"
	"time"
)

func main() {
	personMap := make(map[string]models.Person)

	personMap["john_doe"] = models.Person{Name: "John Doe", Age: 30}

	if person, exists := personMap["john_doe"]; exists {
		fmt.Printf("Found: %s, Age: %d\n", person.Name, person.Age)
	} else {
		fmt.Println("Person not found.")
	}

	utils.PrintGreeting("Welcome to the Go Project!")

	go utils.CountNumbers(5)

	time.Sleep(2 * time.Second)

	utils.DemonstrateScope()
}
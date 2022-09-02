package main

import "fmt"

func main() {
	var names string = "Hi Akash"
	newName := "Akash Sengar"
	namesArray := []string{"Akash", "debjit"}
	for name := range namesArray {
		fmt.Println(name)
	}
	fmt.Println(newName)
	fmt.Println(names)
	fmt.Println(addNumber(4, 3))
	namesArray = append(namesArray, "pRASHANT")
	fmt.Println(namesArray)
}

func addNumber(a int, b int) int {
	return a + b
}

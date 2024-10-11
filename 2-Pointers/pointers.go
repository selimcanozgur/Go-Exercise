package main

import "fmt"

func main() {
	age := 26

	var agePointer *int

	agePointer = &age

	fmt.Println("Age: ",*agePointer)
	getAdultYear(agePointer)
	fmt.Println(age)

}

func getAdultYear (age *int)  {
	*age = *age - 12
}
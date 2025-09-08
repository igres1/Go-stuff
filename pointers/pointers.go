package main

import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age
	fmt.Println("Age:", *agePointer)
	editAgeAdultYears(agePointer)
	fmt.Println(age)
}
func editAgeAdultYears(age *int) {
	//	return *age - 18
	*age = *age - 18
}

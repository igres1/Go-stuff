package main

import "fmt"

func main() {
	age := 32
	var agePointer *int              // the zero value is nil for pointers
	agePointer = &age                // here we pass the mem dir of age to the pointer
	fmt.Println("Age:", *agePointer) //here we desreference the pointer to read the value contained in the mem dir
	editAgeAdultYears(agePointer)
	fmt.Println(age)
}
func editAgeAdultYears(age *int) {
	//	return *age - 18
	*age = *age - 18
}
<<<<<<< HEAD
   

me llamo pascual es un 
=======

/*NOTE: as a great reminder , a pointer stores the BASE memory direction , thats why
we use a type for the pointer ( for example int) as the compiler must now
how much offset it has to read to read the full variable in the contiguous memory*/
>>>>>>> 7a9d8ab8103a6b182e3b1d93e1b370601a59c5cf

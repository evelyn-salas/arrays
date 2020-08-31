package main

import (
	"fmt"
)

func main() {
	//how to make an array
	grades := [3]int{97, 85, 93} //you can indicate the type here
	fmt.Printf("Grades: %v\n", grades)

	//if size is unknown
	//grades := [...]int {97, 85, 93}

	//how to declare an empty array
	var students [5]string
	fmt.Printf("Students: %v\n", students)

	//append and array value
	students[0] = "Evelyn" //pointer starts at zero, just like iota
	fmt.Printf("Students: %v\n", students)

	//call a value withinh array
	fmt.Printf("Student #1: %v\n", students[0])

	//figure out array length
	fmt.Printf("Number of students: %v\n", len(students))

	//arrays within arrays (good for linear algerbra)
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0},
		[3]int{0, 0, 1}}
	fmt.Println(identityMatrix)
	//or (cleaner)
	//var identityMatrix [3][3]int
	//identityMatrix[0] = [3]int{1, 0, 0}
	//identityMatrix[1] = [3]int{0, 1, 0}
	//identityMatrix[2] = [3]int{0, 0, 1}
	//fmt.Println(identityMatrix) //comment out rest to test

	//how go differs, appending arrays
	a := [...]int{1, 2, 3}
	b := a   //b is a copy of a
	b[1] = 5 // this will only change b
	fmt.Println(a)
	fmt.Println(b)

	//how to append both arrays
	x := [...]int{1, 2, 3}
	y := &x //y is pointing at the same data as a
	y[1] = 5
	fmt.Println("this is version 2:")
	fmt.Println(x)
	fmt.Println(y)

}

package main

import (
	"fmt"
)

func main() {
	//arrays//
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

	//slices//
	i := []int{1, 2, 3}
	j := i //this DOES change underlying array
	j[1] = 5
	fmt.Println(i)
	fmt.Println(j)
	fmt.Printf("Length: %v\n", len(a))   //size of slice
	fmt.Printf("Capacity: %v\n", cap(a)) //size of underlying array

	//taking a slice from a slice
	//all of these point to the same underlying array
	q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//q := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //could also be an array
	w := q[:]   //slice of all elements
	e := q[3:]  //slice from 4th element to end
	r := q[:6]  //slice from first 6 elements
	t := q[3:6] //slice from the 4th, 5th, and 6th elements
	q[5] = 42   //the 6th element will change for all slices
	fmt.Println(q)
	fmt.Println(w)
	fmt.Println(e)
	fmt.Println(r)
	fmt.Println(t)

	//declaring slice using make function
	//the slice has an underlying array and they don't have to be equal
	m := make([]int, 0, 3) //the last argument can be left out
	fmt.Println()
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))
	fmt.Printf("Capacity: %v\n", cap(m))

	//adding an element to the slice
	m = append(m, 1, 2, 3, 4, 5)
	fmt.Println()
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))   //length increases
	fmt.Printf("Capacity: %v\n", cap(m)) //capacity increases nonlinearly
	//it is best to get it right in make

	//concatinate slices
	m = append(m, []int{6, 7, 8, 9, 10}...)
	fmt.Println()
	fmt.Println(m)
	fmt.Printf("Length: %v\n", len(m))
	fmt.Printf("Capacity: %v\n", cap(m))

	//trimming slices
	n := m[1:] //trim the start
	fmt.Println()
	fmt.Println(n)
	l := m[:len(m)-1] //trim the end
	fmt.Println(l)
	o := append(m[:2], m[3:]...)
	fmt.Println(o) //now 3 is gone, edits original array

}

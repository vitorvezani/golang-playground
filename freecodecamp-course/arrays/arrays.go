package arrays

import "fmt"

/*
  Arrays:
  - collections of items with the same type
  - fixed size
  - 3 different sintaxes
  - zero based
  - len() and cap()

  Slices:
  - backed by array
  - initialization:
 		- by slice
   	- literal style
		- via make([]int, len, cap) function
  - zero based
  - len() and cap()
	- append
	- reference based
*/

func init() {
	fmt.Println("Init arrays file executed")
}

func Arrays() {
	slicesVsArrays()
	extractingSlices()
	usingMakeFunction()
}

func slicesVsArrays() {
	studants := [...]int{1, 2, 3} // arrays
	grades := []int{1, 2, 3}      // slice
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Studants: %v\n", studants)

	fmt.Printf("Grades length: %v\n", len(grades))
	fmt.Printf("Grades capacity: %v\n", cap(grades))
}

func extractingSlices() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	f := a[3:6]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
}

func usingMakeFunction() {
	// 2 arguments
	a := make([]int, 3)
	fmt.Println(a)
	fmt.Println("len: ", len(a))
	fmt.Println("len: ", cap(a))

	// 3 arguments
	b := make([]int, 3, 100)
	fmt.Println(b)
	fmt.Println("len: ", len(b))
	fmt.Println("len: ", cap(b))

	// appending
	c := []int{}
	c = append(c, 1, 2, 3, 4)
	c = append(c, []int{1, 2, 3, 4}...)

	fmt.Println(c)
	fmt.Println("len: ", len(c))
	fmt.Println("len: ", cap(c))
}

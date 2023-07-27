package main

import "fmt"

func main() {
	fmt.Println("Slices in Go")

	var s []string
	fmt.Println("Uninitilized Slice: ", s, s == nil, len(s))

	// Making a string with predefined length, it increases dynamically
	s = make([]string, 3)
	fmt.Println("Empty string: ", s, "Len: ", len(s), "capacity: ", cap(s))

	// We can assign elements just like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// Like vectors in C++ you can push/append an elements to the array, this returns a new array and must be stored in an array
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Append Function:", s)

	// Copying a slice into another slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copied to c: ", c)

	// Sicing in slices
	name := "atharva"
	// fmt.Printf("%c", name[0])
	slice1 := name[1:5] // Slices including 1st element and excluding 5th element
	fmt.Printf("Slice 1: %s \n", slice1)
	slice2 := name[2:] // Slices including 2 to the last element
	fmt.Printf("Slice 2: %s \n", slice2)
	slice3 := name[:4] // Slices started from the first element to 4th (not including)
	fmt.Printf("Slice 3: %s \n", slice3)

	// Two dimensional slices
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

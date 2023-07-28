package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func divide(x, y int) (float64, error) {
	if y == 0 {
		return float64(0), errors.New("Cannot divide by Zero")
	} else {
		return float64(x / y), nil
	}
}

func main() {
	fmt.Printf("Error Handling in Go\n")

	file, fileErr := ioutil.ReadFile("file.txt")
	// Since I dont have a file name file.txt i am expecting an error
	if fileErr != nil {
		fmt.Println("Error reading file: ", fileErr)
	} else {
		fmt.Println("File says: ", string(file))
	}

	ans, err := divide(10, 2)

	if err != nil {
		fmt.Println("There is an Error: ", err)
	} else {
		fmt.Println("Answer is: ", float64(ans))
	}

}

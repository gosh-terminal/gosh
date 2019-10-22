package main
import "fmt"

func linOfSym() {
	var a float64
	var b float64
	var los float64
	fmt.Print("What is a?: ")
	fmt.Scan(&a)
	fmt.Print("What is b?: ")
	fmt.Scan(&b)
	los = (-1 * b) / (2 * a)
	fmt.Println("Line of symmetry =", los)
}

package main
import "fmt"
import "math"
func pythagTheor() {
	var a float64
	var b float64
	var h float64
	fmt.Print("What is a?: ")
	fmt.Scan(&a)
	fmt.Print("What is b?: ")
	fmt.Scan(&b)
	h = math.Pow(a,2) + math.Pow(b,2)
	fmt.Println("hypot =",h)
}

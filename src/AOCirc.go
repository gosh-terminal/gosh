package main
import "fmt"
import "math"
func AOCirc() {
	var r float64
	var a float64
	fmt.Print("What is the radius?: ")
	fmt.Scan(&r)
	a = math.Pi * (math.Pow(r,2))
	fmt.Println("a = ",a)
}

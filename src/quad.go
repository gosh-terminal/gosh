package main
import "fmt"
import "math"
func quad() {
	var a float64
	var b float64
	var c float64
	var item1 float64
	var item2 float64
	fmt.Print("What is a?: ")
	_, err := fmt.Scan(&a)
	fmt.Print("What is b?: ")
	_, err2 := fmt.Scan(&b)
	fmt.Print("What is c?: ")
	_, err3 := fmt.Scan(&c)
	item1 = ((-b) + math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2 * a
	item2 = ((-b) - math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2 * a
	_ = err
	_ = err2
	_ = err3
	fmt.Println("X =", item1, "and X =", item2)
}

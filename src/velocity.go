package main
import "fmt"
func velocity() {
	var displacement float64
	var time_interval float64
	var velocity float64
	fmt.Print("What is the displacment?: ")
	err, _ := fmt.Scan(&displacement)
	fmt.Print("What is the time interval?: ")
	err2, _ := fmt.Scan(&time_interval)
	_ = err
	_ = err2
	velocity = displacement / time_interval
	fmt.Println("Velocity =", velocity)
}

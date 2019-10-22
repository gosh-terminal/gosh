package main
import "fmt"
func acceleration() {
	var initial_velocity float64
	var final_velocity float64
	var initial_time float64
	var final_time float64
	var delta_t float64
	var acceleration float64
	var delta_v float64
	fmt.Print("What is the initial time?: ")
	err, _ := fmt.Scan(&initial_time)
	fmt.Print("What is the final time?: ")
	err1, _ := fmt.Scan(&final_time)
	fmt.Print("What is the initial velocity?: ")
	err2, _ := fmt.Scan(&initial_velocity)
	fmt.Print("What is the final velocity?: ")
	err3, _ := fmt.Scan(&final_velocity)
	delta_t = final_time - initial_time
	delta_v = final_velocity - initial_velocity
	acceleration = delta_v / delta_t
	fmt.Println("acceleration =", acceleration)
	_ = err
	_ = err1
	_ = err2
	_ = err3
}

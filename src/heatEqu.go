package main

import "fmt"
func heatEqu() {
	var solving_for string
	var q float64
	var m float64
	var c float64
	var dt float64
	var fi float64
	var ini float64
	fmt.Print("What are we solving for?(q,m,c,t): ")
	fmt.Scan(&solving_for)
	if solving_for == "q" {
		fmt.Print("What is m?: ")
		fmt.Scan(&m)
		fmt.Print("What is the final?: ")
		fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		fmt.Scan(&ini)
		fmt.Print("What is the specific heat?: ")
		fmt.Scan(&c)
		q = m * c * (fi - ini)
		fmt.Println("q =", q)
	} else if solving_for == "m" {
		fmt.Print("What is q?: ")
		fmt.Scan(&q)
		fmt.Print("What is the final?: ")
		fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		fmt.Scan(&ini)
		fmt.Print("What is the specific heat?: ")
		fmt.Scan(&c)
		m = q / (c * (fi - ini))
		fmt.Println("m =", m)
	} else if solving_for == "c" {
		fmt.Print("What is q?: ")
		fmt.Scan(&q)
		fmt.Print("What is the final?: ")
		fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		fmt.Scan(&ini)
		fmt.Print("What is m?: ")
		fmt.Scan(&m)
		c = q / ((fi - ini) * m)
		fmt.Println("c =", c)
	} else if solving_for == "t" {
		fmt.Print("What is q?: ")
		fmt.Scan(&q)
		fmt.Print("What is the final?: ")
		fmt.Scan(&fi)
		fmt.Print("What is the initial?: ")
		fmt.Scan(&ini)
		fmt.Print("What is m?: ")
		fmt.Scan(&m)
		dt = q / (c * m)
		fmt.Println("ΔT =", dt)
	}
}

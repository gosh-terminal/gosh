package main

import (
	"fmt"
)
func algeb() {
	var input string
	fmt.Print("Welcome to the algebra calculator available calculators include: quadratic equation(quad),Line of symmetry(los) what calc do you want?: ")
	_, err := fmt.Scan(&input)
	_ = err
	if input == "quad" {
		quad()
	} else if input == "los" {
		linOfSym()
	}
}
func areaCalc() {
	var input string
	fmt.Print("Welcome to the area finder the shapes that this can find the area for so far are trapazoid(trap),circle(circ)\nWhat do you want?: ")
	fmt.Scan(&input)
	if input == "trap" {
		AOTrap()
	} else if input == "circ" {
		AOCirc()
    }
}
func geom() {
	var input string
	fmt.Print("Welcome to the geometry calculator! available calculators include: pythagTheor(pythag),area calc(ac) \nwhat do you want?: ")
	fmt.Scan(&input)
	if input == "pythag" {
		pythagTheor()
	} else if input =="ac" {
		areaCalc()
	}
}
func chem() {
	var input string
	fmt.Print("Welcome to the chem calculator available calculators include: heat equation(heatEqu): ")
	_, err := fmt.Scan(&input)
	_ = err
	if input == "heatEqu" {
		heatEqu()
	}
}
func phys() {
	var input string
	fmt.Print("Welcome to the physics calculator available calculators include: velocity calculator(velocity),acceleration calculator(acc) what calc do you want?: ")
	_, err := fmt.Scan(&input)
	_ = err
	if input == "velocity" {
		velocity()
	} else if input == "acc" {
		acceleration()
	}
}
func misc() {
	fmt.Println("wip")
}
func main() {
	var input string
	fmt.Print("Welcome to Gocalc available calculators include algebra(algeb), geometry(geom), chemistry(chem), misc(misc), physics(phys)\nWhat calculator do you want?: ")
	_, err := fmt.Scan(&input)
	_ = err
	if input == "chem" {
		chem()
	} else if input == "algeb" {
		algeb()
	} else if input == "geom" {
		geom()
	} else if input == "misc" {
		misc()
	} else if input == "phys" {
		phys()
	} else {
		fmt.Println("\n\nERROR! Something went wrong please try again!\n\n")
		main()
	}
}

package main

import "fmt"

func main() {
	cube := New3x3Cube("R")

	fmt.Println("GOAL :\n")
	cube.printCube3x3(cube.goal)
	fmt.Println("\n\nCUBE :\n")
	cube.printCube3x3(cube.cube)
}

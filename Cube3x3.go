package main

import (
	"fmt"
)

type Cubie int

const (
	NIL Cubie = 0
	U0  Cubie = 1
	U1  Cubie = 2
	U2  Cubie = 3
	U3  Cubie = 4
	U4  Cubie = 5
	U5  Cubie = 6
	U6  Cubie = 7
	U7  Cubie = 8
	U8  Cubie = 9
	L0  Cubie = 10
	L1  Cubie = 11
	L2  Cubie = 12
	F0  Cubie = 13
	F1  Cubie = 14
	F2  Cubie = 15
	R0  Cubie = 16
	R1  Cubie = 17
	R2  Cubie = 18
	B0  Cubie = 19
	B1  Cubie = 20
	B2  Cubie = 21
	L3  Cubie = 22
	L4  Cubie = 23
	L5  Cubie = 24
	F3  Cubie = 25
	F4  Cubie = 26
	F5  Cubie = 27
	R3  Cubie = 28
	R4  Cubie = 29
	R5  Cubie = 30
	B3  Cubie = 31
	B4  Cubie = 32
	B5  Cubie = 33
	L6  Cubie = 34
	L7  Cubie = 35
	L8  Cubie = 36
	F6  Cubie = 37
	F7  Cubie = 38
	F8  Cubie = 39
	R6  Cubie = 40
	R7  Cubie = 41
	R8  Cubie = 42
	B6  Cubie = 43
	B7  Cubie = 44
	B8  Cubie = 45
	D0  Cubie = 46
	D1  Cubie = 47
	D2  Cubie = 48
	D3  Cubie = 49
	D4  Cubie = 50
	D5  Cubie = 51
	D6  Cubie = 52
	D7  Cubie = 53
	D8  Cubie = 54
)

type Cube3x3 struct {
	size       int
	rowSize    int
	columnSize int
	goal       []Cubie
	cube       []Cubie
}

func New3x3Cube(startMoves string) *Cube3x3 {
	cube := new(Cube3x3)
	cube.size = 108
	cube.rowSize = 12
	cube.columnSize = 9
	cube.goal = make([]Cubie, cube.size)
	cubie := U0

	for cubieI := range cube.goal {
		cubieIColumn := cubieI % cube.rowSize
		cubieIRow := cubieI / cube.rowSize
		if (cubieIRow > 2 && cubieIRow < 6) || ((cubieIRow < 3 || cubieIRow > 5) && cubieIColumn > 2 && cubieIColumn < 6) {
			cube.goal[cubieI] = cubie
			cubie++
		} else {
			cube.goal[cubieI] = NIL
		}
	}

	cube.cube = cube.shuffle(startMoves)
	return cube
}

func (cube Cube3x3) printCube3x3(cubeArr []Cubie) {
	for cubieI, cubie := range cubeArr {
		if cubieI/cube.rowSize != 0 && cubieI%cube.rowSize == 0 {
			fmt.Println()
		}
		if cubie < 10 && cubie > 0 {
			fmt.Print("\033[107m")
		} else if (cubie > 9 && cubie < 13) || (cubie > 21 && cubie < 25) || (cubie > 33 && cubie < 37) {
			fmt.Print("\033[42m")
		} else if (cubie > 12 && cubie < 16) || (cubie > 24 && cubie < 28) || (cubie > 36 && cubie < 40) {
			fmt.Print("\033[41m")
		} else if (cubie > 15 && cubie < 19) || (cubie > 27 && cubie < 31) || (cubie > 39 && cubie < 43) {
			fmt.Print("\033[44m")
		} else if (cubie > 18 && cubie < 22) || (cubie > 30 && cubie < 34) || (cubie > 42 && cubie < 46) {
			fmt.Print("\033[45m")
		} else if cubie > 45 {
			fmt.Print("\033[43m")
		} else {
			fmt.Print("\033[0m     ")
			continue
		}
		fmt.Printf(" %3d ", cubieI)
	}
	fmt.Println()
}

func (cube Cube3x3) moveFace(upLeft int, clockwise bool) []Cubie {
	tmpCube := make([]Cubie, cube.size)
	copy(tmpCube, cube.cube)

	if clockwise {
		tmpCube[upLeft] = cube.cube[upLeft+cube.rowSize*2]
		tmpCube[upLeft+1] = cube.cube[upLeft+cube.rowSize]
		tmpCube[upLeft+2] = cube.cube[upLeft]
		tmpCube[upLeft+cube.rowSize] = cube.cube[upLeft+cube.rowSize*2+1]
		tmpCube[upLeft+cube.rowSize+2] = cube.cube[upLeft+1]
		tmpCube[upLeft+cube.rowSize*2] = cube.cube[upLeft+cube.rowSize*2+2]
		tmpCube[upLeft+cube.rowSize*2+1] = cube.cube[upLeft+cube.rowSize+2]
		tmpCube[upLeft+cube.rowSize*2+2] = cube.cube[upLeft+2]
	} else {
		tmpCube[upLeft] = cube.cube[upLeft+2]
		tmpCube[upLeft+1] = cube.cube[upLeft+cube.rowSize+2]
		tmpCube[upLeft+2] = cube.cube[upLeft+cube.rowSize*2+2]
		tmpCube[upLeft+cube.rowSize] = cube.cube[upLeft+1]
		tmpCube[upLeft+cube.rowSize+2] = cube.cube[upLeft+cube.rowSize*2+1]
		tmpCube[upLeft+cube.rowSize*2] = cube.cube[upLeft]
		tmpCube[upLeft+cube.rowSize*2+1] = cube.cube[upLeft+cube.rowSize]
		tmpCube[upLeft+cube.rowSize*2+2] = cube.cube[upLeft+cube.rowSize*2]
	}

	return tmpCube
}

func (cube Cube3x3) moveUp(clockwise bool) []Cubie {

	// UP
	tmpCube := cube.moveFace(3, clockwise)

	if clockwise {
		// LEFT, FRONT, RIGHT
		for cubie := 36; cubie < 45; cubie++ {
			tmpCube[cubie] = cube.cube[cubie+3]
		}

		// BACK
		tmpCube[45] = cube.cube[36]
		tmpCube[46] = cube.cube[37]
		tmpCube[47] = cube.cube[38]
	} else {

		// LEFT
		tmpCube[36] = cube.cube[45]
		tmpCube[37] = cube.cube[46]
		tmpCube[38] = cube.cube[47]

		// FRONT, RIGHT, BACK
		for cubie := 39; cubie < 48; cubie++ {
			tmpCube[cubie] = cube.cube[cubie-3]
		}
	}

	return tmpCube
}

func (cube Cube3x3) moveRight(clockwise bool) []Cubie {

	// RIGHT
	tmpCube := cube.moveFace(42, clockwise)

	if clockwise {

		// UP, FRONT
		for cubie := 5; cubie < 66; cubie += 12 {
			tmpCube[cubie] = cube.cube[cubie+12*3]
		}

		// DOWN
		tmpCube[77] = cube.cube[69]
		tmpCube[89] = cube.cube[57]
		tmpCube[101] = cube.cube[45]

		// BACK
		tmpCube[45] = cube.cube[5]
		tmpCube[57] = cube.cube[17]
		tmpCube[69] = cube.cube[29]
	} else {
		// UP, FRONT
		// for cubie := 5; cubie < 66; cubie += 12 {
		// 	tmpCube[cubie] = cube.cube[cubie+12*3]
		// }

		tmpCube[5] = cube.cube[69]
		tmpCube[17] = cube.cube[57]
		tmpCube[29] = cube.cube[45]

		tmpCube[41] = cube.cube[5]
		tmpCube[53] = cube.cube[17]
		tmpCube[65] = cube.cube[29]

		// DOWN
		tmpCube[77] = cube.cube[41]
		tmpCube[89] = cube.cube[53]
		tmpCube[101] = cube.cube[65]

		// BACK
		tmpCube[45] = cube.cube[101]
		tmpCube[57] = cube.cube[89]
		tmpCube[69] = cube.cube[77]
	}

	return tmpCube
}

func (cube Cube3x3) shuffle(startMoves string) []Cubie {
	cube.cube = make([]Cubie, cube.size)
	copy(cube.cube, cube.goal)

	// cube.cube = cube.moveUp(true)
	// cube.cube = cube.moveRight(true)
	// cube.cube = cube.moveRight(false)
	// cube.cube = cube.moveUp(false)
	// cube.cube = cube.moveUp(false)
	return cube.cube
}

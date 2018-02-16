//randomDiceRoller simulates rolls of various polyhedral dice
package main

import (
	"math/rand"
	"fmt"
	"time"
)

//randomNumGen creates a random number to simulate a dice rolls.
func randomNumGen(choice int) int {
	var die int

	source := rand.NewSource(time.Now().UnixNano())
	ranVal := rand.New(source)

	if choice == 1 {	// d6 roll
		die = ranVal.Intn(6)
	} else if choice == 2 {	//d10 roll
		die = ranVal.Intn(10)
	} else if choice == 3 {	//d100 roll
		die = ranVal.Intn(100)
	} else if choice == 4 {	//d4 roll
		die = ranVal.Intn(4)
	} else if choice == 5 {	//d8 roll
		die = ranVal.Intn(8)
	} else if choice == 6 {	//d12 roll
		die = ranVal.Intn(12)
	} else if choice == 7 {	//d20 roll
		die = ranVal.Intn(20)
	} else {
		fmt.Println("Invalid number input")
	}
	return die
}

//multiDie adds dice rolls together, e.g. 2d6, 4d10, etc.
func multiDie(diceNum, dieType int) int {
	var finalRoll int
	var val int

	for val < diceNum {
		finalRoll += randomNumGen(dieType)
		val += 1
	}
	return finalRoll
}

func test()  {
	_1d6 := multiDie(1, 1)
	fmt.Println("1d6 =", _1d6)
	_2d6 := multiDie(2, 1)
	fmt.Println("2d6 =", _2d6)
	_3d6 := multiDie(3, 1)
	fmt.Println("3d6 =", _3d6)
	_4d6 := multiDie(4, 1)
	fmt.Println("4d6 =", _4d6)
	_1d10 := multiDie(1, 2)
	fmt.Println("1d10 =", _1d10)
	_2d10 := multiDie(2, 2)
	fmt.Println("2d10 =", _2d10)
	_3d10 := multiDie(3, 2)
	fmt.Println("3d10 =", _3d10)
	d100 := multiDie(1, 3)
	fmt.Println("1d100 =", d100)
	}

func main() {
	test()
}

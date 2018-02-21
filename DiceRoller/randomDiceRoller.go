//RandomDiceRoller simulates rolls of various polyhedral dice
package DiceRoller

import (
	"fmt"
	"math/rand"
	"time"
)

//RandomNumGenerator creates a random number to simulate a dice rolls.
func RandomNumGenerator(choice int) int {
	var die int

	source := rand.NewSource(time.Now().UnixNano())
	ranVal := rand.New(source)

	switch choice {
	case 1:
		die = ranVal.Intn(6) + 1
	case 2:
		die = ranVal.Intn(10) + 1
	case 3:
		die = ranVal.Intn(100) + 1
	case 4:
		die = ranVal.Intn(4) + 1
	case 5:
		die = ranVal.Intn(8) + 1
	case 6:
		die = ranVal.Intn(12) + 1
	case 7:
		die = ranVal.Intn(20) + 1
	}
	return die
}

//MultiDie adds dice rolls together, e.g. 2d6, 4d10, etc.
func MultiDie(diceNum, dieType int) int {
	var finalRoll int
	var val int

	for val < diceNum {
		finalRoll += RandomNumGenerator(dieType)
		val += 1
	}
	return finalRoll
}

//test submits various dice combinations to check the results
func test() {
	_1d6 := MultiDie(1, 1)
	fmt.Println("1d6 =", _1d6)
	_2d6 := MultiDie(2, 1)
	fmt.Println("2d6 =", _2d6)
	_3d6 := MultiDie(3, 1)
	fmt.Println("3d6 =", _3d6)
	_4d6 := MultiDie(4, 1)
	fmt.Println("4d6 =", _4d6)
	_1d10 := MultiDie(1, 2)
	fmt.Println("1d10 =", _1d10)
	_2d10 := MultiDie(2, 2)
	fmt.Println("2d10 =", _2d10)
	_1d4 := MultiDie(1, 4)
	fmt.Println("1d4 =", _1d4)
	_1d8 := MultiDie(1, 5)
	fmt.Println("1d8 =", _1d8)
	_1d12 := MultiDie(1, 6)
	fmt.Println("1d12 =", _1d12)
	_1d20 := MultiDie(1, 7)
	fmt.Println("1d20 =", _1d20)
	d100 := MultiDie(1, 3)
	fmt.Println("1d100 =", d100)
}

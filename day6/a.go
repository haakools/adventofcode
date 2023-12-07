package main

import (
	"fmt"
)

type race struct {
	maxtime  int
	distance int
}

func sum(slice []int) int {
	result := 0
	for _, val := range slice {
		result += val
	}
	return result
}

//Time:      7  15   30
//Distance:  9  40  200

func main() {
	// for the example
	//race := [3]race{{7, 9}, {15, 40}, {30, 200}}

	race := [4]race{
		{53, 313},
		{89, 1090},
		{76, 1214},
		{98, 1201},
	}

	// foreach race
	var nPossibleWin []int

	for i := range race {
		winCounter := 0
		for j := 0; j < race[i].maxtime; j++ {
			speed := j // in mm
			remaining_time := race[i].maxtime - j
			distanceTraveled := speed * remaining_time
			fmt.Printf("Speed %d, remaining time %d, distance traveled %d\n", speed, remaining_time, distanceTraveled)
			if distanceTraveled > race[i].distance {
				winCounter++
			}
		}

		nPossibleWin = append(nPossibleWin, winCounter)
		fmt.Println("win counter: ", winCounter)
	}
	// get the product
	product := 1
	for _, val := range nPossibleWin {
		product *= val
	}
	// get the product of the number of ways to beat each records
	fmt.Println("Product of winnign configurations: ", product)
}

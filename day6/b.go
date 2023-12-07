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
	//race := race{71530, 940200}

	race := race{53897698, 313109012141201}

	winCounter := 0
	for j := 0; j < race.maxtime; j++ {
		speed := j // in mm
		remaining_time := race.maxtime - j
		distanceTraveled := speed * remaining_time
		//fmt.Printf("Speed %d, remaining time %d, distance traveled %d\n", speed, remaining_time, distanceTraveled)
		if distanceTraveled > race.distance {
			winCounter++
		}
	}
	fmt.Println("Number of possible wins: ", winCounter)
}

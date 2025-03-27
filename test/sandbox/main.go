package main

import "fmt"

func totalFruit(fruits []int) int {
	maxTotalFruit := 0
	// loop the fruites
	for i := 0; i < len(fruits); i++ {
		// set map to store the fruits
		mapFruits := make(map[int]bool)
		// set totalFruit to 0
		totalFruits := 0
		// set j = fruits index
		// loop the fruits based on j
		for j := i; j < len(fruits); j++ {
			// if map doesn't contain the fruit
			_, ok := mapFruits[fruits[j]]
			if !ok {
				// if map size is 2 then break
				if len(mapFruits) == 2 {
					break
				}
				// else add the fruit to map
				mapFruits[fruits[j]] = true
			}
			// increase totalFruit
			totalFruits++
		}
		// if totalFruit is greater than maxTotalFruit then set MaxTotalFruit to totalFruit
		if totalFruits > maxTotalFruit {
			maxTotalFruit = totalFruits
		}
	}
	// return maxTotalFruit
	return maxTotalFruit
}

func main() {
	fruits := []int{1, 2, 3, 2, 2}
	fmt.Println(totalFruit(fruits))
}

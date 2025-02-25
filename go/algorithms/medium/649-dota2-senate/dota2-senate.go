package main

import (
	"fmt"
)

func predictPartyVictory(senate string) string {
	radiant := []int{} // Slice to store the indices of Radiant senators
	dire := []int{}    // Slice to store the indices of Dire senators

	// Initialize the slices
	for i, party := range senate {
		if party == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		// Find the first Radiant senator to ban a Dire senator
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}

		radiant = radiant[1:] // Remove the Radiant senator
		dire = dire[1:]       // Remove the Dire senator
	}

	if len(radiant) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}
func main() {
	fmt.Println(`test`)
	result := predictPartyVictory("RRR")

	fmt.Println(result)
}

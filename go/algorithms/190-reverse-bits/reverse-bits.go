package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		res = res << 1   // Shift res left
		res += (num & 1) // Add the least significant bit of num to res
		num = num >> 1   // Shift num to the right
	}
	return res
}

func main() {
	fmt.Println(`test`)
	// result := reverseBits(uint32(964176192))

	var a uint32 = 8
	fmt.Printf("a & 1 = %d\n", (a & 1))

	fmt.Println(reverseBits(964176192))
}

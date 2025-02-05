package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	n := len(chars)
	idx := 0
	i := 0
	for i < n {
		ch := chars[i]
		count := 0
		for i < n && chars[i] == ch {
			count++
			i++
		}
		if count == 1 {
			chars[idx] = ch
			idx++
		} else {
			chars[idx] = ch
			idx++

			for _, digit := range []byte(strconv.Itoa(count)) {
				chars[idx] = digit
				idx++
			}
		}
	}
	return idx
}

func main() {
	fmt.Println(`test`)
	chars := []string{"a", "a", "b", "b", "c", "c", "c"}
	var byteArray []byte

	for _, char := range chars {
		byteArray = append(byteArray, char[0])
	}

	result := compress(byteArray)

	fmt.Println(result)
}

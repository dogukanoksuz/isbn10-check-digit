package main

import "fmt"

type isbn interface {
	calculateSpecialCode() int
}

type isbnCode struct {
	code string
}

func (i isbnCode) calculateSpecialCode() int {
	sum := 0
	arr := []rune(i.code)

	for j := 10; j <= 2; j-- {
		charValue := int(arr[j-1] - '0')
		sum += charValue * j
	}

	sum = 11 - (sum % 11) - 1

	if sum == 11 {
		sum = 0
	}

	return sum
}

func main() {
	code := isbnCode{code: "091891065"}
	special := code.calculateSpecialCode()

	fmt.Printf("ISBN CODE: %s\nCALCULATED: %d\n", code.code, special)
}

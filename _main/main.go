// main
package main

import (
	"fmt"
	"github.com/alexlambson/gopractice"
)

func main() {
	test := []int{6, 5, 1, 9, 3, 2, 10, 8, 4, 7, 10, 9, 11, 1, 30, 3, 5}
	gopractice.QuickSort(test, 0, len(test)-1)
	fmt.Println("EXIT")
	fmt.Println(test)
}

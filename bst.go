// bst
package gopractice

import (
	"fmt"
)

var numbersToSort = []int{12, 8, 1, 5, 6, 9, 4, 2, 7, 10, 59, 22, 45, 78}

type nodeInt struct {
	value  int
	parent *nodeInt
	left   *nodeInt
	right  *nodeInt
}

func (elt *nodeInt) Print() {
	fmt.Println(elt.value)
}

func (elt *nodeInt) GetLeft() *nodeInt {
	return elt.left
}

func (elt *nodeInt) GetRight() *nodeInt {
	return elt.right
}

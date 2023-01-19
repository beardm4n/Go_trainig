package main

import (
	"fmt"
)

func main() {
	list := make([]int, 5, 7)
	fmt.Println("list", list, "length", len(list), "capacity", cap(list))

	list = Append(list, 1)
	fmt.Println("list", list, "length", len(list), "capacity", cap(list))

	list = Append(list, 2, 3)
	fmt.Println("list", list, "length", len(list), "capacity", cap(list))

	list = Append(list, 4, 5, 6)
	fmt.Println("list", list, "length", len(list), "capacity", cap(list))

	list = Append(list, 7, 8, 9, 10)
	fmt.Println("list", list, "length", len(list), "capacity", cap(list))
}

func Append(list []int, elements ...int) []int {
	var res []int

	resLen := len(list) + len(elements)

	if (resLen <= cap(list)) {
		res = list[:resLen]
	} else {
		resCap := resLen

		if resCap < 2 * len(list) {
			resCap = 2 * len(list)
		}

		res = make([]int, resLen, resCap)

		copy(res, list)
	}

	for i, el := range elements {
		res[len(list) + i] = el
	}

	return res
}
package main

import (
	"fmt"
	"strconv"
	"strings"
)
// можно решить без строк?
func main() {

	arr := []int{28, 2, 20, 19, 18, 12, 14}
	val := 2

	valstr := strconv.Itoa(val)
	str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), ","), "[]")

	strarr := strings.Split(str, ",")

	fmt.Println(strarr)
	fmt.Printf("%T\n", strarr)
	var numbers []int
	for _, v := range strarr {
		num, err := strconv.Atoi(v)
		if err == nil && v != valstr {
			numbers = append(numbers, num)
		}

	}
	fmt.Println(numbers)

}

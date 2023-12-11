package main

import "fmt"

func main() {
	//в этой задаче излишнее перевыделение памяти
	arr := []int{1, 2, 2, 3, 4, 4, 4, 4, 4, 5}
	fmt.Println(removeElement(arr))

	fmt.Println(arr)
}
func removeElement(slice []int) int {
	for i, x := range slice {
		for j := i + 1; j < len(slice); j++ {
			if slice[j] == x {
				slice = append(slice[:j], slice[j+1:]...)
				j--
			} else {
				break
			}
		}
	}

	fmt.Println(slice)
	return len(slice)
}

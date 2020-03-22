package main

import "fmt"

func NameValue(my_list []string) []int {
	var result = make([]int, len(my_list))

	for idx, str := range my_list {
		for _, chr := range str {
			if chr >= 'a' && chr <= 'z' {
				result[idx] += int(chr-'a') + 1
			}
		}

		result[idx] = result[idx] * (idx + 1)
	}

	return result
}

func main() {
	t1 := []string{"abc", "abc", "abc", "abc"}
	fmt.Println(NameValue(t1))

	t2 := []string{"codewars", "abc", "xyz"}
	fmt.Println(NameValue(t2))
}

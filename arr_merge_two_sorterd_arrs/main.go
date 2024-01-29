package main

import "fmt"

func main() {
	res := merge([]int{1, 2, 5, 9}, []int{1, 2, 3, 4, 6, 7, 8})
	fmt.Printf("result: %+v", res)

}

func merge(arr1 []int, arr2 []int) []int {
	len1 := len(arr1)
	len2 := len(arr2)
	i, j := 0, 0
	res := make([]int, 0, len1+len2)
	for x := 0; x < min(len1, len2); x++ {
		if i > len1-1 {
			res = append(res, arr2[j:]...)
			break
		}

		if j > len2-1 {
			res = append(res, arr1[i:]...)
			break
		}

		if arr1[i] <= arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}

	return res
}

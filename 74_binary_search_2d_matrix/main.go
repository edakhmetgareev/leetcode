package main

import "fmt"

// You are given an m x n integer matrix with the following two properties:
// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.
//
// You must write a solution in O(log(m * n)) time complexity.

func main() {
	matrix1 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(searchMatrix(matrix1, 5))
}

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)*len(matrix[0])-1

	for left < right {
		// 1, 1
		middle := (left + right) / 2

		i, j := middle/len(matrix[0]), middle%len(matrix[0])
		fmt.Println(matrix[i][j])
		if matrix[i][j] == target {
			return true
		}

		if target < matrix[i][j] {
			right = middle + 1
			continue
		}

		if target > matrix[i][j] {
			left = middle
			continue
		}
	}

	return false
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle values.
//
// For examples, if arr = [2,3,4], the median is 3.
// For examples, if arr = [1,2,3,4], the median is (2 + 3) / 2 = 2.5.
// You are given an integer array nums and an integer k. There is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.
//
// Return the median array for each window in the original array. Answers within 10-5 of the actual value will be accepted.
//
// Example 1:
// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000]
// Explanation:
// Window position                Median
// ---------------                -----
// [1  3  -1] -3  5  3  6  7        1
// 1 [3  -1  -3] 5  3  6  7       -1
// 1  3 [-1  -3  5] 3  6  7       -1
// 1  3  -1 [-3  5  3] 6  7        3
// 1  3  -1  -3 [5  3  6] 7        5
// 1  3  -1  -3  5 [3  6  7]       6

// Example 2:
// Input: nums = [1,2,3,4,2,3,1,4,2], k = 3
// Output: [2.00000,3.00000,3.00000,3.00000,2.00000,3.00000,2.00000]
//
// Constraints:
//
// 1 <= k <= nums.length <= 105
// -231 <= nums[i] <= 231 - 1

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now).String())
	}()
	data, err := ReadInts()
	if err != nil {
		log.Fatal(err)
	}
	_ = data
	// result := medianSlidingWindow(data, 5000)
	result := medianSlidingWindow([]int{1, 2}, 1)
	fmt.Println(result)
}

func ReadInts() ([]int, error) {
	file, err := os.ReadFile("/Users/evgeniiakhmetgareev/projects/letcode/480_sliding_window_median/data_source.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	r := strings.NewReader(string(file))
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

// Current version is too slow.
func medianSlidingWindow(nums []int, k int) []float64 {
	var result []float64
	window := make([]int, k)
	copy(window, nums[0:k])
	sort.Ints(window)
	for i := 0; i+k <= len(nums); i++ {
		if k == 1 {
			result = append(result, getMedian(nums[i:i+1], k))
			continue
		}

		if i > 0 {
			removeIdx := search(window, nums[i-1])
			addElm := nums[i+k-1]
			window = reSort(window, addElm, removeIdx)
		}

		result = append(result, getMedian(window, k))
	}

	return result
}

func reSort(window []int, addElm int, removeIdx int) []int {
	result := make([]int, len(window))
	copy(result, window)

	if result[removeIdx] == addElm {
		return result
	}

	result = append(result[:removeIdx], result[removeIdx+1:]...)
	var idx int
	left, right := 0, len(result)-1
	for left <= right {
		idx = (right + left) / 2

		if left == right || result[idx] == addElm {
			break
		}

		if result[idx] < addElm {
			left = idx + 1
		}

		if result[idx] > addElm {
			right = idx - 1
		}
	}

	if result[idx] < addElm {
		idx++
	}

	result = append(result[:idx+1], result[idx:]...)
	result[idx] = addElm

	return result
}

func getMedian(window []int, k int) float64 {
	if k == 1 {
		return float64(window[0])
	}
	if k%2 != 0 {
		middle := len(window) / 2
		return float64(window[middle])
	}

	middle1, middle2 := len(window)/2-1, len(window)/2
	return float64(window[middle2]+window[middle1]) / 2
}

func search(nums []int, target int) int {
	right := len(nums) - 1
	left := 0
	for left <= right {
		idx := (right + left) / 2

		if nums[idx] == target {
			return idx
		}

		if nums[idx] > target {
			right = idx - 1
			continue
		}

		if nums[idx] < target {
			left = idx + 1
			continue
		}
	}

	return 0
}

// //

// not mine
func medianSlidingWindow3(nums []int, k int) []float64 {
	var res []float64

	var minHeap minH
	var maxHeap maxH

	for i, num := range nums {
		if len(maxHeap) == 0 || num <= maxHeap[0] {
			maxHeap.add(num)
		} else {
			minHeap.add(num)
		}

		if len(maxHeap) > len(minHeap)+1 {
			minHeap.add(maxHeap.pop())
		} else if len(minHeap) > len(maxHeap) {
			maxHeap.add(minHeap.pop())
		}

		if i-k+1 >= 0 {
			removedIndex := i - k

			if removedIndex >= 0 {
				if nums[removedIndex] <= maxHeap[0] {
					maxHeap.remove(nums[removedIndex])
				} else {
					minHeap.remove(nums[removedIndex])
				}

				if len(maxHeap) > len(minHeap)+1 {
					minHeap.add(maxHeap.pop())
				} else if len(minHeap) > len(maxHeap) {
					maxHeap.add(minHeap.pop())
				}
			}

			var median float64

			if len(minHeap) == len(maxHeap) {
				median = float64(minHeap[0]+maxHeap[0]) / float64(2)
			} else {
				median = float64(maxHeap[0])
			}

			res = append(res, median)
		}
	}

	return res
}

type minH []int

type maxH []int

func (h *maxH) remove(item int) {
	for i := 0; i < len(*h); i++ {
		if item == (*h)[i] {
			(*h)[i] = (*h)[len(*h)-1]
			*h = (*h)[:len(*h)-1]

			if i == 0 {
				h.heapDown(0)
				return
			}

			parent := (i - 1) / 2

			if parent >= 0 && i < len(*h) {
				h.heapDown(i)
				h.heapUp(i)
				return
			}
		}
	}
}

func (h *minH) remove(item int) {
	for i := 0; i < len(*h); i++ {
		if item == (*h)[i] {
			(*h)[i] = (*h)[len(*h)-1]
			*h = (*h)[:len(*h)-1]

			if i == 0 {
				h.heapDown(0)
				return
			}

			parent := (i - 1) / 2

			if parent >= 0 && i < len(*h) {
				h.heapDown(i)
				h.heapUp(i)
				return
			}
		}
	}
}

func (h *maxH) pop() int {
	poppedItem := (*h)[0]

	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.heapDown(0)

	return poppedItem
}

func (h *minH) pop() int {
	poppedItem := (*h)[0]

	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.heapDown(0)

	return poppedItem
}

func (h *minH) add(num int) {
	*h = append(*h, num)
	h.heapUp(len(*h) - 1)
}

func (h *maxH) add(num int) {
	*h = append(*h, num)
	h.heapUp(len(*h) - 1)
}

func (h *minH) heapDown(p int) {
	l, r := 2*p+1, 2*p+2
	smaller := p

	if l < len(*h) && (*h)[l] < (*h)[smaller] {
		smaller = l
	}

	if r < len(*h) && (*h)[r] < (*h)[smaller] {
		smaller = r
	}

	if smaller != p {
		(*h)[smaller], (*h)[p] = (*h)[p], (*h)[smaller]
		h.heapDown(smaller)
	}
}

func (h *minH) heapUp(p int) {
	parent := (p - 1) / 2

	if parent >= 0 && (*h)[p] < (*h)[parent] {
		(*h)[p], (*h)[parent] = (*h)[parent], (*h)[p]
		h.heapUp(parent)
	}
}

func (h *maxH) heapDown(p int) {
	l, r := 2*p+1, 2*p+2
	bigger := p

	if l < len(*h) && (*h)[l] > (*h)[bigger] {
		bigger = l
	}

	if r < len(*h) && (*h)[r] > (*h)[bigger] {
		bigger = r
	}

	if bigger != p {
		(*h)[bigger], (*h)[p] = (*h)[p], (*h)[bigger]
		h.heapDown(bigger)
	}
}

func (h *maxH) heapUp(p int) {
	parent := (p - 1) / 2

	if parent >= 0 && (*h)[p] > (*h)[parent] {
		(*h)[p], (*h)[parent] = (*h)[parent], (*h)[p]
		h.heapUp(parent)
	}
}

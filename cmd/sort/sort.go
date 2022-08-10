package main

import (
	"fmt"
)

func SelectSort(a []int, n int) {
	if n < 0 {
		return
	}
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] > a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

func InsertSort(a []int, n int) {
	if n < 0 {
		return
	}
	for i := 1; i < n; i++ {
		temp := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > temp {
				a[j+1] = a[j]
				continue
			}
			break
		}
		a[j+1] = temp
	}
}

func BubbleSort(a []int, n int) {
	if n < 0 {
		return
	}
	for i := 0; i < n; i++ {
		isSwap := false
		for j := 0; j < n-1-i; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}

func Swap(i, j *int) {
	temp := *i
	*i = *j
	*j = temp
}

func Mingmingsuijishu(A []int, n int) {
	quick_sort(A, n)
	deduplication(A, n)
}

func deduplication(A []int, n int) {
	for i := 0; i < n; {
		fmt.Println(A[i])
		step := 1
		for j := i + 1; j < n; {
			if A[i] == A[j] {
				step++
				j++
				continue
			}
			break
		}
		i += step
	}
}

func quick_sort(A []int, n int) {
	quick_sort_c(A, 0, n-1)
}

func quick_sort_c(A []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(A, p, r)
	quick_sort_c(A, p, q-1)
	quick_sort_c(A, q+1, r)
}

func partition(A []int, p, r int) int {
	pivot := A[r]
	i := p
	for j := p; j < r; j++ {
		if A[j] < pivot {
			temp := A[j]
			A[j] = A[i]
			A[i] = temp
			i++
		}
	}
	temp := A[i]
	A[i] = A[r]
	A[r] = temp
	return i
}

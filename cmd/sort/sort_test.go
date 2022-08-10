package main

import (
	"leetcode/pkg/util"
	"math/rand"
	"testing"
)

func TestSelectSort(t *testing.T) {
	testCase := util.GenerateCase(100)
	t.Log(testCase)
	SelectSort(testCase, 100)
	for i := 0; i < len(testCase)-1; i++ {
		if testCase[i] < testCase[i+1] {
			t.Error("select sort is not correct")
			break
		}
	}
	t.Log(testCase)
}

func TestInsertSort(t *testing.T) {
	testCase := util.GenerateCase(100)
	t.Log(testCase)
	InsertSort(testCase, 100)
	for i := 0; i < len(testCase)-1; i++ {
		if testCase[i] < testCase[i+1] {
			t.Error("insert sort is not correct")
			break
		}
	}
	t.Log(testCase)
}

func TestBubbleSort(t *testing.T) {
	testCase := util.GenerateCase(100)
	t.Log(testCase)
	BubbleSort(testCase, 100)
	for i := 0; i < len(testCase)-1; i++ {
		if testCase[i] < testCase[i+1] {
			t.Error("bubble sort is not correct")
			break
		}
	}
	t.Log(testCase)
}

func TestSwap(t *testing.T) {
	i := 9
	j := 8
	t.Log(i, j)
	Swap(&i, &j)
	if i != 8 || j != 9 {
		t.Error("invalid swap")
	}
	t.Log(i, j)
}

func TestQuickSort(t *testing.T) {
	arr := make([]int, 0, 500)
	for i := 0; i < 6; i++ {
		j := int(rand.Float64() * 500)
		arr = append(arr, j)
	}
	quick_sort(arr, 6)
	deduplication(arr, 6)
}

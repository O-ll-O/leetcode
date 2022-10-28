package util

import (
	"fmt"
	"math/rand"
	"reflect"
	"unsafe"
)

func GenerateCase(n int) []int {
	if n < 0 {
		n = 100
	}
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		j := int(rand.Float64() * 500)
		arr = append(arr, j)
	}
	return arr
}

func StrToBytes(s string) []byte {
	sptr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(unsafe.Pointer(&s))
	fmt.Println(unsafe.Pointer(sptr))
	b := (*[]byte)(unsafe.Pointer(sptr))
	fmt.Println(unsafe.Pointer(b))
	return *b
}

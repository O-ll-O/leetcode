package main

import (
	"bufio"
	"fmt"
	"os"
	"unsafe"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	source := sc.Bytes()
	for i, j := 0, len(source)-1; i < j; i, j = i+1, j-1 {
		source[i], source[j] = source[j], source[i]
	}
	fmt.Print(BytesToStr(source))
}

func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

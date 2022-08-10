package main

import (
	"bufio"

	"os"
	"strconv"
)

func main() {
	bucket := make([]int, 0)
	r := bufio.NewReader(os.Stdin)
	for {
		in, _, _ := r.ReadLine()
		if len(in) == 0 {
			Mingmingsuijishu(bucket, len(bucket))
			break
		}

		count, _ := strconv.Atoi(string(in))

		for i := 0; i < count; i++ {
			nb, _, _ := r.ReadLine()
			n, _ := strconv.Atoi(string(nb))
			bucket = append(bucket, n)
		}
	}

}

package main

import (
	"bufio"
	"fmt"

	"os"
	"strconv"
	"strings"
)

//knapsack problem
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	arr := strings.Split(sc.Text(), " ")
	total, _ := strconv.Atoi(arr[0])
	total /= 10
	m, _ := strconv.Atoi(arr[1])
	W := make([][3]int, m)
	V := make([][3]int, m)

	for i := 0; i < m; i++ {
		sc.Scan()
		params := strings.Split(sc.Text(), " ")

		v, _ := strconv.Atoi(params[0])
		v /= 10
		p, _ := strconv.Atoi(params[1])
		q, _ := strconv.Atoi(params[2])
		if q == 0 {
			W[i][0] = v
			V[i][0] = p * v
			continue
		}
		if q > 0 {
			q--
		}
		if W[q][1] == 0 {
			W[q][1] = v
			V[q][1] = p * v
			continue
		}
		if W[q][2] == 0 {
			W[q][2] = v
			V[q][2] = p * v
			continue
		}
	}

	for i := 0; i < len(W); {
		if W[i][0] == 0 {
			W = append(W[:i], W[i+1:]...)
			V = append(V[:i], V[i+1:]...)
			continue
		}
		i++
	}
	dp := make([][]int, len(W)+1)
	for i := 0; i < len(W)+1; i++ {
		dp[i] = make([]int, total+1)
	}

	for i := 1; i < len(W)+1; i++ {
		w1 := W[i-1][0]
		w2 := W[i-1][1]
		w3 := W[i-1][2]
		v1 := V[i-1][0]
		v2 := V[i-1][1]
		v3 := V[i-1][2]
		for j := 0; j < total+1; j++ {
			//# 1. 不放入:
			dp[i][j] = dp[i-1][j]
			//# 2. 放入一个主件
			if j-w1 >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-w1]+v1)
			}
			//# 3. 1个主件+附件1
			if j-w1-w2 >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-w2-w1]+v1+v2)
			}
			//# 4. 一个主件+附件2
			if j-w1-w3 >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-w3-w1]+v1+v3)
			}

			//# 5. 一个主见+附件1+附件2
			if j-w1-w2-w3 >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-w3-w2-w1]+v1+v2+v3)
			}
		}
	}
	fmt.Print(dp[len(W)][total] * 10)
}
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var j int

func main() {
	var ip, mask string
	var lineParts []string
	var a, b, c, d, e, errCount, private int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lineParts = strings.Split(scanner.Text(), "~")
		ip = lineParts[0]
		mask = lineParts[1]

		ipSlice := strToInts(ip)
		maskSlice := strToInts(mask)

		if ipSlice[0] == 0 || ipSlice[0] == 127 {
			continue // 类似于【0.*.*.*】和【127.*.*.*】的IP地址不属于上述输入的任意一类，也不属于不合法ip地址，计数时请忽略
		}

		if ipSlice == nil || maskSlice == nil || maskSlice[0] == 0 || maskSlice[3] == 255 || isErrMask(maskSlice) {

			errCount++
			continue
		}
		// 以下开始 IP类型计数
		//private
		/*
		   私网IP范围是：

		   从10.0.0.0到10.255.255.255

		   从172.16.0.0到172.31.255.255

		   从192.168.0.0到192.168.255.255
		*/
		if (ipSlice[0] == 172 && (ipSlice[1] == 16 || ipSlice[1] == 31)) ||
			ipSlice[0] == 10 || (ipSlice[0] == 192 && ipSlice[1] == 168) {
			private++
		}
		//通过高8位用来区分ip的类别,假设高八位的值为x
		//e 11110 000  x>>3==0+2+4+8+16
		if ipSlice[0] >= 240 {
			e++
			continue
		}
		//d 1110 0000  x>>4==0+2+4+8
		if ipSlice[0] >= 224 {
			d++
			continue
		}
		//c 110 00000  x>>5==0+2+4
		if ipSlice[0] >= 192 {
			c++
			continue
		}
		//b 10 000000  x>>6==0+2
		if ipSlice[0] >= 128 {
			b++
			continue
		}
		//a 0 0000000  x>>7==0
		if ipSlice[0] >= 1 {
			a++
			continue
		}

	}

	fmt.Println(a, b, c, d, e, errCount, private)
}
func isErrMask(maskSlice []byte) bool {
	maskInt := int(binary.BigEndian.Uint32(maskSlice))
	min1Idx := -1
	max0Idx := -1
	for i := 0; i < 32; i++ {
		if maskInt&1 == 0 {
			max0Idx = i
		}
		if maskInt&1 == 1 && min1Idx == -1 {
			min1Idx = i
		}

		maskInt = maskInt >> 1
	}
	return max0Idx > min1Idx
}

func strToInts(ip string) []byte {
	s := strings.Split(ip, ".")
	if len(s) != 4 {
		return nil
	}

	r := make([]byte, 4)

	for i := 0; i < 4; i++ {
		num, err := strconv.Atoi(s[i])
		if err != nil {
			return nil
		}
		if num > 255 {
			return nil
		}
		r[i] = uint8(num)
	}
	return r
}

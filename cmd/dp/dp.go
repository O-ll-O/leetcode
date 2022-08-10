package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sumv, number, things := CollectData()

	log.Print(sumv, number, things)
}

func CollectData() (int, int, []Thing) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	arr := strings.Split(sc.Text(), " ")
	sumv, _ := strconv.Atoi(arr[0])
	number, _ := strconv.Atoi(arr[1])
	items := make([]Thing, number)
	for i := 0; i < number; i++ {
		sc.Scan()
		arr := strings.Split(sc.Text(), " ")
		v, _ := strconv.Atoi(arr[0])
		p, _ := strconv.Atoi(arr[1])
		q, _ := strconv.Atoi(arr[2])
		item := Thing{v, p, q, nil, nil}
		items = append(items, item)
	}
	return sumv, number, items
}

type Thing struct {
	V      int
	P      int
	Q      int
	thing1 *Thing
	thing2 *Thing
}

type Things []Thing

func (ts Things) Conver2() {

}

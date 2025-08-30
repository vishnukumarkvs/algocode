package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// anagram
	q := []string{"ola", "alo"}

	a := strings.Split(q[0], "")
	sort.Strings(a)
	aa := strings.Join(a, "")

	b := strings.Split(q[1], "")
	sort.Strings(b)
	bb := strings.Join(b, "")

	if aa == bb {
		fmt.Println("anagram")
	}

	// count numbers

	n1 := []string{"1", "2", "3"}

	cnt := 0

	for _, i := range n1 {
		tmp, _ := strconv.Atoi(i)
		cnt += tmp
	}
	fmt.Println(cnt)

}

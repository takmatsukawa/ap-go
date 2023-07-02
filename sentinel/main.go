package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func toIntSlice(str string) []int {
	strs := strings.Split(str, " ")
	ints := make([]int, len(strs))
	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	A := toIntSlice(sc.Text())
	sc.Scan()
	B := toIntSlice(sc.Text())

	A = append(A, -2*1e9, 2*1e9)
	sort.Ints(A)

	ans := int(2 * 1e9)
	for _, b := range B {
		i := sort.Search(len(A), func(i int) bool { return A[i] >= b })
		ans = min(ans, abs(b-A[i-1]), abs(A[i]-b))
	}

	fmt.Println(ans)
}

func min(args ...int) int {
	res := args[0]
	for _, v := range args {
		if v < res {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/* func calcMove(s []int) int {
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return i
		}
	}
	return 0
} */

func calcMove(s []int) int {
	l, h := 0, len(s)-1

	for l <= h {
		m := l + (h-l)/2

		if m > 0 && s[m-1] > s[m] {
			return m
		} else if m+1 < len(s) && s[m] > s[m+1] {
			return m + 1
		}

		if s[m] >= s[h] {
			l = m + 1
		} else if s[m] <= s[l] {
			h = m - 1
		} else {
			return 0
		}
	}

	return 0
}

func main() {
	rand.Seed(time.Now().Unix())

	s := make([]int, 8)
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	m := rand.Intn(6)
	for i := 0; i < m; i++ {
		s = append(s[len(s)-1:], s[:len(s)-1]...)
	}

	fmt.Println(s)
	fmt.Println(calcMove(s))
}

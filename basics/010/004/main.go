package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 低买高买
func maxProfit(p []int) int {
	s := 0

	for i := 1; i < len(p); i++ {
		if p[i-1] < p[i] {
			j := i
			for j < len(p) && p[j-1] <= p[j] {
				j++
			}
			s += p[j-1] - p[i-1]
			i = j
		}
	}

	return s
}

func main() {
	rand.Seed(time.Now().Unix())

	s := make([]int, 5)
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}

	fmt.Println(s)
	fmt.Println(maxProfit(s))
}

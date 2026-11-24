package main

import (
	"fmt"
	"sort"
)

// FourSum returns all unique quadruplets of values from nums summing to
// target, in sorted order: sort once, fix two values in nested loops, and
// close each pair with two pointers. Duplicate values are skipped at every
// level so no quadruplet repeats.
func FourSum(nums []int, target int) [][]int {
	a := append([]int{}, nums...)
	sort.Ints(a)
	n := len(a)
	out := [][]int{}
	for i := 0; i < n-3; i++ {
		if i > 0 && a[i] == a[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && a[j] == a[j-1] {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				s := int64(a[i]) + int64(a[j]) + int64(a[l]) + int64(a[r])
				switch {
				case s == int64(target):
					out = append(out, []int{a[i], a[j], a[l], a[r]})
					l++
					r--
					for l < r && a[l] == a[l-1] {
						l++
					}
					for l < r && a[r] == a[r+1] {
						r--
					}
				case s < int64(target):
					l++
				default:
					r--
				}
			}
		}
	}
	return out
}

func main() {
	fmt.Println(FourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

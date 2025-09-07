package main

import "fmt"

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func putIfAbsent(slice []int, val int) []int {
	if !contains(slice, val) {
		slice = append(slice, val)
	}
	return slice
}

func unique(slice []int) []int {
	res := []int{}
	for _, v := range slice {
		res = putIfAbsent(res, v)
	}
	return res
}

func intersect(a, b []int) []int {
	res := []int{}
	for _, v := range a {
		if contains(b, v) {
			res = putIfAbsent(res, v)
		}
	}
	return res
}

func union(a, b []int) []int {
	res := []int{}
	for _, v := range a {
		res = putIfAbsent(res, v)
	}
	for _, v := range b {
		res = putIfAbsent(res, v)
	}
	return res
}

func main() {
	a := []int{5, 1, 2, 5}
	b := []int{4, 2, 5, 1, 1, 2}

	ua := unique(a)
	ub := unique(b)
	i := intersect(a, b)
	u := union(a, b)

	fmt.Println(ua, ub)
	fmt.Println(i)
	fmt.Println(u)
}

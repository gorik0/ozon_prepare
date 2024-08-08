package main

import "testing"

func TestFindSquareAquiredAllPoints(t *testing.T) {
	p1 := [2]int{1, 10}
	p2 := [2]int{0, 2}
	p3 := [2]int{5, 9}
	points := [][2]int{p1, p2, p3}
	println(FindSquareAquiredAllPoints(points))
}

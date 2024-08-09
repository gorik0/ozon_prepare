package main

import "log"

func main() {
	p1 := Point{1, 1}
	p2 := Point{2, 2}
	p9 := Point{9, 9}
	p8 := Point{8, 8}
	p4 := Point{4, 4}
	p5 := Point{5, 5}
	ma = map[int][]Point{1: {p1, p2}, 2: {p9, p8}, 3: {p4, p5}}
	//ma[0] = [][2]int{p1, p2}
	//ma[1] = [][2]int{p4, p5}
	//ma[2] = [][2]int{p9, p8, p1}
	//log.Println(len(FindCombinations(ma)))

	//l1 := []string{"a", "b"}
	//l2 := []string{"c", "d"}
	//Through([][]string{l1, l2}, "")
	//log.Println(combinations)

	log.Println(ma)
	println("Helo")
	key := make([]int, 0)
	for ke, _ := range ma {
		key = append(key, ke)

	}
	acc := make([]Point, 0)
	ThMa(key, acc)
	log.Println(combinations)
	println(len(combinations))
}

type Point struct {
	x int
	y int
}

var combinations [][]Point
var ma map[int][]Point

func ThMa(keys []int, accum []Point) {
	var item []Point
	last := len(keys) == 1
	n := len(ma[keys[0]])
	for i := range n {
		item = append(accum, ma[keys[0]][i])
		if last {
			combinations = append(combinations, item)
		} else {
			ThMa(keys[1:], item)
		}

	}
}

//func Through(terms [][]string, accum string) {
//
//	last := len(terms) == 1
//	n := len(terms[0])
//	for i := range n {
//		item := accum + terms[0][i]
//		if last {
//			combinations = append(combinations, item)
//		} else {
//			Through(terms[1:], item)
//		}
//
//	}
//}

func FindCombinations(slices map[int][][2]int) [][][2]int {
	combins := make([][][2]int, 0)
	var comb = make([][2]int, 0)

	indSlices := make([]int, 0)
	for _, slice := range slices {
		indSlices = append(indSlices, len(slice)-1)
	}
gorik:
	for {

		for i := range len(slices) {
			slice := slices[i]

			if i == 0 {
				comb = make([][2]int, 0)
			}
			if i == 0 && indSlices[i] == -1 {
				break gorik
			}
			if indSlices[i] == -1 {
				indSlices[i] = len(slice) - 1
			}

			//log.Println("sss", slice[indSlices[i]])
			comb = append(comb, slice[indSlices[i]])

			if i == len(slices)-1 {
				indSlices[i] -= 1
				updateIndSlice(indSlices)
			}

			//	:::: get elemnt from slice for index

		}
		combins = append(combins, comb)

	}
	return combins
}

func updateIndSlice(s []int) {
	for i := len(s) - 1; i > 0; i-- {

		if s[i] < 0 {
			s[i-1] -= 1
		} else {
			break
		}

	}
}

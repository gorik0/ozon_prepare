package main

import "log"

func main() {
	p1 := [2]int{1, 1}
	p2 := [2]int{2, 2}
	p9 := [2]int{9, 9}
	p8 := [2]int{8, 8}
	p4 := [2]int{4, 4}
	p5 := [2]int{5, 5}
	ma := make(map[int][][2]int)
	ma[0] = [][2]int{p1, p2}
	ma[1] = [][2]int{p4, p5}
	ma[2] = [][2]int{p9, p8, p1}
	log.Println(len(FindCombinations(ma)))

	println("Helo")
}

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

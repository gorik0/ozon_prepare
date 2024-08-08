package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

func main() {
	var (
		dataCount, width, length, resAmount, resTypes, x, y, minSquare int

		//coord                                               []int
		//ress                                                [][]int
		ma  map[int][][2]int
		in  *bufio.Reader
		out *bufio.Writer
	)
	_ = minSquare

	// IN--OUT
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// IN--OUT

	//:: scan DATA count
	fmt.Fscan(in, &dataCount)
	//:: scan DATA count

	for range dataCount {
		//	 scan (width,length)
		fmt.Fscan(in, &width, &length)
		//	 scan (width,length)

		//	 scan ressTypes
		fmt.Fscan(in, &resTypes)
		//	 scan resTypes
		ma = make(map[int][][2]int, resTypes)
		for i := range resTypes {

			//	 scan ressAmount
			fmt.Fscan(in, &resAmount)
			log.Println("resAmount:", resAmount)
			//	 scan ressAmount
			ma[i] = make([][2]int, resAmount)

			for j := range resAmount {

				//	 scan coord
				fmt.Fscan(in, &x, &y)
				log.Println("x :::", x)
				log.Println("y :::", y)
				ma[i][j] = [2]int{x, y}
				//	 scan coord

			}

		}

		log.Printf("dataCount = %d, width = %d, length = %d\n", dataCount, width, length)
		log.Println("resTypes = ", resTypes)
		log.Println("ma = ", ma)

		combinations := FindCombinations(ma)
		log.Println("Combinatiosn ::; ", combinations)
		for _, combination := range combinations {

			sq := FindSquareAquiredAllPoints(combination)
			println(sq)
			if minSquare == 0 {
				minSquare = sq
			} else if sq < minSquare {
				minSquare = sq
			}
		}
		println("minSquare ::: ", minSquare)

		//var points [][2]int
		//var p [2]int
		//
		//points = append(points, p)
		//

	}

	//for i, quntCombin := range ressQuantitytTypes {

	//}

}

func FindSquareAquiredAllPoints(points [][2]int) (squareAquired int) {
	var x, y []int
	for _, point := range points {
		x = append(x, point[0])
		y = append(y, point[1])
	}
	xMax := slices.Max(x)
	yMax := slices.Max(y)
	yMin := slices.Min(y)
	xMin := slices.Min(x)

	return findSquareByTwoPoints([2]int{xMin, yMin}, [2]int{xMax, yMax})

}

func findSquareByTwoPoints(p1 [2]int, p2 [2]int) int {
	x := p1[0] - p2[0]
	y := p1[1] - p2[1]
	return int(math.Abs(float64(x * y)))

}

/*
1
2 3
2
2
1 3
2 2
2
2 3
2 1

l = [slice1,slice2,slice3]

	for i  in range slice1{
			for j  in range slice2{
					for k  in range slice3{
						...Делать что нибудь...

slice1= [2]
slice2= [3]

[0,0]
[0,1]
[0,2]
[1,0]
[1,1]
[1,2]
}
}
}
*/
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

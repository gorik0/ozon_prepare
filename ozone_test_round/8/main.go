package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		dataCount, width, length, resAmount, resTypes, x, y int
		//coord                                               []int
		//ress                                                [][]int
		ma  map[int][][]int
		in  *bufio.Reader
		out *bufio.Writer
	)

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
		ma = make(map[int][][]int, resTypes)
		for i := range resTypes {

			//	 scan ressAmount
			fmt.Fscan(in, &resAmount)
			log.Println("resAmount:", resAmount)
			//	 scan ressAmount
			ma[i] = make([][]int, resAmount)

			for j := range resAmount {

				ma[i][j] = make([]int, 2)
				//	 scan coord
				fmt.Fscan(in, &x, &y)
				log.Println("x :::", x)
				log.Println("y :::", y)
				ma[i][j] = []int{x, y}
				//	 scan coord

			}

		}
		log.Printf("dataCount = %d, width = %d, length = %d\n", dataCount, width, length)
		log.Println("resTypes = ", resTypes)
		log.Println("ma = ", ma)

	}
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

*/

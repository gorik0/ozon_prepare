package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		sequenceLength, dataCount, posInInfo, element, childrenCount int
		in                                                           *bufio.Reader
		out                                                          *bufio.Writer
		mappaChildrens                                               map[int]interface{}
		listHead, roots                                              []int
	)

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	//	throughout data
	fmt.Fscan(in, &dataCount)
	for range dataCount {
		mappaChildrens = make(map[int]interface{})
		listHead = make([]int, 0)
		posInInfo = 1
		fmt.Fscan(in, &sequenceLength)

		for range sequenceLength {
			fmt.Fscan(in, &element)
			switch posInInfo {
			case 1:
				listHead = append(listHead, element)
				posInInfo++

			case 2:
				if element == 0 {
					posInInfo = 1
					break

				}
				childrenCount = element

				posInInfo++

			default:

				mappaChildrens[element] = element
				if posInInfo == childrenCount+2 {
					posInInfo = 1
					break
				}
				posInInfo++

			}
		}
		for _, i := range listHead {
			if _, ok := mappaChildrens[i]; !ok {
				roots = append(roots, i)

			}

		}

	}
	for _, root := range roots {

		fmt.Fprintln(out, root)
	}

}

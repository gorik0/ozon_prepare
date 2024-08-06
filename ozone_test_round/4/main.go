package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		maxLength, dataCount, sequenceCount, first, second, maxim, element, pos, secondpos int
		in                                                                                 *bufio.Reader
		out                                                                                *bufio.Writer
		maxLoadCount, seq                                                                  []int
	)
	_ = maxLength
	_ = maxLoadCount
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	//	throughout data
	fmt.Fscan(in, &dataCount)
	for range dataCount {
		fmt.Fscan(in, &sequenceCount)
		seq = make([]int, sequenceCount)
		first = -1
		second = -1
		maxim = 1
		maxLength = 1

		for i := range sequenceCount {
			fmt.Fscan(in, &seq[i])
		}
		for pos < len(seq) {

			element = seq[pos]
			switch {

			case element == first || element == second:
				maxim++
				break
			case first == -1:
				first = element
				break
			case second == -1:
				second = element
				secondpos = pos
				maxim++
				break

			case element != first || element != second:

				first = second
				second = -1
				if maxLength < maxim {
					maxLength = maxim
				}
				maxim = 1
				pos = secondpos

			}
			pos++

		}
		pos = 0
		if maxLength < maxim {
			maxLength = maxim
		}

		maxLoadCount = append(maxLoadCount, maxLength)

	}
	for _, i := range maxLoadCount {
		fmt.Fprintln(out, i)

	}
}

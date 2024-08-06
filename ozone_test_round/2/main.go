package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		discount, productsCount, dataCount, cost, toReturn int
		moneyMissed                                        float64
		in                                                 *bufio.Reader
		out                                                *bufio.Writer
	)

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	//	throughout data
	fmt.Fscan(in, &dataCount)
	for range dataCount {

		fmt.Fscan(in, &productsCount, &discount)

		for range productsCount {

			fmt.Fscan(in, &cost)
			toReturn = cost * discount % 100

			moneyMissed += float64(toReturn) / 100
		}
		fmt.Fprintf(out, "%.2f\n", moneyMissed)
		moneyMissed = 0
	}

}

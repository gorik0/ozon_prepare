package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"time"
)

func main() {
	var (
		dataCount, boxesCount, carsCount, carTonnage int
		boxes                                        []int
		logistics                                    []int
		in                                           *bufio.Reader
		out                                          *bufio.Writer
	)

	ti := time.Now()
	_ = boxesCount
	_ = boxes
	in = bufio.NewReader(os.Stdin)
	logistics = make([]int, 0)
	out = bufio.NewWriter(os.Stdout)

	defer out.Flush()

	fmt.Fscan(in, &dataCount)
	for range dataCount {
		fmt.Fscan(in, &carsCount, &carTonnage)

		fmt.Fscan(in, &boxesCount)
		boxes = make([]int, boxesCount)
		for i := range boxesCount {
			fmt.Fscan(in, &boxes[i])

		}

		for i, el := range boxes {
			boxes[i] = int(math.Pow(2, float64(el)))

		}
		//remain := EquipCar(boxes, int(carTonnage))
		//log.Println("REMAIN :::: ", remain)
		logistic := 0
		carsRemain := carsCount
		for {
			boxes = EquipCar(boxes, int(carTonnage))
			carsRemain--
			if len(boxes) != 0 && carsRemain == 0 {
				carsRemain = carsCount
				logistic++
			} else if len(boxes) == 0 {
				logistic++
				break
			}
		}
		logistics = append(logistics, logistic)
		//Run(in, out, boxes, carsCount, carTonnage)
	}

	for _, i := range logistics {

		fmt.Fprintln(out, i)
	}
	println(time.Since(ti).Seconds())

}

func EquipCar(boxes []int, tonnage int) (boxesRemain []int) {
	var offset int
	sort.Slice(boxes, func(i, j int) bool { return boxes[i] > boxes[j] })
	boxesRemain = make([]int, len(boxes))
	copy(boxesRemain, boxes)
	for i, box := range boxes {
		if box > tonnage {
			continue
		} else {
			tonnage -= box
			if tonnage == 0 {
				if i == 0 {
					return boxesRemain[1:]

				} else if i == len(boxes)-1 {
					return boxesRemain[:i-offset]
				} else {

					return append(boxesRemain[:i-offset], boxesRemain[i+1-offset:]...)
				}
			} else {
				if i == 0 {
					boxesRemain = boxesRemain[1:]
					offset++

				} else if i == len(boxes)-1 {
					boxesRemain = boxesRemain[:i-offset]
					offset++
				} else {

					boxesRemain = append(boxesRemain[:i-offset], boxesRemain[i+1-offset:]...)
					offset++
				}
			}
		}

	}
	return
}

/*

	1
	2 10
	3 1 3 1 3 1 3 1 0

	1
	2 10
	1

	1
	2 17
	6
	0 3 1 4 3 3



*/

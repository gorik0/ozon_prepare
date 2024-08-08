package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	var (
		dataCount, boxesCount, carsCount, carTonnage int
		boxes                                        []float64
		in                                           *bufio.Reader
		out                                          *bufio.Writer
	)

	_ = boxesCount
	_ = boxes
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	defer out.Flush()

	fmt.Fscanln(in, &dataCount)
	for range dataCount {
		fmt.Fscanln(in, &carsCount, &carTonnage)

		s, _ := in.ReadString('\n')
		println(s)
		s = strings.Replace(s, " ", ",", -1)
		s = "[" + s + "]"
		er := json.Unmarshal([]byte(s), &boxes)
		log.Println(er)

		for i, el := range boxes {
			boxes[i] = math.Pow(float64(2), el)

		}
		log.Println("Tonnage ::::", carTonnage)
		log.Println("Al ::::: ", boxes)

		//remain := EquipCar(boxes, float64(carTonnage))
		//log.Println("REMAIN :::: ", remain)
		logistic := 0
		carsRemain := carsCount
		for {
			boxes = EquipCar(boxes, float64(carTonnage))
			log.Printf("Boxes remain ::: %v\n", boxes)
			carsRemain--
			if len(boxes) != 0 && carsRemain == 0 {
				carsRemain = carsCount
				logistic++
			} else if len(boxes) == 0 {
				logistic++
				break
			}
		}
		log.Printf("Logisstic ::: %v\n", logistic)
		//Run(in, out, boxes, carsCount, carTonnage)
	}

}

func EquipCar(boxes []float64, tonnage float64) (boxesRemain []float64) {
	var offset int
	sort.Slice(boxes, func(i, j int) bool { return boxes[i] > boxes[j] })
	boxesRemain = make([]float64, len(boxes))
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

func Run(in *bufio.Reader, out *bufio.Writer, boxes []float64, carCount int, tonnage int) {

	//	::: Sort boxes
	log.Println(boxes)
	sort.Slice(boxes, func(i, j int) bool { return boxes[i] > boxes[j] })
	log.Println(boxes)
	boxesRemain := FindOptimusForCar(boxes, float64(tonnage))
	log.Println("boxesRemain:::", boxesRemain)
}

func FindOptimusForCar(boxes []float64, tonnage float64) (boxesrRemain []float64) {
	log.Println("BOXES", boxes)
	copy(boxesrRemain, boxes)
	sort.Slice(boxes, func(i, j int) bool { return boxes[i] > boxes[j] })
	for i, box := range boxes {
		if box == tonnage {
			if i == 0 {
				return boxesrRemain[1:]
			}
			return append(boxesrRemain[:i], boxesrRemain[i+1:]...)
		} else if tonnage > box {
			tonnage -= box
			if i == 0 {
				boxesrRemain = FindOptimusForCar(append(boxesrRemain[1:]), tonnage)
			} else {

				boxesrRemain = FindOptimusForCar(append(boxesrRemain[:i], boxesrRemain[i+1:]...), tonnage)
			}

		}

	}
	return boxesrRemain

}

/*

1
2 10
1 0 2 1 1 3 1



*/

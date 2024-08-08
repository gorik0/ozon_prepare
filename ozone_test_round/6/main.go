package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var (
		dataCount, boxesCount, carsCount, carTonnage int
		boxesMappa                                   map[int]int
		logistics                                    []int
		in                                           *bufio.Reader
		out                                          *bufio.Writer
	)

	in = bufio.NewReader(os.Stdin)
	logistics = make([]int, 0)
	out = bufio.NewWriter(os.Stdout)

	defer out.Flush()

	fmt.Fscan(in, &dataCount)
	//ti := time.Now()
	for range dataCount {
		fmt.Fscan(in, &carsCount, &carTonnage)

		fmt.Fscan(in, &boxesCount)
		boxesMappa = make(map[int]int)
		for range boxesCount {
			//fmt.Fscan(in, &boxes[i])
			var boxi int
			fmt.Fscan(in, &boxi)
			boxi = int(math.Pow(2, float64(boxi)))
			PushBoxi(boxi, boxesMappa)

		}

		//for i, el := range boxes {
		//	boxes[i] = int(math.Pow(2, float64(el)))
		//
		//}
		//println("before equip ___", time.Since(ti).Seconds())
		//remain := EquipCar(boxes, int(carTonnage))
		//log.Println("REMAIN :::: ", remain)
		logistic := 0
		carsRemain := carsCount

		//log.Println(boxesMappa)
		for {
			EquipCar(boxesMappa, int(carTonnage))
			carsRemain--
			if len(boxesMappa) != 0 && carsRemain == 0 {
				carsRemain = carsCount
				logistic++
			} else if len(boxesMappa) == 0 {
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
	//println(time.Since(ti).Seconds())

}

func PushBoxi(boxi int, mappa map[int]int) {
	//println("b", boxi)
	if boxCount, ok := mappa[boxi]; ok {
		//println("exist")
		mappa[boxi] = boxCount + 1
	} else {
		mappa[boxi] = 1
	}

}

func EquipCar(ma map[int]int, tonnage int) {

	boxeKeys := make([]int, 0)
	for key, _ := range ma {
		boxeKeys = append(boxeKeys, key)

	}
	sort.Sort(sort.Reverse(sort.IntSlice(boxeKeys)))
	for _, key := range boxeKeys {

		if key > tonnage {
			continue
		} else {
			tonnage -= key
			if tonnage == 0 {

				DeleteBoxFromMa(ma, key)
				return

			} else {
				DeleteBoxFromMa(ma, key)
				//:::: Try delete boxes of this weight
				if _, ok := ma[key]; ok {
					for tonnage >= key {
						tonnage -= key
						DeleteBoxFromMa(ma, key)
						if _, ok := ma[key]; !ok {
							break
						}
						if tonnage == 0 {
							return
						}
					}

				}
			}
		}

	}
	return
}

func DeleteBoxFromMa(ma map[int]int, key int) {
	if boxesCount, ok := ma[key]; ok {
		ma[key] = boxesCount - 1
		if boxesCount == 1 {
			delete(ma, key)
		}

	}

}

/*

	1
	1 10
2
	3 1

	1
1 20
7
3 1 3 1 3 1 0

1
1 41
6
5 5 3 0 2 1



*/

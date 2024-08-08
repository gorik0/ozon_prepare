package main

import "log"

func main() {

	a := []int{1, 2, 3, 4, 5}
	log.Println(a[1:])
}

func Changeslice(a []int) {
	println("ssssssss")
	println(a)
	a[0] = 99999
	println(a)
	a = append(a[:1], a[2])
	println(a)
	a[0] = 0
	log.Println(a)
}

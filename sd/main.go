package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	var a int
	//scanner := bufio.NewScanner(os.Stdin)
	fmt.Fscanln(in, &a)
	for range a {
		line, _, _ := in.ReadLine()
		println("line:", string(line))

	}

}

func deleteFromMa(ma []interface{}) {
	ma = append(ma[:1], ma[2])
}

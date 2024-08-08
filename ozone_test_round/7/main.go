package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		newJobers, oldJobers int
		newLogins, oldLogins []string
		in                   *bufio.Reader
		out                  *bufio.Writer
	)

	_ = newJobers
	_ = newLogins
	// IN--OUT
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// IN--OUT

	// scan OLD
	fmt.Fscan(in, &oldJobers)
	oldLogins = make([]string, oldJobers)
	for i := range oldJobers {
		fmt.Fscan(in, &oldLogins[i])
	}
	// scan OLD

	// scan NEW
	fmt.Fscan(in, &newJobers)
	newLogins = make([]string, newJobers)
	for i := range newJobers {
		fmt.Fscan(in, &newLogins[i])
	}
	// scan NEW

	log.Println("oldLogins:", oldLogins)
	log.Println("newLogins:", newLogins)

	var l1, l2 string = "bacccab", "abcccba"

	log.Printf("l1 ::: %s\n", l1)
	log.Printf("l2 ::: %s\n", l2)
	log.Println(LoginsIsEq(l1, l2))

}

func LoginsIsEq(l1 string, l2 string) bool {
	if l1 == l2 {
		return true
	} else {
		if len(l1) != len(l2) {
			return false
		}
		return LoginsIsEqWithReplace(l1, l2)
	}
}

func LoginsIsEqWithReplace(l1 string, l2 string) bool {
	for i := range l1 {

		if l1[i] != l2[i] {
			if i == len(l1)-1 {
				println(1)
				println("i :::", i)
				return false
			} else {

				if i+2 != len(l1) {
					log.Println(2)
					println("i :::", i)

					return l1[i] == l2[i+1] && l2[i] == l1[i+1] && l2[i+2:] == l1[i+2:]
				} else {
					println(3)
					println("i :::", i)

					return l1[i] == l2[i+1] && l2[i] == l1[i+1]

				}
			}
		}
	}
	return false
}

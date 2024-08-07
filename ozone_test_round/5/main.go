package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	querlyBracesOpen  = "{"
	querlyBracesClose = "}"
	squaBracesOpen    = "["
	squaBracesClose   = "]"
)

func main() {
	var (
		jsonLines, dataCount int
		in                   *bufio.Reader
		out                  *bufio.Writer
		stringJson           string
		//open, close        int
		jso      map[string]interface{}
		jsoArray []interface{}
	)
	jso = map[string]interface{}{}
	jsoArray = []interface{}{}
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	//	throughout data
	fmt.Fscan(in, &dataCount)
	for range dataCount {
		fmt.Fscan(in, &jsonLines)

		for jsonLines > 0 {

			var stri string
			stri, _ = in.ReadString('\n')

			stringJson = stringJson + strings.Trim(stri, "\n")
			if stri != "\n" {
				jsonLines--
			}
		}
	}
	println(stringJson)
	//stringJson = `[{}, [], {}, 	{}, "string"]`
	//Handle(stringJson)
	err := json.Unmarshal([]byte(stringJson), &jso)

	if err != nil {
		println("no map ...")

		if err = json.Unmarshal([]byte(stringJson), &jsoArray); err != nil {

			println("no list too")
			return
		}
		throughList(jsoArray)
		log.Println(jsoArray)
		return
	}
	throughMap(jso)
	println(jso)
	//log.Printf("HSON ::: %v", jso)

}

func throughList(li []interface{}) bool {
	isEmpty := true

	for _, i := range li {
		log.Println("I ::: ", i)
		//	::: is list
		if liSmall, ok := i.([]interface{}); ok {
			if throughList(liSmall) {
				println("emptty")
				//li = append(li[:pos], li[pos+1:]...)
			}
		} else if mapSmall, ok := i.(map[string]interface{}); ok {
			if throughMap(mapSmall) {
				println("emptty")
				//li = append(li[:pos], li[pos+1:]...)
			}
		}
		isEmpty = false
	}
	return isEmpty
}
func throughMap(ma map[string]interface{}) bool {
	isEmpty := true
	for _, i := range ma {

		log.Println("I ::: ", i)
		if liSmall, ok := i.([]interface{}); ok {
			if throughList(liSmall) {
				println("empty ..")
			}
		} else if mapSmall, ok := i.(map[string]interface{}); ok {
			if throughMap(mapSmall) {
				println("empty ..")
			}
		}
		isEmpty = false
	}
	return isEmpty
}

/*
1
2
[{}, [1,[],true], {}, 	{}, "string"
]


*/

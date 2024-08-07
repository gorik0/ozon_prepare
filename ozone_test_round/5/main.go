package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
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
		jso             map[string]interface{}
		jsoArray, jsons []interface{}
	)
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	//	throughout data
	line, _, _ := in.ReadLine()
	dataCount, _ = strconv.Atoi(string(line))
	for range dataCount {
		jsoArray = []interface{}{}
		jso = map[string]interface{}{}
		stringJson = ""

		line, _, _ := in.ReadLine()
		jsonLines, _ = strconv.Atoi(string(line))
		log.Println(jsonLines)
		for range jsonLines {

			var stri string
			line, _, _ := in.ReadLine()
			stri = string(line)

			stringJson = stringJson + strings.Trim(stri, "\n")
		}

		err := json.Unmarshal([]byte(stringJson), &jso)

		if err != nil {

			if err = json.Unmarshal([]byte(stringJson), &jsoArray); err != nil {

				return
			}
			_, offset := throughList(jsoArray)

			jsons = append(jsons, jsoArray[:len(jsoArray)-offset])
			continue

		}
		throughMap(jso)

		jsons = append(jsons, jso)

	}
	jj, _ := json.Marshal(jsons)
	fmt.Fprintf(out, "%v", string(jj))

}

func throughList(li []interface{}) (bool, int) {
	isEmpty := true
	var offset int
	var newLi []interface{} = make([]interface{}, len(li))
	copy(newLi, li)
	for pos, i := range li {
		//	::: is list
		if liSmall, ok := i.([]interface{}); ok {
			if empty, _ := throughList(liSmall); empty {

				newLi = append(newLi[:pos-offset], newLi[pos+1-offset:]...)
				offset++
			}
		} else if mapSmall, ok := i.(map[string]interface{}); ok {
			if throughMap(mapSmall) {
				newLi = append(newLi[:pos-offset], newLi[pos+1-offset:]...)
				offset++
			}
		}
		isEmpty = false
	}
	copy(li, newLi)

	return isEmpty, offset
}
func throughMap(ma map[string]interface{}) bool {
	isEmpty := true
	for key, i := range ma {

		if liSmall, ok := i.([]interface{}); ok {
			if empty, _ := throughList(liSmall); empty {
				delete(ma, key)

			} else {

				isEmpty = false
			}
		} else if mapSmall, ok := i.(map[string]interface{}); ok {
			if throughMap(mapSmall) {
				delete(ma, key)

			} else {
				isEmpty = false

			}
		} else {

			isEmpty = false
		}
	}
	return isEmpty
}

/*


3
6
{



"a": "f",
"b": {"c": {"d": [], "e": ["ababa"]}},
"c": {"k": {}},
"d": {"d": {"e": {}}}
}
2
[{}, [], {}, 	{}, "string"
]
3
[{"one":
	[{"two":
		[{"three":"four"}]}]}]



*/

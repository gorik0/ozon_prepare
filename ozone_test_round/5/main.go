package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
		for range jsonLines {

			var stri string
			line, _, _ := in.ReadLine()
			stri = string(line)

			stringJson = stringJson + strings.Trim(stri, "\n")
		}
		//stringJson = `[[[],["fqj"],"ffo"]]`

		err := json.Unmarshal([]byte(stringJson), &jso)

		if err != nil {

			if err = json.Unmarshal([]byte(stringJson), &jsoArray); err != nil {

				return
			}
			l := throughList(jsoArray)

			jsons = append(jsons, l)
			continue

		}
		throughMap(jso)

		jsons = append(jsons, jso)

	}
	jj, _ := json.Marshal(jsons)
	fmt.Fprintf(out, "%v", string(jj))

}

func throughList(li []interface{}) []interface{} {

	var newLi []interface{} = make([]interface{}, len(li))
	copy(newLi, li)
	var isEmpty = true
	var offset int
	for pos, i := range li {
		//	::: is list
		if liSmall, ok := i.([]interface{}); ok {
			liSmall = throughList(liSmall)

			if liSmall == nil {
				newLi = append(newLi[:pos-offset], newLi[pos+1-offset:]...)

				offset++

			} else {
				newLi[pos-offset] = liSmall
				isEmpty = false
			}
		} else if mapSmall, ok := i.(map[string]interface{}); ok {
			if throughMap(mapSmall) {
				newLi = append(newLi[:pos-offset], newLi[pos+1-offset:]...)
				offset++
			} else {
				isEmpty = false
			}
		} else {
			isEmpty = false
		}

	}
	if isEmpty {
		return nil
	}
	return newLi
}
func throughMap(ma map[string]interface{}) bool {
	isEmpty := true
	for key, i := range ma {

		if liSmall, ok := i.([]interface{}); ok {
			liSmall = throughList(liSmall)
			ma[key] = liSmall
			if liSmall == nil {
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

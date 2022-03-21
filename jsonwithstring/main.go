package main

import (
	"fmt"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/json"
)

var myjsonstr = `{
	"name":"John",
	"age":30,
	"cars":[ "Ford", "BMW", "Fiat" ]
	}`

func main() {
	hogosuru.Init()

	if j, err := json.Parse(myjsonstr); hogosuru.AssertErr(err) {

		extractdata := j.Map()
		if arrayincars, ok := extractdata.(map[string]interface{})["cars"].([]interface{}); ok {

			if len(arrayincars) > 0 {
				//ensure that the first element is a string
				if valuestr, ok := arrayincars[0].(string); ok {
					fmt.Printf("First Element in cars is %s\n", valuestr)
				}

			}
		}

		if arrayincars, ok := extractdata.(map[string]interface{})["boat"].([]string); ok {

			if len(arrayincars) > 0 {
				fmt.Printf("First Element in boat is %s\n", arrayincars[0])
			}
		} else {
			fmt.Printf("Boat is not present in json\n")
		}
	}

}

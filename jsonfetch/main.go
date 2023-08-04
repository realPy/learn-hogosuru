package main

import (
	"fmt"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/fetch"
	"github.com/realPy/hogosuru/base/json"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/response"
)

func main() {
	hogosuru.Init()

	var httpHeaders map[string]interface{} = map[string]interface{}{"Content-Type": "application/json"}
	var fetchOpts map[string]interface{} = map[string]interface{}{"method": "GET", "headers": httpHeaders}

	//Start promise and wait result
	if f, err := fetch.New("mockjson/get.json", fetchOpts); hogosuru.AssertErr(err) {
		f.Then(func(r response.Response) *promise.Promise {
			//when header is ready the function is execute
			if status, err := r.Status(); hogosuru.AssertErr(err) {
				if status == 200 {
					if promiseGetJSONData, err := r.Text(); hogosuru.AssertErr(err) {
						//when data is ready the text is get
						promiseGetJSONData.Then(func(i interface{}) *promise.Promise {
							//we must ensure that the object send is a string like expected
							if textstr, ok := i.(string); ok {
								if j, err := json.Parse(textstr); hogosuru.AssertErr(err) {
									jsonmap := j.Map()

									if headers, ok := jsonmap.(map[string]interface{})["headers"]; ok {

										if uaheader, ok := headers.(map[string]interface{})["User-Agent"].(string); ok {
											fmt.Printf("User-Agent %s\n", uaheader)
										}

									}

								}

							}
							return nil
						}, func(e error) {
							fmt.Printf("An error occurs %s\n", err.Error())
						})

					}
				} else {

					fmt.Printf("Must return a 200 HTTP Code")

				}
			}
			return nil
		}, func(e error) {
			fmt.Printf("An error occurs %s\n", err.Error())

		})
	}

	ch := make(chan struct{})
	<-ch

}

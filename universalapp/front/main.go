package main

import (
	"fmt"
	"net/url"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/fetch"
	"github.com/realPy/hogosuru/htmlbuttonelement"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmllinkelement"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/response"
)

func OnClickButton() {

	dataPost := url.Values{}
	dataPost.Set("username", "testfetch")

	var headers map[string]interface{} = map[string]interface{}{"Content-Type": "application/x-www-form-urlencoded"}
	var fetchOpts map[string]interface{} = map[string]interface{}{"method": "POST", "headers": headers, "body": dataPost.Encode()}

	//Start promise and wait result
	if f, err := fetch.New("app/data", fetchOpts); hogosuru.AssertErr(err) {
		f.Then(func(r response.Response) *promise.Promise {

			return nil
		}, func(e error) {

			fmt.Printf("An error occured: %s\n", e.Error())
		})

	}

}

func main() {
	hogosuru.Init()

	// we get the current document
	if doc, err := document.New(); hogosuru.AssertErr(err) {

		//we got the head
		if head, err := doc.Head(); hogosuru.AssertErr(err) {

			if link, err := htmllinkelement.New(doc); hogosuru.AssertErr(err) {

				link.SetRel("stylesheet")
				link.SetHref("https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css")
				head.AppendChild(link.Node)

			}

		}

		//we get the current body
		if body := doc.Body_(); hogosuru.AssertErr(err) {

			if div, err := htmldivelement.New(doc); hogosuru.AssertErr(err) {

				if list, err := div.ClassList(); hogosuru.AssertErr(err) {
					list.Add("buttons")
				}

				if buttonprimary, err := htmlbuttonelement.New(doc); hogosuru.AssertErr(err) {

					buttonprimary.SetTextContent("Primary")
					//we get the class list attribute
					if list, err := buttonprimary.ClassList(); hogosuru.AssertErr(err) {
						list.Add("button")
						list.Add("is-primary")
					}

					buttonprimary.OnClick(func(e event.Event) {
						OnClickButton()
					})

					div.Append(buttonprimary.Element)
				}

				body.Append(div.Element)

			}
		}
	}

	ch := make(chan struct{})
	<-ch

}

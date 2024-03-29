package main

import (
	"fmt"
	"net/url"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/formdata"
	"github.com/realPy/hogosuru/base/htmlbuttonelement"
	"github.com/realPy/hogosuru/base/htmldivelement"
	"github.com/realPy/hogosuru/base/htmllinkelement"
	"github.com/realPy/hogosuru/base/xmlhttprequest"
)

func OnClickButton() {

	endpoint, _ := url.Parse("app/data")

	if xhr, err := xmlhttprequest.New(); err == nil {

		xhr.Open("POST", endpoint.String())
		f, _ := formdata.New()

		f.Append("username", "test")

		xhr.SetOnload(func(i interface{}) {

			if text, err := xhr.ResponseText(); err == nil {
				fmt.Printf("Resultat: %s\n", text)
			}

			if header, err := xhr.GetResponseHeader("Content-Type"); err == nil {
				fmt.Printf("Resultat: %s\n", header)
			}

		})
		xhr.Send(f)

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

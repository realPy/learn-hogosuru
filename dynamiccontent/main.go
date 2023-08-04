package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/console"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/htmlbuttonelement"
	"github.com/realPy/hogosuru/base/htmldivelement"
	"github.com/realPy/hogosuru/base/htmlspanelement"
)

func main() {
	hogosuru.Init()

	c, _ := console.New()

	// we get the current document
	if doc, err := document.New(); hogosuru.AssertErr(err) {

		//we get the current body
		if body := doc.Body_(); hogosuru.AssertErr(err) {

			//we empty all things in body
			body.SetTextContent("")

			//we create a div
			if div, err := htmldivelement.New(doc); hogosuru.AssertErr(err) {

				//we create a button and a span with text
				if span, err := htmlspanelement.New(doc); hogosuru.AssertErr(err) {
					span.SetTextContent("Please click on this button")
					span.SetID("spanwithtext")
					div.Append(span.Element)
				}

				//we create a button and a span with text
				if button, err := htmlbuttonelement.New(doc); hogosuru.AssertErr(err) {
					button.SetTextContent("click on this button")
					button.SetID("mycustombutton")
					button.OnClick(func(e event.Event) {

						c.Debug("You click on this button")
					})
					div.Append(button.Element)
				}

				//we attach div to the body

				body.Append(div.Element)

			}

		}

	}

	ch := make(chan struct{})
	<-ch

}

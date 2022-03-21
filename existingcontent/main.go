package main

import (
	"errors"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/console"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmlbuttonelement"
	"github.com/realPy/hogosuru/htmlspanelement"
	"github.com/realPy/hogosuru/node"
)

var domhtml = `<div>
<span id="myspantextid"></span>
<button id="mybuttonid"></button>
</div>`

func GetSpanByID(elemsearch element.Element, ID string) (htmlspanelement.HtmlSpanElement, error) {
	var e htmlspanelement.HtmlSpanElement
	var err error
	var elem node.Node
	var elemInstance interface{}
	var ok bool
	if elem, err = elemsearch.QuerySelector(ID); hogosuru.AssertErr(err) {

		if elemInstance, err = elem.Discover(); hogosuru.AssertErr(err) {

			if e, ok = elemInstance.(htmlspanelement.HtmlSpanElement); !ok {
				err = errors.New(ID + " is not a span")
			}
		}

	}
	return e, err
}

func GetButtonByID(elemsearch element.Element, ID string) (htmlbuttonelement.HtmlButtonElement, error) {
	var e htmlbuttonelement.HtmlButtonElement
	var err error
	var elem node.Node
	var elemInstance interface{}
	var ok bool
	if elem, err = elemsearch.QuerySelector(ID); hogosuru.AssertErr(err) {

		if elemInstance, err = elem.Discover(); hogosuru.AssertErr(err) {

			if e, ok = elemInstance.(htmlbuttonelement.HtmlButtonElement); !ok {
				err = errors.New(ID + " is not a button")
			}
		}

	}
	return e, err
}

func main() {
	hogosuru.Init()

	c, _ := console.New()

	// we get the current document
	if doc, err := document.New(); hogosuru.AssertErr(err) {

		//we get the current body
		if body := doc.Body_(); hogosuru.AssertErr(err) {
			//be carefull use a sanytiser to prevent xss if you're not trust the src
			body.SetInnerHTML(domhtml)

			//search the element name "myspantextid" in the body

			if span, err := GetSpanByID(body.Element, "#myspantextid"); hogosuru.AssertErr(err) {

				span.SetTextContent("Please click here :)")
			}

			if button, err := GetButtonByID(body.Element, "#mybuttonid"); hogosuru.AssertErr(err) {

				button.SetTextContent("The text button content")
				button.OnClick(func(e event.Event) {
					c.Debug("You click on this button")
				})
			}

		}

	}

	ch := make(chan struct{})
	<-ch

}

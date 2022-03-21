package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmlbuttonelement"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmllinkelement"
)

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

		if body, err := doc.Body(); hogosuru.AssertErr(err) {
			//lets create some button design with bulma

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

						if list, err := buttonprimary.ClassList(); hogosuru.AssertErr(err) {

							list.Remove("is-primary")
							list.Add("is-warning")
						}
					})

					div.Append(buttonprimary.Element)
				}

				if buttondanger, err := htmlbuttonelement.New(doc); hogosuru.AssertErr(err) {

					buttondanger.SetTextContent("Danger")
					//we get the class list attribute
					if list, err := buttondanger.ClassList(); hogosuru.AssertErr(err) {
						list.Add("button")
						list.Add("is-danger")
					}

					buttondanger.OnClick(func(e event.Event) {

						if list, err := buttondanger.ClassList(); hogosuru.AssertErr(err) {

							list.Remove("is-danger")
							list.Add("is-info")
						}
					})
					div.Append(buttondanger.Element)
				}

				body.Append(div.Element)

			}

		}

	}

	ch := make(chan struct{})
	<-ch

}

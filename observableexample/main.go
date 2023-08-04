package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/htmlbuttonelement"
	"github.com/realPy/hogosuru/base/htmldivelement"
	"github.com/realPy/hogosuru/base/htmllinkelement"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/keyobservable"
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

				if buttondanger, err := htmlbuttonelement.New(doc); hogosuru.AssertErr(err) {

					buttondanger.SetTextContent("Danger wait 5s....")
					//we get the class list attribute
					if list, err := buttondanger.ClassList(); hogosuru.AssertErr(err) {
						list.Add("button")
						list.Add("is-danger")
					}

					keyobservable.KeyObservable().RegisterFunc("dangertextbutton", func(value interface{}) {

						if textbutton, ok := value.(string); ok {
							buttondanger.SetTextContent(textbutton)
						}

					})
					//we register the upgrade text content

					div.Append(buttondanger.Element)
				}

				body.Append(div.Element)

			}

		}

	}

	//we create a promise wait 5 second and set the "dangertextbutton" key to another value,
	if w1, err := promise.SetTimeout(5000); hogosuru.AssertErr(err) {
		w1.Then(func(i interface{}) *promise.Promise {

			keyobservable.KeyObservable().Set("dangertextbutton", "Boom!", false)

			return nil
		}, nil)
	}

	ch := make(chan struct{})
	<-ch

}

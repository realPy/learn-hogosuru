package main

import (
	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/htmlanchorelement"
	"github.com/realPy/hogosuru/htmllinkelement"
)

var domtemplate = `<div class="dropdown" id="dropdowntest">
  <div class="dropdown-trigger">
    <button class="button" aria-haspopup="true" aria-controls="dropdown-menu2" id="dropdownbutton">
      <span>Click me</span>
	  <span class="icon is-small">
	  <i class="fas fa-angle-down" aria-hidden="true"></i>
	  </span>
    </button>
  </div>
  <div class="dropdown-menu" id="dropdown-menu2" role="menu">
    <div class="dropdown-content">
	  <template id="mytemplatedropdown">
	  <a href="#"  class="dropdown-item" id="itemtpl">
		  <span>Hello </span><span id="itemtext"></span>
      </a>
	  </template>
	  <div id="dropdownitemscontent"></div>
    </div>
  </div>
</div>`

var arraydynamiccontent []string = []string{"World", "Me", "Cat"}

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

			if link, err := htmllinkelement.New(doc); hogosuru.AssertErr(err) {

				link.SetRel("stylesheet")
				link.SetHref("https://bulma.io/vendor/fontawesome-free-5.15.2-web/css/all.min.css")
				head.AppendChild(link.Node)

			}

		}

		if body, err := doc.Body(); hogosuru.AssertErr(err) {
			//we insert the html code
			body.InsertAdjacentHTML("beforeend", domtemplate)

			if dropdowntest, err := GetDivBySelector(body.Element, "#dropdowntest"); hogosuru.AssertErr(err) {
				if dropdownbutton, err := GetButtonBySelector(body.Element, "#dropdownbutton"); hogosuru.AssertErr(err) {
					dropdownbutton.OnClick(func(e event.Event) {
						//when click we open or close the drop down
						if list, err := dropdowntest.ClassList(); hogosuru.AssertErr(err) {
							list.Toggle("is-active")
						}

					})
				}

			}

			if contentitems, err := GetDivBySelector(body.Element, "#dropdownitemscontent"); hogosuru.AssertErr(err) {
				contentitems.SetTextContent("")
				if t, err := GetTemplateBySelector(body.Element, "#mytemplatedropdown"); hogosuru.AssertErr(err) {
					//now we want the content of the template to duplicate it
					if fragment, err := t.Content(); hogosuru.AssertErr(err) {
						if cloneNode, err := fragment.GetElementById("itemtpl"); hogosuru.AssertErr(err) {

							//we loop for each content
							for _, name := range arraydynamiccontent {
								if clone, err := doc.ImportNode(cloneNode.Node, true); hogosuru.AssertErr(err) {
									if a, ok := clone.(htmlanchorelement.HtmlAnchorElement); ok {

										//we search the span txtContent to modify value

										if spantxt, err := GetSpanBySelector(a.Element, "#itemtext"); hogosuru.AssertErr(err) {
											spantxt.SetTextContent(name)
										}

										contentitems.Append(a.Element)

									}
								}
							}

						}

					}
				}

			}

			//we search element like contents to loop the data and template

		}

	}

	ch := make(chan struct{})
	<-ch

}

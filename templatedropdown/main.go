package main

import (
	"errors"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/htmlanchorelement"
	"github.com/realPy/hogosuru/base/htmlbuttonelement"
	"github.com/realPy/hogosuru/base/htmldivelement"
	"github.com/realPy/hogosuru/base/htmllinkelement"
	"github.com/realPy/hogosuru/base/htmlspanelement"
	"github.com/realPy/hogosuru/base/htmltemplateelement"
	"github.com/realPy/hogosuru/base/node"
	"github.com/realPy/hogosuru/htmlstruct"
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

type DropDownTemplate struct {
	Item htmlanchorelement.HtmlAnchorElement `hogosuru:"#itemtpl"`
}

type DropDownTitem struct {
	Item htmlspanelement.HtmlSpanElement `hogosuru:"#itemtext"`
}

type MainWindow struct {
	DropdownTest         htmldivelement.HtmlDivElement           `hogosuru:"#dropdowntest"`
	DropdownButton       htmlbuttonelement.HtmlButtonElement     `hogosuru:"#dropdownbutton"`
	DropdownItemsContent htmldivelement.HtmlDivElement           `hogosuru:"#dropdownitemscontent"`
	Mytemplatedropdown   htmltemplateelement.HtmlTemplateElement `hogosuru:"#mytemplatedropdown"`
	Itemtemplate         DropDownTemplate
}

func ClonableStruct(doc document.Document, root node.Node, i interface{}) (element.Element, error) {
	clone, err := doc.ImportNode(root.Node_(), true)
	if err != nil {
		return element.Element{}, err
	}
	el, ok := clone.(element.ElementFrom)
	if !ok {
		return element.Element{}, errors.New("can't clone struct: not an element")
	}

	clonelement := el.Element_()
	err = htmlstruct.Unmarshal(clonelement, i)
	if err != nil {
		panic(err)
	}
	return clonelement, err
}

func main() {
	hogosuru.Init()
	var w MainWindow
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
			htmlstruct.Unmarshal(doc, &w)

			w.DropdownButton.OnClick(func(e event.Event) {
				//when click we open or close the drop down
				if list, err := w.DropdownTest.ClassList(); hogosuru.AssertErr(err) {
					list.Toggle("is-active")
				}

			})

			w.DropdownItemsContent.SetTextContent("")

			if fragment, err := w.Mytemplatedropdown.Content(); hogosuru.AssertErr(err) {
				htmlstruct.Unmarshal(fragment, &w.Itemtemplate)

				for _, name := range arraydynamiccontent {
					var item DropDownTitem
					rootnode, err := ClonableStruct(doc, w.Itemtemplate.Item.Node, &item)
					if err == nil {
						item.Item.SetTextContent(name)
						w.DropdownItemsContent.Append(rootnode)
					}
				}

			}

		}

	}

	ch := make(chan struct{})
	<-ch

}

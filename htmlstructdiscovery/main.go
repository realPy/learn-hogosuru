package main

import (
	"fmt"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/fetch"
	"github.com/realPy/hogosuru/base/htmlbodyelement"
	"github.com/realPy/hogosuru/base/htmlbuttonelement"
	"github.com/realPy/hogosuru/base/htmldivelement"
	"github.com/realPy/hogosuru/base/htmlheadingelement"
	"github.com/realPy/hogosuru/base/node"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/response"
	"github.com/realPy/hogosuru/hogosurudebug"
	"github.com/realPy/hogosuru/htmlstruct"
	"github.com/realPy/hogosuru/routing"
)

type MyMainWindow struct {
	Body        htmlbodyelement.HtmlBodyElement         `hogosuru:"body:nth-of-type(1)"`
	Divoups     htmldivelement.HtmlDivElement           `hogosuru:"#oups"`
	Divs        []htmldivelement.HtmlDivElement         `hogosuru:"[]"`
	H1s         []htmlheadingelement.HtmlHeadingElement `hogosuru:"[]:1"`
	PressButton htmlbuttonelement.HtmlButtonElement     `hogosuru:"#press"`
	P2          htmlbuttonelement.HtmlButtonElement     `hogosuru:"button.inner"`
	Divtoto     []htmldivelement.HtmlDivElement         `hogosuru:"[]div.toto"`
}

// //////////////////////////::
type GlobalContainer struct {
	global element.Element
}

func (w *GlobalContainer) OnFinishedLoaded() {

}

func (w *GlobalContainer) Node(r routing.Rendering) node.Node {

	return w.global.Node_()
}

func (w *GlobalContainer) OnEndChildRendering(r routing.Rendering) {

}

func (w *GlobalContainer) OnEndChildsRendering() {

}

func (w *GlobalContainer) OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []routing.Rendering) {

	var ret *promise.Promise

	if f, err := fetch.New("main.html"); hogosuru.AssertErr(err) {
		textpromise, _ := f.Then(func(r response.Response) *promise.Promise {

			if promise, err := r.Text(); hogosuru.AssertErr(err) {
				return &promise
			}

			return nil

		}, nil)

		textpromise.Then(func(i interface{}) *promise.Promise {

			if element, err := d.DocumentElement(); hogosuru.AssertErr(err) {

				element.SetInnerHTML(i.(string))

				var w MyMainWindow

				htmlstruct.Unmarshal(d, &w)

				for _, div := range w.Divs {
					div.Style_().SetProperty("background-color", "lightgoldenrodyellow")
				}

				for _, h1 := range w.H1s {
					h1.Style_().SetProperty("background-color", "red")
				}

				w.PressButton.OnClick(func(e event.Event) {
					fmt.Printf("Hello world\n")
				})
				w.P2.OnClick(func(e event.Event) {
					fmt.Printf("InnerBox\n")
				})

				for _, div := range w.Divtoto {
					div.Style_().SetProperty("background-color", "blue")
				}

			}
			return nil
		}, nil)

		ret = &textpromise

	}

	return ret, []routing.Rendering{}
}

func (w *GlobalContainer) OnUnload() {

}

func main() {

	hogosuru.Init()
	hogosurudebug.EnableDebug()

	routing.Router().DefaultRendering(&GlobalContainer{})

	routing.Router().Start(routing.HASHROUTE)

	ch := make(chan struct{})
	<-ch

}

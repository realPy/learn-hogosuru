package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/blob"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/dragevent"
	"github.com/realPy/hogosuru/file"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmllinkelement"
	"github.com/realPy/hogosuru/promise"
)

type HashCalculate struct {
	name string
	hash string
}

func sha256File(f file.File) {
	var sha256result string

	//allocate memory in handler is not recommend
	var buffer []byte = make([]byte, 128*1024)
	hashsha256 := sha256.New()

	stream := blob.NewBlobStream(f.Blob)

	p, _ := stream.AsyncRead(buffer, func(b []byte, i int) {
		hashsha256.Write(b[:i])
	})

	p.Then(func(i interface{}) *promise.Promise {

		sha256result = hex.EncodeToString(hashsha256.Sum(nil))

		hogosuru.KeyObservable().Set("newhash", HashCalculate{name: f.Name_(), hash: sha256result}, false)
		return nil
	}, nil)
}

func main() {

	hogosuru.Init()

	if doc, err := document.New(); hogosuru.AssertErr(err) {

		//we got the head
		if head, err := doc.Head(); hogosuru.AssertErr(err) {

			if link, err := htmllinkelement.New(doc); hogosuru.AssertErr(err) {

				link.SetRel("stylesheet")
				link.SetHref("https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css")
				head.AppendChild(link.Node)

			}

		}

		doc.OnDragOver(func(e dragevent.DragEvent) {

			e.PreventDefault()

		})

		if body, err := doc.Body(); hogosuru.AssertErr(err) {

			if div, err := htmldivelement.New(doc); hogosuru.AssertErr(err) {

				if list, err := div.ClassList(); hogosuru.AssertErr(err) {
					list.Add("box")

				}

				div.SetTextContent("Drag files here")
				div.OnDrop(func(e dragevent.DragEvent) {

					e.PreventDefault()
					if dt, err := e.DataTransfer(); hogosuru.AssertErr(err) {
						if files, err := dt.Files(); hogosuru.AssertErr(err) {
							if l, err := files.Length(); hogosuru.AssertErr(err) {
								for i := 0; i < l; i++ {
									if f, err := files.Item(i); hogosuru.AssertErr(err) {

										sha256File(f)
									}
								}
							}

						}
					}

				})

				hogosuru.KeyObservable().RegisterFunc("newhash", func(value interface{}) {

					if h, ok := value.(HashCalculate); ok {
						if div, err := htmldivelement.New(doc); hogosuru.AssertErr(err) {
							div.SetTextContent(fmt.Sprintf("%s: %s", h.name, h.hash))
							body.Append(div.Element)
						}

					}

				})

				body.Append(div.Element)

			}

		}

	}

	ch := make(chan struct{})
	<-ch

}

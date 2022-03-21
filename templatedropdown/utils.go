package main

import (
	"errors"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlanchorelement"
	"github.com/realPy/hogosuru/htmlbuttonelement"
	"github.com/realPy/hogosuru/htmldivelement"
	"github.com/realPy/hogosuru/htmlspanelement"
	"github.com/realPy/hogosuru/htmltemplateelement"
	"github.com/realPy/hogosuru/node"
)

func GetTemplateBySelector(elemsearch element.Element, selector string) (htmltemplateelement.HtmlTemplateElement, error) {
	var e htmltemplateelement.HtmlTemplateElement
	var err error
	var elem node.Node
	var elemInstance interface{}
	var ok bool
	if elem, err = elemsearch.QuerySelector(selector); hogosuru.AssertErr(err) {

		if elemInstance, err = elem.Discover(); hogosuru.AssertErr(err) {

			if e, ok = elemInstance.(htmltemplateelement.HtmlTemplateElement); !ok {
				err = errors.New(selector + " is not a template")
			}
		}

	}
	return e, err
}

func GetDivBySelector(elemsearch element.Element, selector string) (htmldivelement.HtmlDivElement, error) {
	var e htmldivelement.HtmlDivElement
	var err error
	var elem node.Node
	var elemInstance interface{}
	var ok bool
	if elem, err = elemsearch.QuerySelector(selector); hogosuru.AssertErr(err) {

		if elemInstance, err = elem.Discover(); hogosuru.AssertErr(err) {

			if e, ok = elemInstance.(htmldivelement.HtmlDivElement); !ok {
				err = errors.New(selector + " is not a div")
			}
		}

	}
	return e, err
}

func GetSpanBySelector(elemsearch element.Element, selector string) (htmlspanelement.HtmlSpanElement, error) {
	var e htmlspanelement.HtmlSpanElement
	var err error
	var elem node.Node
	var elemInstance interface{}
	var ok bool
	if elem, err = elemsearch.QuerySelector(selector); hogosuru.AssertErr(err) {

		if elemInstance, err = elem.Discover(); hogosuru.AssertErr(err) {

			if e, ok = elemInstance.(htmlspanelement.HtmlSpanElement); !ok {
				err = errors.New(selector + " is not a span")
			}
		}

	}
	return e, err
}

func GetButtonBySelector(elemsearch element.Element, selector string) (htmlbuttonelement.HtmlButtonElement, error) {
	var e htmlbuttonelement.HtmlButtonElement
	var err error
	var elem node.Node
	var elemInstance interface{}
	var ok bool
	if elem, err = elemsearch.QuerySelector(selector); hogosuru.AssertErr(err) {

		if elemInstance, err = elem.Discover(); hogosuru.AssertErr(err) {

			if e, ok = elemInstance.(htmlbuttonelement.HtmlButtonElement); !ok {
				err = errors.New(selector + " is not a button")
			}
		}

	}
	return e, err
}

func GetAnchorBySelector(elemsearch element.Element, selector string) (htmlanchorelement.HtmlAnchorElement, error) {
	var e htmlanchorelement.HtmlAnchorElement
	var err error
	var elem node.Node
	var elemInstance interface{}
	var ok bool
	if elem, err = elemsearch.QuerySelector(selector); hogosuru.AssertErr(err) {

		if elemInstance, err = elem.Discover(); hogosuru.AssertErr(err) {

			if e, ok = elemInstance.(htmlanchorelement.HtmlAnchorElement); !ok {
				err = errors.New(selector + " is not an anchor")
			}
		}

	}
	return e, err
}

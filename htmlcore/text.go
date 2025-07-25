// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package htmlcore

import (
	"slices"

	"golang.org/x/net/html"
)

// ExtractText recursively extracts all of the text from the children
// of the given [*html.Node], adding any appropriate inline markup for
// formatted text. It adds any non-text elements to the given [core.Widget]
// using [readHTMLNode]. It should not be called on text nodes themselves;
// for that, you can directly access the [html.Node.Data] field. It uses
// the given page URL for context when resolving URLs, but it can be
// omitted if not available.
func ExtractText(ctx *Context) string {
	if ctx.Node.FirstChild == nil {
		return ""
	}
	tx, _ := extractTextExclude(ctx, ctx.Node.FirstChild)
	return tx
}

// ExtractTextExclude recursively extracts all of the text from the children
// of the given [*html.Node], adding any appropriate inline markup for
// formatted text. It adds any non-text elements to the given [core.Widget]
// using [readHTMLNode], except those types listed in the excludeSubs list.
// If one of those types is encountered, it is returned.
func ExtractTextExclude(ctx *Context, excludeSubs ...string) (string, *html.Node) {
	if ctx.Node.FirstChild == nil {
		return "", nil
	}
	return extractTextExclude(ctx, ctx.Node.FirstChild, excludeSubs...)
}

func extractTextExclude(ctx *Context, n *html.Node, excludeSubs ...string) (string, *html.Node) {
	str := ""
	if n.Type == html.TextNode {
		str += n.Data
	}
	it := isText(n)
	if !it {
		if slices.Contains(excludeSubs, n.Data) {
			return str, n
		}
		readHTMLNode(ctx, ctx.Parent(), n)
		// readHTMLNode already handles children and siblings, so we return.
		// TODO: for something like [TestButtonInHeadingBug] this will not
		// have the right behavior, but that is a rare use case and this
		// heuristic is much simpler.
		return str, nil
	}
	if n.FirstChild != nil {
		start, end := nodeString(n)
		tx, excl := extractTextExclude(ctx, n.FirstChild, excludeSubs...)
		str = start + tx + end
		if excl != nil {
			return str, excl
		}
	}
	if n.NextSibling != nil {
		tx, excl := extractTextExclude(ctx, n.NextSibling, excludeSubs...)
		str += tx
		if excl != nil {
			return str, excl
		}
	}
	return str, nil
}

// nodeString returns the given node as starting and ending strings in the format:
//
//	<tag attr0="value0" attr1="value1">
//
// and
//
//	</tag>
//
// It returns "", "" if the given node is not an [html.ElementNode]
func nodeString(n *html.Node) (start, end string) {
	if n.Type != html.ElementNode {
		return
	}
	tag := n.Data
	start = "<" + tag
	for _, a := range n.Attr {
		start += " " + a.Key + "=" + `"` + a.Val + `"`
	}
	start += ">"
	end = "</" + tag + ">"
	return
}

// textTags are all of the node tags that result in a true return value for [isText].
var textTags = []string{
	"a", "abbr", "b", "bdi", "bdo", "br", "cite", "code", "data", "dfn",
	"em", "i", "kbd", "mark", "q", "rp", "rt", "ruby", "s", "samp", "small",
	"span", "strong", "sub", "sup", "time", "u", "var", "wbr",
}

// isText returns true if the given node is a [html.TextNode] or
// an [html.ElementNode] designed for holding text (a, span, b, code, etc),
// and false otherwise.
func isText(n *html.Node) bool {
	if n.Type == html.TextNode {
		return true
	}
	if n.Type == html.ElementNode {
		tag := n.Data
		return slices.Contains(textTags, tag)
	}
	return false
}

// Copyright (c) 2025, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lines

import (
	"testing"

	_ "github.com/MobinYengejehi/core/system/driver"
	"github.com/stretchr/testify/assert"
)

func TestMarkup(t *testing.T) {
	src := `func (ls *Lines) deleteTextRectImpl(st, ed textpos.Pos) *textpos.Edit {
	tbe := ls.regionRect(st, ed)
	if tbe == nil {
	return nil
	}
`

	lns, vid := NewLinesFromBytes("dummy.go", 40, []byte(src))
	vw := lns.view(vid)
	assert.Equal(t, src+"\n", lns.String())

	mu0 := `[monospace bold fill-color]: "func"
[monospace]: " "
[monospace]: "("
[monospace]: "ls"
[monospace]: " "
[monospace fill-color]: "*"
[monospace]: "Lines"
[monospace]: ")"
[monospace]: " "
[monospace]: "deleteTextRectImpl"
[monospace]: "("
[monospace]: "st"
[monospace]: ","
[monospace]: " "
`
	mu1 := `[monospace]: "ed"
[monospace]: " "
[monospace]: "textpos"
[monospace]: "."
[monospace]: "Pos"
[monospace]: ")"
[monospace]: " "
[monospace fill-color]: "*"
[monospace]: "textpos"
[monospace]: "."
[monospace]: "Edit"
[monospace]: " "
[monospace]: "{"
`
	// fmt.Println(vw.markup[0])
	assert.Equal(t, mu0, vw.markup[0].String())

	// fmt.Println(vw.markup[1])
	assert.Equal(t, mu1, vw.markup[1].String())
}

func TestLineWrap(t *testing.T) {
	src := `The [rich.Text](http://rich.text.com) type is the standard representation for formatted text, used as the input to the "shaped" package for text layout and rendering. It is encoded purely using "[]rune" slices for each span, with the _style_ information **represented** with special rune values at the start of each span. This is an efficient and GPU-friendly pure-value format that avoids any issues of style struct pointer management etc.
`

	lns, vid := NewLinesFromBytes("dummy.md", 80, []byte(src))
	vw := lns.view(vid)
	assert.Equal(t, src+"\n", lns.String())

	tmu := []string{`[monospace]: "The "
[monospace fill-color]: "[rich.Text]"
[monospace fill-color]: "(http://rich.text.com)"
[monospace]: " type is the standard representation for "
`,

		`[monospace]: "formatted text, used as the input to the "
[monospace fill-color]: ""shaped""
[monospace]: " package for text layout and "
`,

		`[monospace]: "rendering. It is encoded purely using "
[monospace fill-color]: ""[]rune""
[monospace]: " slices for each span, with the"
[monospace italic]: " "
`,

		`[monospace italic]: "_style_"
[monospace]: " information"
[monospace bold]: " **represented**"
[monospace]: " with special rune values at the start of "
`,

		`[monospace]: "each span. This is an efficient and GPU-friendly pure-value format that avoids "
`,
		`[monospace]: "any issues of style struct pointer management etc."
`,
	}

	join := `The [rich.Text](http://rich.text.com) type is the standard representation for 
formatted text, used as the input to the "shaped" package for text layout and 
rendering. It is encoded purely using "[]rune" slices for each span, with the 
_style_ information **represented** with special rune values at the start of 
each span. This is an efficient and GPU-friendly pure-value format that avoids 
any issues of style struct pointer management etc.
`
	assert.Equal(t, 6, vw.viewLines)

	jtxt := ""
	for i := range vw.viewLines {
		trg := tmu[i]
		// fmt.Println(vw.markup[i])
		assert.Equal(t, trg, vw.markup[i].String())
		jtxt += string(vw.markup[i].Join()) + "\n"
	}
	// fmt.Println(jtxt)
	assert.Equal(t, join, jtxt)
}

func TestMarkupSpaces(t *testing.T) {
	src := `Name           string
`

	lns, vid := NewLinesFromBytes("dummy.go", 40, []byte(src))
	vw := lns.view(vid)
	assert.Equal(t, src+"\n", lns.String())

	mu0 := `[monospace]: "Name"
[monospace]: "           "
[monospace bold fill-color]: "string"
`
	// fmt.Println(lns.markup[0])
	// fmt.Println(vw.markup[0])
	assert.Equal(t, mu0, vw.markup[0].String())
}

func TestLongLineWrap(t *testing.T) {
	src := `The [rich.Text](http://rich.text.com) type is the standard representation for formatted text, usedastheinputtotheshapedpackagefortextlayoutandrenderingItisencodedpurelyusingruneslicesforeachspanwiththestyleinformationrepresentedwithspecial rune values at the start of each span. This is an efficient and GPU-friendly pure-value format that avoids any issues of style struct pointer management etc.
`

	lns, vid := NewLinesFromBytes("dummy.md", 80, []byte(src))
	vw := lns.view(vid)
	assert.Equal(t, src+"\n", lns.String())

	tmu := []string{`[monospace]: "The "
[monospace fill-color]: "[rich.Text]"
[monospace fill-color]: "(http://rich.text.com)"
[monospace]: " type is the standard representation for "
`,

		`[monospace]: "formatted text, "
`,

		`[monospace]: "usedastheinputtotheshapedpackagefortextlayoutandrenderingItisencodedpurelyusingr"
`,

		`[monospace]: "uneslicesforeachspanwiththestyleinformationrepresentedwithspecial "
`,

		`[monospace]: "rune values at the start of each span. This is an efficient and GPU-friendly "
`,
		`[monospace]: "pure-value format that avoids any issues of style struct pointer management etc."
`,
	}

	assert.Equal(t, 6, vw.viewLines)

	join := `The [rich.Text](http://rich.text.com) type is the standard representation for 
formatted text, 
usedastheinputtotheshapedpackagefortextlayoutandrenderingItisencodedpurelyusingr
uneslicesforeachspanwiththestyleinformationrepresentedwithspecial 
rune values at the start of each span. This is an efficient and GPU-friendly 
pure-value format that avoids any issues of style struct pointer management etc.
`

	jtxt := ""
	for i := range vw.viewLines {
		trg := tmu[i]
		// fmt.Println(vw.markup[i])
		assert.Equal(t, trg, vw.markup[i].String())
		jtxt += string(vw.markup[i].Join()) + "\n"
	}
	// fmt.Println(jtxt)
	assert.Equal(t, join, jtxt)
}

func TestLineWrapSVG(t *testing.T) {
	src := `<svg xmlns="http://www.w3.org/2000/svg" width="256" height="256" viewBox="0 0 1 1"><path d="M.833.675a.35.35 0 1 1 0-.35" style="stroke:#005bc0;stroke-width:.27;fill:none"/><circle cx=".53" cy=".5" r=".23" style="fill:#fbbd0e;stroke:none"/></svg>
`

	lns, vid := NewLinesFromBytes("dummy.svg", 44, []byte(src))
	vw := lns.view(vid)
	assert.Equal(t, src+"\n", lns.String())

	tmu := []string{`[monospace fill-color]: "<svg"
[monospace]: " "
[monospace fill-color]: "xmlns="
[monospace fill-color]: ""http://www.w3.org/2000/svg""
[monospace]: " "
`,

		`[monospace fill-color]: "width="
[monospace fill-color]: ""256""
[monospace]: " "
[monospace fill-color]: "height="
[monospace fill-color]: ""256""
[monospace]: " "
[monospace fill-color]: "viewBox="
[monospace fill-color]: ""0 0 1 "
`,

		`[monospace fill-color]: "1""
[monospace fill-color]: "><path"
[monospace]: " "
[monospace fill-color]: "d="
[monospace fill-color]: ""M.833.675a.35.35 0 1 1 0-.35""
[monospace]: " "
`,

		`[monospace fill-color]: "style="
[monospace fill-color]: ""stroke:#005bc0;stroke-width:.27;fill:"
`,

		`[monospace fill-color]: "none""
[monospace fill-color]: "/><circle"
[monospace]: " "
`,
		`[monospace fill-color]: "cx="
[monospace fill-color]: "".53""
[monospace]: " "
[monospace fill-color]: "cy="
[monospace fill-color]: "".5""
[monospace]: " "
[monospace fill-color]: "r="
[monospace fill-color]: "".23""
[monospace]: " "
`,
		`[monospace fill-color]: "style="
[monospace fill-color]: ""fill:#fbbd0e;stroke:none""
[monospace fill-color]: "/></svg>"
`,
	}
	assert.Equal(t, 7, vw.viewLines)

	join := `<svg xmlns="http://www.w3.org/2000/svg" 
width="256" height="256" viewBox="0 0 1 
1"><path d="M.833.675a.35.35 0 1 1 0-.35" 
style="stroke:#005bc0;stroke-width:.27;fill:
none"/><circle 
cx=".53" cy=".5" r=".23" 
style="fill:#fbbd0e;stroke:none"/></svg>
`

	jtxt := ""
	for i := range vw.viewLines {
		trg := tmu[i]
		// fmt.Println(vw.markup[i])
		assert.Equal(t, trg, vw.markup[i].String())
		jtxt += string(vw.markup[i].Join()) + "\n"
	}
	// fmt.Println(jtxt)
	assert.Equal(t, join, jtxt)
}

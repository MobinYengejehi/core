// Copyright (c) 2025, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package highlighting

import (
	"fmt"
	"testing"

	"github.com/MobinYengejehi/core/base/fileinfo"
	_ "github.com/MobinYengejehi/core/system/driver"
	"github.com/MobinYengejehi/core/text/parse"
	"github.com/MobinYengejehi/core/text/parse/lexer"
	"github.com/MobinYengejehi/core/text/rich"
	"github.com/MobinYengejehi/core/text/runes"
	"github.com/MobinYengejehi/core/text/token"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/stretchr/testify/assert"
)

func TestMarkup(t *testing.T) {

	src := `	if len(txt) > maxLineLen { // avoid overflow`
	rsrc := []rune(src)

	fi, err := fileinfo.NewFileInfo("dummy.go")
	assert.Error(t, err)

	var pst parse.FileStates
	pst.SetSrc("dummy.go", "", fi.Known)

	hi := Highlighter{}
	hi.Init(fi, &pst)
	hi.SetStyle(HighlightingName("emacs"))

	fs := pst.Done() // initialize
	fs.Src.SetBytes([]byte(src))

	lex, err := hi.MarkupTagsLine(0, rsrc)
	assert.NoError(t, err)

	hitrg := `[{NameFunction: if 1 3 {0 0}} {NameBuiltin 4 7 {0 0}} {PunctGpLParen 7 8 {0 0}} {+1:Name 8 11 {0 0}} {PunctGpRParen 11 12 {0 0}} {OpRelGreater 13 14 {0 0}} {Name 15 25 {0 0}} {PunctGpLBrace 26 27 {0 0}} {+1:EOS 27 27 {0 0}} {+1:Comment 28 45 {0 0}}]`
	assert.Equal(t, hitrg, fmt.Sprint(lex))
	// fmt.Println(lex)

	// this "avoid" is what drives the need for depth in styles
	// we're marking it as misspelled
	aix := runes.Index(rsrc, []rune("avoid"))
	ot := []lexer.Lex{lexer.Lex{Token: token.KeyToken{Token: token.TextSpellErr, Depth: 1}, Start: aix, End: aix + 5}}

	// todo: it doesn't detect the offset of the embedded avoid token here!

	sty := rich.NewStyle()
	sty.Family = rich.Monospace
	tx := MarkupLineRich(hi.Style, sty, rsrc, lex, ot)

	rtx := `[monospace]: "	"
[monospace fill-color]: "if"
[monospace]: " "
[monospace fill-color]: "len"
[monospace]: "("
[monospace]: "txt"
[monospace]: ")"
[monospace]: " "
[monospace fill-color]: ">"
[monospace]: " "
[monospace]: "maxLineLen"
[monospace]: " "
[monospace]: "{"
[monospace]: " "
[monospace italic fill-color]: "// "
[monospace italic dotted-underline fill-color]: "avoid"
[monospace italic fill-color]: " overflow"
`
	// fmt.Println(tx)
	assert.Equal(t, rtx, fmt.Sprint(tx))

	for i, r := range rsrc {
		si, sn, ri := tx.Index(i)
		if tx[si][ri] != r {
			fmt.Println(i, string(r), string(tx[si][ri]), si, ri, sn)
		}
		assert.Equal(t, string(r), string(tx[si][ri]))
	}

	rht := `	<span class="nf">if</span> <span class="nb">len</span><span class="">(</span><span class="n">txt</span><span class="">)</span> <span class="">></span> <span class="n">maxLineLen</span> <span class="">{</span><span class="EOS"></span> <span class="c">// <span class="te">avoid</span> overflow</span>`

	b := MarkupLineHTML(rsrc, lex, ot, NoEscapeHTML)
	assert.Equal(t, rht, fmt.Sprint(string(b)))

}

func TestMarkupSpaces(t *testing.T) {

	src := `Name        string`
	rsrc := []rune(src)

	fi, err := fileinfo.NewFileInfo("dummy.go")
	assert.Error(t, err)

	var pst parse.FileStates
	pst.SetSrc("dummy.go", "", fi.Known)

	hi := Highlighter{}
	hi.Init(fi, &pst)
	hi.SetStyle(HighlightingName("emacs"))

	fs := pst.Done() // initialize
	fs.Src.SetBytes([]byte(src))

	lex, err := hi.MarkupTagsLine(0, rsrc)
	assert.NoError(t, err)

	hitrg := `[{Name 0 4 {0 0}} {KeywordType: string 12 18 {0 0}} {EOS 18 18 {0 0}}]`
	assert.Equal(t, hitrg, fmt.Sprint(lex))
	// fmt.Println(lex)

	sty := rich.NewStyle()
	sty.Family = rich.Monospace
	tx := MarkupLineRich(hi.Style, sty, rsrc, lex, nil)

	rtx := `[monospace]: "Name"
[monospace]: "        "
[monospace bold fill-color]: "string"
`
	// fmt.Println(tx)
	assert.Equal(t, rtx, fmt.Sprint(tx))

	for i, r := range rsrc {
		si, sn, ri := tx.Index(i)
		if tx[si][ri] != r {
			fmt.Println(i, string(r), string(tx[si][ri]), si, ri, sn)
		}
		assert.Equal(t, string(r), string(tx[si][ri]))
	}
}

func TestMarkupPathsAsLinks(t *testing.T) {
	flds := []string{
		"./path/file.go",
		"/absolute/path/file.go",
		"../relative/path/file.go",
		"file.go",
		"./commands.go:68:6: ps redeclared in this block",
	}

	res := []string{
		`[link [file:///./path/file.go] underline fill-color]: "./path/file.go"
[{End Special}]: ""
`,
		`[link [file:////absolute/path/file.go] underline fill-color]: "/absolute/path/file.go"
[{End Special}]: ""
`,
		`[link [file:///../relative/path/file.go] underline fill-color]: "../relative/path/file.go"
[{End Special}]: ""
`,
		`[]: "file.go"
`,
		`[link [file:///./commands.go#L68C6] underline fill-color]: "./commands.go:68:6:"
[{End Special}]: ""
[]: " ps redeclared in this block"
`,
	}

	for i, fld := range flds {
		rfd := []rune(fld)
		mu := rich.NewPlainText(rfd)
		nmu := MarkupPathsAsLinks(rfd, mu, 2)
		// fmt.Println(i, nmu) // todo: make it a test
		assert.Equal(t, res[i], nmu.String())
	}
}

func TestMarkupDiff(t *testing.T) {
	src := `diff --git a/code/cdebug/cdelve/cdelve.go b/code/cdebug/cdelve/cdelve.goindex 83ee192..6d2e820 100644"`
	rsrc := []rune(src)

	hi := Highlighter{}
	hi.SetStyle(HighlightingName("emacs"))

	clex := lexers.Get("diff")
	ctags, _ := ChromaTagsLine(clex, src)

	// hitrg := `[{Name 0 4 {0 0}} {KeywordType: string 12 18 {0 0}} {EOS 18 18 {0 0}}]`
	// assert.Equal(t, hitrg, fmt.Sprint(lex))
	// fmt.Println(ctags)

	sty := rich.NewStyle()
	sty.Family = rich.Monospace
	tx := MarkupLineRich(hi.Style, sty, rsrc, ctags, nil)

	rtx := `[monospace bold fill-color]: "diff --git a/code/cdebug/cdelve/cdelve.go b/code/cdebug/cdelve/cdelve.goindex 83ee192..6d2e820 100644""
`
	// _ = rtx
	// fmt.Println(tx)
	assert.Equal(t, rtx, fmt.Sprint(tx))

	for i, r := range rsrc {
		si, sn, ri := tx.Index(i)
		if tx[si][ri] != r {
			fmt.Println(i, string(r), string(tx[si][ri]), si, ri, sn)
		}
		assert.Equal(t, string(r), string(tx[si][ri]))
	}
}

func TestMarkupSpellErr(t *testing.T) {
	src := `this is a mispel word [test](link) and anotherr bad word`
	rsrc := []rune(src)

	hi := Highlighter{}
	hi.SetStyle(HighlightingName("emacs"))

	clex := lexers.Get("markdown")
	ctags, _ := ChromaTagsLine(clex, src)
	etags := lexer.Line{}
	etags.AddLex(token.KeyToken{Token: token.TextSpellErr}, 10, 16)
	etags.AddLex(token.KeyToken{Token: token.TextSpellErr}, 39, 47)

	sty := rich.NewStyle()
	sty.Family = rich.Monospace
	tx := MarkupLineRich(hi.Style, sty, rsrc, ctags, etags)

	rtx := `[monospace]: "this is a "
[monospace dotted-underline]: "mispel"
[monospace]: " word ["
[monospace fill-color]: "test"
[monospace]: "]("
[monospace fill-color]: "link"
[monospace]: ") and "
[monospace dotted-underline]: "anotherr"
[monospace]: " bad word"
`
	_ = rtx
	// fmt.Println(tx)
	assert.Equal(t, rtx, fmt.Sprint(tx))

	for i, r := range rsrc {
		si, sn, ri := tx.Index(i)
		if tx[si][ri] != r {
			fmt.Println(i, string(r), string(tx[si][ri]), si, ri, sn)
		}
		assert.Equal(t, string(r), string(tx[si][ri]))
	}
}

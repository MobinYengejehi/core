// Copyright (c) 2020, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lines

import (
	"github.com/MobinYengejehi/core/base/fileinfo"
	"github.com/MobinYengejehi/core/base/indent"
	"github.com/MobinYengejehi/core/text/parse"
	"github.com/MobinYengejehi/core/text/text"
)

// Settings contains settings for editing text lines.
type Settings struct {
	text.EditorSettings

	// CommentLine are character(s) that start a single-line comment;
	// if empty then multi-line comment syntax will be used.
	CommentLine string

	// CommentStart are character(s) that start a multi-line comment
	// or one that requires both start and end.
	CommentStart string

	// Commentend are character(s) that end a multi-line comment
	// or one that requires both start and end.
	CommentEnd string
}

// CommentStrings returns the comment start and end strings,
// using line-based CommentLn first if set and falling back
// on multi-line / general purpose start / end syntax.
func (tb *Settings) CommentStrings() (comst, comed string) {
	comst = tb.CommentLine
	if comst == "" {
		comst = tb.CommentStart
		comed = tb.CommentEnd
	}
	return
}

// IndentChar returns the indent character based on SpaceIndent option
func (tb *Settings) IndentChar() indent.Character {
	if tb.SpaceIndent {
		return indent.Space
	}
	return indent.Tab
}

// ConfigKnown configures options based on the supported language info in parse.
// Returns true if supported.
func (tb *Settings) ConfigKnown(sup fileinfo.Known) bool {
	if sup == fileinfo.Unknown {
		return false
	}
	lp, ok := parse.StandardLanguageProperties[sup]
	if !ok {
		return false
	}
	tb.CommentLine = lp.CommentLn
	tb.CommentStart = lp.CommentSt
	tb.CommentEnd = lp.CommentEd
	for _, flg := range lp.Flags {
		switch flg {
		case parse.IndentSpace:
			tb.SpaceIndent = true
		case parse.IndentTab:
			tb.SpaceIndent = false
		}
	}
	return true
}

// KnownComments returns the comment strings for supported file types,
// and returns the standard C-style comments otherwise.
func KnownComments(fpath string) (comLn, comSt, comEd string) {
	comLn = "//"
	comSt = "/*"
	comEd = "*/"
	mtyp, _, err := fileinfo.MimeFromFile(fpath)
	if err != nil {
		return
	}
	sup := fileinfo.MimeKnown(mtyp)
	if sup == fileinfo.Unknown {
		return
	}
	lp, ok := parse.StandardLanguageProperties[sup]
	if !ok {
		return
	}
	comLn = lp.CommentLn
	comSt = lp.CommentSt
	comEd = lp.CommentEd
	return
}

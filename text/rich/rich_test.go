// Copyright (c) 2025, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rich

import (
	"image/color"
	"testing"

	"github.com/MobinYengejehi/core/colors"
	"github.com/MobinYengejehi/core/text/runes"
	"github.com/MobinYengejehi/core/text/textpos"
	"github.com/stretchr/testify/assert"
)

func TestColors(t *testing.T) {
	c := color.RGBA{22, 55, 77, 255}
	r := ColorToRune(c)
	rc := ColorFromRune(r)
	assert.Equal(t, c, rc)
}

func TestStyle(t *testing.T) {
	s := NewStyle()
	s.Family = Math
	s.Special = MathInline
	s.SetBackground(colors.Blue)

	sr := RuneFromSpecial(s.Special)
	ss := RuneToSpecial(sr)
	assert.Equal(t, s.Special, ss)

	rs := s.ToRunes()

	assert.Equal(t, 3, len(rs))
	assert.Equal(t, 1, s.Decoration.NumColors())

	ns := &Style{}
	ns.FromRunes(rs)

	assert.Equal(t, s, ns)
}

func TestText(t *testing.T) {
	src := "The lazy fox typed in some familiar text"
	sr := []rune(src)
	tx := Text{}
	plain := NewStyle() // .SetFamily(Monospace)
	ital := plain.Clone().SetSlant(Italic)
	ital.SetStrokeColor(colors.Red)
	// ital.SetFillColor(colors.Red)
	boldBig := plain.Clone().SetWeight(Bold).SetSize(1.5)
	tx.AddSpan(plain, sr[:4])
	tx.AddSpan(ital, sr[4:8])
	fam := []rune("familiar")
	ix := runes.Index(sr, fam)
	tx.AddSpan(plain, sr[8:ix])
	tx.AddSpan(boldBig, sr[ix:ix+8])
	tx.AddSpan(plain, sr[ix+8:])

	str := tx.String()
	trg := `[]: "The "
[italic stroke-color]: "lazy"
[]: " fox typed in some "
[1.50x bold]: "familiar"
[]: " text"
`
	assert.Equal(t, trg, str)

	os := tx.Join()
	assert.Equal(t, src, string(os))

	for i := range src {
		assert.Equal(t, rune(src[i]), tx.At(i))
	}

	ssi := tx.SplitSpan(12)
	trg = `[]: "The "
[italic stroke-color]: "lazy"
[]: " fox"
[]: " typed in some "
[1.50x bold]: "familiar"
[]: " text"
`
	// fmt.Println(tx)
	assert.Equal(t, 3, ssi)
	assert.Equal(t, trg, tx.String())

	idxTests := []struct {
		idx int
		si  int
		sn  int
		ri  int
	}{
		{0, 0, 2, 2},
		{2, 0, 2, 4},
		{4, 1, 3, 3},
		{7, 1, 3, 6},
		{8, 2, 2, 2},
		{9, 2, 2, 3},
		{11, 2, 2, 5},
		{16, 3, 2, 6},
	}
	for _, test := range idxTests {
		si, sn, ri := tx.Index(test.idx)
		stx := string(tx[si][ri:])
		trg := string(sr[test.idx : test.idx+3])
		// fmt.Printf("%d\tsi:%d\tsn:%d\tri:%d\tsisrc: %q txt: %q\n", test.idx, si, sn, ri, stx, trg)
		assert.Equal(t, test.si, si)
		assert.Equal(t, test.sn, sn)
		assert.Equal(t, test.ri, ri)
		assert.Equal(t, trg[0], stx[0])
	}

	// spl := tx.Split()
	// for i := range spl {
	// 	fmt.Println(string(spl[i]))
	// }

	tx.SetSpanStyle(3, ital)
	trg = `[]: "The "
[italic stroke-color]: "lazy"
[]: " fox"
[italic stroke-color]: " typed in some "
[1.50x bold]: "familiar"
[]: " text"
`
	// fmt.Println(tx)
	assert.Equal(t, trg, tx.String())
}

func TestLink(t *testing.T) {
	src := "Pre link link text post link"
	tx := Text{}
	plain := NewStyle()
	ital := NewStyle().SetSlant(Italic)
	ital.SetStrokeColor(colors.Red)
	boldBig := NewStyle().SetWeight(Bold).SetSize(1.5)
	tx.AddSpanString(plain, "Pre link ")
	tx.AddLink(ital, "https://example.com", "link text")
	tx.AddSpanString(boldBig, " post link")

	str := tx.String()
	trg := `[]: "Pre link "
[italic link [https://example.com] stroke-color]: "link text"
[{End Special}]: ""
[1.50x bold]: " post link"
`
	assert.Equal(t, trg, str)

	os := tx.Join()
	assert.Equal(t, src, string(os))

	for i := range src {
		assert.Equal(t, rune(src[i]), tx.At(i))
	}

	lks := tx.GetLinks()
	assert.Equal(t, 1, len(lks))
	assert.Equal(t, textpos.Range{9, 18}, lks[0].Range)
	assert.Equal(t, "link text", lks[0].Label)
	assert.Equal(t, "https://example.com", lks[0].URL)
}

func TestSplitSpaces(t *testing.T) {
	src := "Pre link text post link "
	tx := NewPlainText([]rune(src))
	tx.SplitSpaces()
	trg := `[]: "Pre "
[]: "link "
[]: "text "
[]: "post "
[]: "link "
`
	// fmt.Println(tx)
	assert.Equal(t, trg, tx.String())
}

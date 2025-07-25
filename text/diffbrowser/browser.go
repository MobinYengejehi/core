// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbrowser

//go:generate core generate

import (
	"github.com/MobinYengejehi/core/base/fsx"
	"github.com/MobinYengejehi/core/base/stringsx"
	"github.com/MobinYengejehi/core/core"
	"github.com/MobinYengejehi/core/events"
	"github.com/MobinYengejehi/core/styles"
	"github.com/MobinYengejehi/core/text/textcore"
	"github.com/MobinYengejehi/core/tree"
)

// Browser is a diff browser, for browsing a set of paired files
// for viewing differences between them, organized into a tree
// structure, e.g., reflecting their source in a filesystem.
type Browser struct {
	core.Frame

	// starting paths for the files being compared
	PathA, PathB string
}

func (br *Browser) Init() {
	br.Frame.Init()
	br.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
	})
	br.OnShow(func(e events.Event) {
		br.OpenFiles()
	})

	tree.AddChildAt(br, "splits", func(w *core.Splits) {
		w.SetSplits(.15, .85)
		tree.AddChildAt(w, "treeframe", func(w *core.Frame) {
			w.Styler(func(s *styles.Style) {
				s.Direction = styles.Column
				s.Overflow.Set(styles.OverflowAuto)
				s.Grow.Set(1, 1)
			})
			tree.AddChildAt(w, "tree", func(w *Node) {})
		})
		tree.AddChildAt(w, "tabs", func(w *core.Tabs) {
			w.Type = core.FunctionalTabs
		})
	})
}

// NewBrowserWindow opens a new diff Browser in a new window
func NewBrowserWindow() (*Browser, *core.Body) {
	b := core.NewBody("Diff browser")
	br := NewBrowser(b)
	br.UpdateTree() // must have tree
	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(br.MakeToolbar)
	})
	return br, b
}

func (br *Browser) Splits() *core.Splits {
	return br.FindPath("splits").(*core.Splits)
}

func (br *Browser) Tree() *Node {
	sp := br.Splits()
	return sp.Child(0).AsTree().Child(0).(*Node)
}

func (br *Browser) Tabs() *core.Tabs {
	return br.FindPath("splits/tabs").(*core.Tabs)
}

// OpenFiles Updates the tree based on files
func (br *Browser) OpenFiles() { //types:add
	tv := br.Tree()
	if tv == nil {
		return
	}
	tv.Open()
}

func (br *Browser) MakeToolbar(p *tree.Plan) {
	// tree.Add(p, func(w *core.FuncButton) {
	// 	w.SetFunc(br.OpenFiles).SetText("").SetIcon(icons.Refresh).SetShortcut("Command+U")
	// })
}

// ViewDiff views diff for given file Node, returning a textcore.DiffEditor
func (br *Browser) ViewDiff(fn *Node) *textcore.DiffEditor {
	df := fsx.DirAndFile(fn.FileA)
	tabs := br.Tabs()
	tab := tabs.RecycleTab(df)
	if tab.HasChildren() {
		dv := tab.Child(1).(*textcore.DiffEditor)
		return dv
	}
	tb := core.NewToolbar(tab)
	de := textcore.NewDiffEditor(tab)
	tb.Maker(de.MakeToolbar)
	de.SetFileA(fn.FileA).SetFileB(fn.FileB).SetRevisionA(fn.RevA).SetRevisionB(fn.RevB)
	de.DiffStrings(stringsx.SplitLines(fn.TextA), stringsx.SplitLines(fn.TextB))
	br.Update()
	return de
}

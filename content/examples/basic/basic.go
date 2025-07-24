// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"

	"github.com/MobinYengejehi/core/content"
	"github.com/MobinYengejehi/core/core"
	"github.com/MobinYengejehi/core/htmlcore"
	_ "github.com/MobinYengejehi/core/yaegicore"
)

//go:embed content
var econtent embed.FS

func main() {
	b := core.NewBody("Cogent Content Example")
	ct := content.NewContent(b).SetContent(econtent)
	ct.Context.AddWikilinkHandler(htmlcore.GoDocWikilink("doc", "github.com/MobinYengejehi/core"))
	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(ct.MakeToolbar)
	})
	b.RunMainWindow()
}

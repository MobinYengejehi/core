// Copyright (c) 2019, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js

package xyz

import (
	"io/fs"

	"github.com/MobinYengejehi/core/text/fonts/noto"
	"github.com/MobinYengejehi/core/text/fonts/robotomono"
	"github.com/MobinYengejehi/core/text/shaped/shapers/shapedgt"
)

func initTextShaper(sc *Scene) {
	sc.TextShaper = shapedgt.NewShaperWithFonts([]fs.FS{noto.Embedded, robotomono.Embedded})
}

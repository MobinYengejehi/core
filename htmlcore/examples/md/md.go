// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	_ "embed"

	"github.com/MobinYengejehi/core/base/errors"
	"github.com/MobinYengejehi/core/core"
	"github.com/MobinYengejehi/core/htmlcore"
	_ "github.com/MobinYengejehi/core/text/tex" // include this to get math
)

//go:embed example.md
var content string

func main() {
	b := core.NewBody("MD Example")
	errors.Log(htmlcore.ReadMDString(htmlcore.NewContext(), b, content))
	b.RunMainWindow()
}

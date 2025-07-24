// Copyright (c) 2025, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js

package shapers

import (
	"github.com/MobinYengejehi/core/text/shaped"
	"github.com/MobinYengejehi/core/text/shaped/shapers/shapedjs"
)

func init() {
	shaped.NewShaper = shapedjs.NewShaper
}

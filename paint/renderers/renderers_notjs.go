// Copyright (c) 2025, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js

package renderers

import (
	"github.com/MobinYengejehi/core/paint"
	"github.com/MobinYengejehi/core/paint/renderers/rasterx"
	_ "github.com/MobinYengejehi/core/text/shaped/shapers"
)

func init() {
	paint.NewSourceRenderer = rasterx.New
	paint.NewImageRenderer = rasterx.New
}

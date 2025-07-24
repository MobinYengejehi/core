// Copyright 2023 Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package offscreen

import (
	"github.com/MobinYengejehi/core/system"
	"github.com/MobinYengejehi/core/system/composer"
	"github.com/MobinYengejehi/core/system/driver/base"
)

// Window is the implementation of [system.Window] for the offscreen platform.
type Window struct {
	base.WindowMulti[*App, *composer.ComposerDrawer]
}

func (w *Window) Screen() *system.Screen {
	return TheApp.Screen(0)
}

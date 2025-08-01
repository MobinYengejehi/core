// Copyright 2023 Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build android && !offscreen

package driver

import (
	"github.com/MobinYengejehi/core/system/driver/android"
)

func init() {
	android.Init()
}

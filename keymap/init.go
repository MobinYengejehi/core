// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package keymap

import (
	"github.com/MobinYengejehi/core/system"

	// we have to import system/driver here so that it is initialized
	// in time for us to the get the system platform
	_ "github.com/MobinYengejehi/core/system/driver"
)

func init() {
	AvailableMaps.CopyFrom(StandardMaps)
	switch system.TheApp.SystemPlatform() {
	case system.MacOS:
		DefaultMap = "MacStandard"
	case system.Windows:
		DefaultMap = "WindowsStandard"
	default:
		DefaultMap = "LinuxStandard"
	}
	SetActiveMapName(DefaultMap)
}

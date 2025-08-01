// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"testing"

	"github.com/MobinYengejehi/core/styles"
)

func TestSceneInit(t *testing.T) {
	TheApp.SetSceneInit(func(sc *Scene) {
		sc.SetWidgetInit(func(w Widget) {
			switch w := w.(type) {
			case *Button:
				w.Styler(func(s *styles.Style) {
					s.Border.Radius = styles.BorderRadiusSmall
				})
			}
		})
	})
	defer func() {
		TheApp.SetSceneInit(nil)
	}()
	b := NewBody()
	NewButton(b).SetText("Test")
	b.AssertRender(t, "app/scene-init")
}

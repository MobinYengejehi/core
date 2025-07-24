// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"testing"

	"github.com/MobinYengejehi/core/colors"
	"github.com/MobinYengejehi/core/math32"
	"github.com/MobinYengejehi/core/paint"
)

func TestCanvas(t *testing.T) {
	b := NewBody()
	NewCanvas(b).SetDraw(func(pc *paint.Painter) {
		pc.Stroke.Color = colors.Uniform(colors.Blue)
		pc.MoveTo(0.15, 0.3)
		pc.LineTo(0.3, 0.15)
		pc.Draw()
		pc.Stroke.Color = nil

		pc.FillBox(math32.Vec2(0.7, 0.3), math32.Vec2(0.2, 0.5), colors.Scheme.Success.Container)
		pc.Fill.Color = colors.Uniform(colors.Orange)
		pc.Circle(0.4, 0.5, 0.15)
		pc.Draw()
	})
	b.AssertRender(t, "canvas/basic")
}

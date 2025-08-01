// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/MobinYengejehi/core/colors"
	"github.com/MobinYengejehi/core/cursors"
	"github.com/MobinYengejehi/core/events"
	"github.com/MobinYengejehi/core/math32"
	"github.com/MobinYengejehi/core/styles"
	"github.com/MobinYengejehi/core/styles/abilities"
	"github.com/MobinYengejehi/core/styles/units"
)

// Handle represents a draggable handle that can be used to
// control the size of an element. The [styles.Style.Direction]
// controls the direction in which the handle moves.
type Handle struct {
	WidgetBase

	// Min is the minimum value that the handle can go to
	// (typically the lower bound of the dialog/splits)
	Min float32

	// Max is the maximum value that the handle can go to
	// (typically the upper bound of the dialog/splits)
	Max float32

	// Pos is the current position of the handle on the
	// scale of [Handle.Min] to [Handle.Max].
	Pos float32
}

func (hl *Handle) Init() {
	hl.WidgetBase.Init()
	hl.Styler(func(s *styles.Style) {
		s.SetAbilities(true, abilities.Clickable, abilities.Focusable, abilities.Hoverable, abilities.Slideable, abilities.ScrollableUnattended)

		s.Border.Radius = styles.BorderRadiusFull
		s.Background = colors.Scheme.OutlineVariant
	})
	hl.FinalStyler(func(s *styles.Style) {
		if s.Direction == styles.Row {
			s.Min.X.Dp(6)
			s.Min.Y.Em(2)
			s.Margin.SetHorizontal(units.Dp(6))
		} else {
			s.Min.X.Em(2)
			s.Min.Y.Dp(6)
			s.Margin.SetVertical(units.Dp(6))
		}

		if !hl.IsReadOnly() {
			if s.Direction == styles.Row {
				s.Cursor = cursors.ResizeEW
			} else {
				s.Cursor = cursors.ResizeNS
			}
		}
	})

	hl.On(events.SlideMove, func(e events.Event) {
		e.SetHandled()
		pos := hl.parentWidget().PointToRelPos(e.Pos())
		hl.Pos = math32.FromPoint(pos).Dim(hl.Styles.Direction.Dim())
		hl.SendChange(e)
	})
}

// Value returns the value on a normalized scale of 0-1,
// based on [Handle.Pos], [Handle.Min], and [Handle.Max].
func (hl *Handle) Value() float32 {
	return (hl.Pos - hl.Min) / (hl.Max - hl.Min)
}

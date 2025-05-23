// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	"cogentcore.org/core/base/slicesx"
	"cogentcore.org/core/math32"
)

// Group groups together SVG elements.
// Provides a common transform for all group elements
// and shared style properties.
type Group struct {
	NodeBase
}

func (g *Group) SVGName() string { return "g" }

func (g *Group) EnforceSVGName() bool { return false }

func (g *Group) BBoxes(sv *SVG, parTransform math32.Matrix2) {
	g.BBoxesFromChildren(sv, parTransform)
}

func (g *Group) IsVisible(sv *SVG) bool {
	if g == nil || g.This == nil || !g.Paint.Display { // does not check g.Paint.Off!
		return false
	}
	nvis := g.VisBBox == math32.Box2{}
	if nvis && !g.isDef {
		return false
	}
	return true
}

func (g *Group) Render(sv *SVG) {
	if !g.IsVisible(sv) {
		return
	}
	pc := g.Painter(sv)
	pc.PushContext(&g.Paint, nil)
	g.RenderChildren(sv)
	pc.PopContext()
}

// ApplyTransform applies the given 2D transform to the geometry of this node
// each node must define this for itself
func (g *Group) ApplyTransform(sv *SVG, xf math32.Matrix2) {
	g.Paint.Transform.SetMul(xf)
	g.SetProperty("transform", g.Paint.Transform.String())
}

// ApplyDeltaTransform applies the given 2D delta transforms to the geometry of this node
// relative to given point.  Trans translation and point are in top-level coordinates,
// so must be transformed into local coords first.
// Point is upper left corner of selection box that anchors the translation and scaling,
// and for rotation it is the center point around which to rotate
func (g *Group) ApplyDeltaTransform(sv *SVG, trans math32.Vector2, scale math32.Vector2, rot float32, pt math32.Vector2) {
	xf, lpt := g.DeltaTransform(trans, scale, rot, pt, false) // group does NOT include self
	g.Paint.Transform.SetMulCenter(xf, lpt)
	g.SetProperty("transform", g.Paint.Transform.String())
}

// WriteGeom writes the geometry of the node to a slice of floating point numbers
// the length and ordering of which is specific to each node type.
// Slice must be passed and will be resized if not the correct length.
func (g *Group) WriteGeom(sv *SVG, dat *[]float32) {
	*dat = slicesx.SetLength(*dat, 6)
	g.WriteTransform(*dat, 0)
}

// ReadGeom reads the geometry of the node from a slice of floating point numbers
// the length and ordering of which is specific to each node type.
func (g *Group) ReadGeom(sv *SVG, dat []float32) {
	g.ReadTransform(dat, 0)
}

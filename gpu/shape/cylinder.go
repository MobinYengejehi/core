// Copyright 2022 Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shape

import (
	"github.com/MobinYengejehi/core/math32"
)

// Cylinder is a generalized cylinder shape, including a cone
// or truncated cone by having different size circles at either end.
// Height is up along the Y axis.
type Cylinder struct { //types:add -setters
	ShapeBase

	// height of the cylinder
	Height float32

	// radius of the top -- set to 0 for a cone
	TopRad float32

	// radius of the bottom
	BotRad float32

	// number of radial segments (32 is a reasonable default for full circle)
	RadialSegs int `min:"1"`

	// number of height segments
	HeightSegs int

	// render the top disc
	Top bool

	// render the bottom disc
	Bottom bool

	// starting angle in degrees, relative to -1,0,0 left side starting point
	AngStart float32 `min:"0" max:"360" step:"5"`

	// total angle to generate in degrees (max 360)
	AngLen float32 `min:"0" max:"360" step:"5"`
}

// NewCylinder returns a Cylinder shape with given radius, height,
// number of radial segments, number of height segments,
// and presence of a top and/or bottom cap.
// Height is along the Y axis.
func NewCylinder(height, radius float32, radialSegs, heightSegs int, top, bottom bool) *Cylinder {
	cy := &Cylinder{}
	cy.Defaults()

	cy.Height = height
	cy.TopRad = radius
	cy.BotRad = radius
	cy.RadialSegs = radialSegs
	cy.HeightSegs = heightSegs
	cy.Top = top
	cy.Bottom = bottom
	return cy
}

// NewCone returns a cone shape with the specified base radius, height,
// number of radial segments, number of height segments, and presence of a bottom cap.
// Height is along the Y axis.
func NewCone(height, radius float32, radialSegs, heightSegs int, bottom bool) *Cylinder {
	cy := &Cylinder{}
	cy.Defaults()

	cy.Height = height
	cy.TopRad = 0
	cy.BotRad = radius
	cy.RadialSegs = radialSegs
	cy.HeightSegs = heightSegs
	cy.Top = false
	cy.Bottom = bottom
	return cy
}

func (cy *Cylinder) Defaults() {
	cy.Height = 1
	cy.TopRad = 0.5
	cy.BotRad = 0.5
	cy.RadialSegs = 32
	cy.HeightSegs = 32
	cy.Top = true
	cy.Bottom = true
	cy.AngStart = 0
	cy.AngLen = 360
}

func (cy *Cylinder) MeshSize() (numVertex, nIndex int, hasColor bool) {
	numVertex, nIndex = CylinderSectorN(cy.RadialSegs, cy.HeightSegs, cy.Top, cy.Bottom)
	return
}

// SetCylinderSector sets points in given allocated arrays
func (cy *Cylinder) Set(vertex, normal, texcoord, clrs math32.ArrayF32, index math32.ArrayU32) {
	cy.CBBox = SetCylinderSector(vertex, normal, texcoord, index, cy.VertexOffset, cy.IndexOffset, cy.Height, cy.TopRad, cy.BotRad, cy.RadialSegs, cy.HeightSegs, cy.AngStart, cy.AngLen, cy.Top, cy.Bottom, cy.Pos)
}

////////////////////////////////////////////////////////////////

// CylinderSectorN returns the N's for the cylinder (truncated cone) sector
// vertex and index data with given parameters
func CylinderSectorN(radialSegs, heightSegs int, top, bottom bool) (numVertex, nIndex int) {
	numVertex = (heightSegs + 1) * (radialSegs + 1)
	nIndex = radialSegs * heightSegs * 6
	if top {
		numVertex += radialSegs*2 + 2
		nIndex += radialSegs * 3
	}
	if bottom {
		numVertex += radialSegs*2 + 2
		nIndex += radialSegs * 3
	}
	return
}

// SetCone creates a cone mesh with the specified base radius, height,
// vertex, normal, tex, index data at given starting *vertex* index
// (i.e., multiply this *3 to get actual float offset in Vtx array),
// number of radial segments, number of height segments, and presence of a bottom cap.
// Height is along the Y axis.
// pos is an arbitrary offset (for composing shapes).
func SetCone(vertex, normal, texcoord math32.ArrayF32, index math32.ArrayU32, vtxOff, idxOff int, height, radius float32, radialSegs, heightSegs int, bottom bool, pos math32.Vector3) math32.Box3 {
	return SetCylinderSector(vertex, normal, texcoord, index, vtxOff, idxOff, height, 0, radius, radialSegs, heightSegs, 0, 360, false, bottom, pos)
}

// SetCylinderSector creates a generalized cylinder (truncated cone) sector
// vertex, normal, tex, index data at given starting *vertex* index
// (i.e., multiply this *3 to get actual float offset in Vtx array),
// with the specified top and bottom radii, height, number of radial segments,
// number of height segments, sector start angle in degrees (start = -1,0,0)
// sector size angle in degrees, and presence of a top and/or bottom cap.
// Height is along the Y axis.
// pos is an arbitrary offset (for composing shapes).
func SetCylinderSector(vertex, normal, texcoord math32.ArrayF32, index math32.ArrayU32, vtxOff, idxOff int, height, topRad, botRad float32, radialSegs, heightSegs int, angStart, angLen float32, top, bottom bool, pos math32.Vector3) math32.Box3 {
	hHt := height / 2
	vtxs := [][]int{}
	uvsOrig := [][]math32.Vector2{}

	angStRad := math32.DegToRad(angStart)
	angLenRad := math32.DegToRad(angLen)

	bb := math32.Box3{}
	bb.SetEmpty()

	idx := 0
	vidx := vtxOff * 3
	tidx := vtxOff * 2

	var pt math32.Vector3
	for y := 0; y <= heightSegs; y++ {
		var vtxsRow = []int{}
		var uvsRow = []math32.Vector2{}
		v := float32(y) / float32(heightSegs)
		radius := v*(botRad-topRad) + topRad
		for x := 0; x <= radialSegs; x++ {
			u := float32(x) / float32(radialSegs)
			pt.X = -radius * math32.Cos(u*angLenRad+angStRad)
			pt.Y = -v*height + hHt
			pt.Z = radius * math32.Sin(u*angLenRad+angStRad)
			pt.SetAdd(pos)
			vertex.SetVector3(vidx+idx*3, pt)
			bb.ExpandByPoint(pt)
			vtxsRow = append(vtxsRow, idx)
			uvsRow = append(uvsRow, math32.Vec2(u, 1.0-v))
			idx++
		}
		vtxs = append(vtxs, vtxsRow)
		uvsOrig = append(uvsOrig, uvsRow)
	}

	tanTheta := (botRad - topRad) / height
	var na, nb math32.Vector3

	vOff := uint32(vtxOff)
	ii := idxOff
	for x := 0; x < radialSegs; x++ {
		if topRad != 0 {
			vertex.GetVector3(3*vtxs[0][x], &na)
			vertex.GetVector3(3*vtxs[0][x+1], &nb)
		} else {
			vertex.GetVector3(3*vtxs[1][x], &na)
			vertex.GetVector3(3*vtxs[1][x+1], &nb)
		}

		na.Y = math32.Sqrt(na.X*na.X+na.Z*na.Z) * tanTheta
		na.SetNormal()
		nb.Y = math32.Sqrt(nb.X*nb.X+nb.Z*nb.Z) * tanTheta
		nb.SetNormal()

		for y := 0; y < heightSegs; y++ {
			v1 := vtxs[y][x]
			v2 := vtxs[y+1][x]
			v3 := vtxs[y+1][x+1]
			v4 := vtxs[y][x+1]

			n1 := na
			n2 := na
			n3 := nb
			n4 := nb

			uv1 := uvsOrig[y][x]
			uv2 := uvsOrig[y+1][x]
			uv3 := uvsOrig[y+1][x+1]
			uv4 := uvsOrig[y][x+1]

			index.Set(ii, vOff+uint32(v1), vOff+uint32(v2), vOff+uint32(v4))
			ii += 3
			normal.SetVector3(3*v1, n1)
			normal.SetVector3(3*v2, n2)
			normal.SetVector3(3*v4, n4)

			index.Set(ii, vOff+uint32(v2), vOff+uint32(v3), vOff+uint32(v4))
			ii += 3
			normal.SetVector3(3*v2, n2)
			normal.SetVector3(3*v3, n3)
			normal.SetVector3(3*v4, n4)

			texcoord.SetVector2(2*v1, uv1)
			texcoord.SetVector2(2*v2, uv2)
			texcoord.SetVector2(2*v3, uv3)
			texcoord.SetVector2(2*v4, uv4)
		}
	}

	// Top cap
	if top && topRad > 0 {
		// Array of vertex indicesOrig to build used to build the faces.
		idxsOrig := []uint32{}

		// Appends top segments vtxs and builds array of its idxsOrig
		var uv1, uv2, uv3 math32.Vector2
		for x := 0; x < radialSegs; x++ {
			uv1 = uvsOrig[0][x]
			uv2 = uvsOrig[0][x+1]
			uv3 = math32.Vec2(uv2.X, 0)
			// Appends CENTER with its own UV.
			vertex.Set(vidx+idx*3, 0, hHt, 0)
			normal.Set(vidx+idx*3, 0, 1, 0)
			texcoord.SetVector2(tidx+idx*2, uv3)
			idxsOrig = append(idxsOrig, uint32(idx))
			idx++
			// Appends vertex
			v := math32.Vector3{}
			vi := vtxs[0][x]
			vertex.GetVector3(3*vi, &v)
			vertex.SetVector3(vidx+idx*3, v)
			normal.Set(vidx+idx*3, 0, 1, 0)
			texcoord.SetVector2(tidx+idx*2, uv1)
			idxsOrig = append(idxsOrig, uint32(idx))
			idx++
		}
		// Appends copy of first vertex (center)
		var pt, norm math32.Vector3
		var uv math32.Vector2
		vertex.GetVector3(3*int(idxsOrig[0]), &pt)
		normal.GetVector3(3*int(idxsOrig[0]), &norm)
		texcoord.GetVector2(2*int(idxsOrig[0]), &uv)
		vertex.SetVector3(vidx+idx*3, pt)
		normal.SetVector3(vidx+idx*3, norm)
		texcoord.SetVector2(tidx+idx*2, uv)
		idxsOrig = append(idxsOrig, uint32(idx))
		idx++

		// Appends copy of second vertex (v1) USING LAST UVec2
		vertex.GetVector3(3*int(idxsOrig[1]), &pt)
		normal.GetVector3(3*int(idxsOrig[1]), &norm)
		vertex.SetVector3(vidx+idx*3, pt)
		normal.SetVector3(vidx+idx*3, norm)
		texcoord.SetVector2(tidx+idx*2, uv2)
		idxsOrig = append(idxsOrig, uint32(idx))
		idx++

		// Append faces idxsOrig
		for x := 0; x < radialSegs; x++ {
			pos := 2 * x
			i1 := idxsOrig[pos]
			i2 := idxsOrig[pos+1]
			i3 := idxsOrig[pos+3]
			index.Set(ii, uint32(vOff+i1), uint32(vOff+i2), uint32(vOff+i3))
			ii += 3
		}
	}

	// Bottom cap
	if bottom && botRad > 0 {
		// Array of vertex idxsOrig to build used to build the faces.
		idxsOrig := []uint32{}

		// Appends top segments vtxs and builds array of its idxsOrig
		var uv1, uv2, uv3 math32.Vector2
		for x := 0; x < radialSegs; x++ {
			uv1 = uvsOrig[heightSegs][x]
			uv2 = uvsOrig[heightSegs][x+1]
			uv3 = math32.Vec2(uv2.X, 1)
			// Appends CENTER with its own UV.
			vertex.Set(vidx+idx*3, 0, -hHt, 0)
			normal.Set(vidx+idx*3, 0, -1, 0)
			texcoord.SetVector2(tidx+idx*2, uv3)
			idxsOrig = append(idxsOrig, uint32(idx))
			idx++
			// Appends vertex
			v := math32.Vector3{}
			vi := vtxs[heightSegs][x]
			vertex.GetVector3(3*vi, &v)
			vertex.SetVector3(vidx+idx*3, v)
			normal.Set(vidx+idx*3, 0, -1, 0)
			texcoord.SetVector2(tidx+idx*2, uv1)
			idxsOrig = append(idxsOrig, uint32(idx))
			idx++
		}

		// Appends copy of first vertex (center)
		var pt, norm math32.Vector3
		var uv math32.Vector2
		vertex.GetVector3(3*int(idxsOrig[0]), &pt)
		normal.GetVector3(3*int(idxsOrig[0]), &norm)
		texcoord.GetVector2(2*int(idxsOrig[0]), &uv)
		vertex.SetVector3(vidx+idx*3, pt)
		normal.SetVector3(vidx+idx*3, norm)
		texcoord.SetVector2(tidx+idx*2, uv)
		idxsOrig = append(idxsOrig, uint32(idx))
		idx++

		// Appends copy of second vertex (v1) USING LAST UVec2
		vertex.GetVector3(3*int(idxsOrig[1]), &pt)
		normal.GetVector3(3*int(idxsOrig[1]), &norm)
		vertex.SetVector3(vidx+idx*3, pt)
		normal.SetVector3(vidx+idx*3, norm)
		texcoord.SetVector2(tidx+idx*2, uv2)
		idxsOrig = append(idxsOrig, uint32(idx))
		idx++

		// Appends faces idxsOrig
		for x := 0; x < radialSegs; x++ {
			pos := 2 * x
			i1 := idxsOrig[pos]
			i2 := idxsOrig[pos+3]
			i3 := idxsOrig[pos+1]
			index.Set(ii, uint32(vOff+i1), uint32(vOff+i2), uint32(vOff+i3))
			ii += 3
		}
	}

	return bb
}

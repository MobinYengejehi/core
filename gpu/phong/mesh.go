// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phong

import (
	"fmt"
	"log"

	"github.com/MobinYengejehi/core/base/errors"
	"github.com/MobinYengejehi/core/gpu"
	"github.com/MobinYengejehi/core/gpu/shape"
)

// ResetMeshes resets the meshes for reconfiguring.
func (ph *Phong) ResetMeshes() {
	ph.Lock()
	defer ph.Unlock()

	ph.meshes.Reset()
	vgp := ph.System.Vars().VertexGroup()
	vgp.SetNValues(1)
}

// SetMesh sets a Mesh using the [shape.Mesh] interface for the source
// of the mesh data, and sets the values directly.
// If Mesh already exists, then data is updated.
// It is ready for [UseMesh] after this point.
func (ph *Phong) SetMesh(name string, mesh shape.Mesh) {
	ph.Lock()
	defer ph.Unlock()

	vgp := ph.System.Vars().VertexGroup()
	md := shape.NewMeshData(mesh)
	idx, ok := ph.meshes.Map[name]
	if !ok {
		idx = ph.meshes.Len()
		ph.meshes.Add(name, md)
		vgp.SetNValues(ph.meshes.Len())
	} else {
		ph.meshes.Order[idx].Value = md
	}
	ph.configMesh(md, idx)

	gpu.SetValueFrom(errors.Log1(vgp.ValueByIndex("Pos", idx)), md.Vertex)
	gpu.SetValueFrom(errors.Log1(vgp.ValueByIndex("Normal", idx)), md.Normal)
	gpu.SetValueFrom(errors.Log1(vgp.ValueByIndex("TexCoord", idx)), md.TexCoord)
	if md.HasColor {
		gpu.SetValueFrom(errors.Log1(vgp.ValueByIndex("VertexColor", idx)), md.Colors)
	} else if idx == 0 { // set dummy vertexcolor for first guy
		gpu.SetValueFrom(errors.Log1(vgp.ValueByIndex("VertexColor", idx)), make([]float32, md.NumVertex*4))
	}
	gpu.SetValueFrom(errors.Log1(vgp.ValueByIndex("Index", idx)), md.Index)
}

func (ph *Phong) configMesh(md *shape.MeshData, idx int) {
	vgp := ph.System.Vars().VertexGroup()
	errors.Log1(vgp.ValueByIndex("Pos", idx)).SetDynamicN(md.NumVertex)
	errors.Log1(vgp.ValueByIndex("Normal", idx)).SetDynamicN(md.NumVertex)
	errors.Log1(vgp.ValueByIndex("TexCoord", idx)).SetDynamicN(md.NumVertex)
	errors.Log1(vgp.ValueByIndex("Index", idx)).SetDynamicN(md.NumIndex)
	vc := errors.Log1(vgp.ValueByIndex("VertexColor", idx))
	if md.HasColor {
		vc.SetDynamicN(md.NumVertex)
	} else {
		vc.SetDynamicN(1)
	}
}

// UseMesh selects mesh by name for current render step.
// Mesh must have been added / updated via SetMesh method.
// If mesh has per-vertex colors, these are selected for rendering,
// and texture is turned off.  UseTexture* after this to override.
func (ph *Phong) UseMesh(name string) error {
	ph.Lock()
	defer ph.Unlock()

	idx, ok := ph.meshes.IndexByKeyTry(name)
	if !ok {
		err := fmt.Errorf("phong:UseMeshName -- name not found: %s", name)
		if gpu.Debug {
			log.Println(err)
		}
	}
	return ph.useMeshIndex(idx)
}

// useMeshIndex selects mesh by index for current render step.
// If mesh has per-vertex colors, these are selected for rendering,
// and texture is turned off.  UseTexture* after this to override.
func (ph *Phong) useMeshIndex(idx int) error {
	sy := ph.System
	md := ph.meshes.ValueByIndex(idx)
	sy.Vars().SetCurrentValue(gpu.VertexGroup, "Pos", idx)
	sy.Vars().SetCurrentValue(gpu.VertexGroup, "Normal", idx)
	sy.Vars().SetCurrentValue(gpu.VertexGroup, "TexCoord", idx)
	sy.Vars().SetCurrentValue(gpu.VertexGroup, "Index", idx)
	if md.HasColor {
		sy.Vars().SetCurrentValue(gpu.VertexGroup, "VertexColor", idx)
		ph.UseVertexColor = true
		ph.UseCurTexture = false
	} else {
		sy.Vars().SetCurrentValue(gpu.VertexGroup, "VertexColor", 0)
		ph.UseVertexColor = false
	}
	return nil
}

// Copyright (c) 2019, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//go:generate core generate

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"

	"github.com/MobinYengejehi/core/base/iox/imagex"
	"github.com/MobinYengejehi/core/colors"
	"github.com/MobinYengejehi/core/colors/colormap"
	"github.com/MobinYengejehi/core/core"
	"github.com/MobinYengejehi/core/events"
	"github.com/MobinYengejehi/core/gpu"
	"github.com/MobinYengejehi/core/icons"
	"github.com/MobinYengejehi/core/math32"
	"github.com/MobinYengejehi/core/styles"
	"github.com/MobinYengejehi/core/styles/abilities"
	"github.com/MobinYengejehi/core/svg"
	"github.com/MobinYengejehi/core/tree"
	"github.com/MobinYengejehi/core/xyz"
	"github.com/MobinYengejehi/core/xyz/physics"
	"github.com/MobinYengejehi/core/xyz/physics/world"
	"github.com/MobinYengejehi/core/xyz/physics/world2d"
	"github.com/MobinYengejehi/core/xyz/xyzcore"
)

var NoGUI bool

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-nogui" {
		NoGUI = true
	}
	ev := &Env{}
	ev.Defaults()
	if NoGUI {
		ev.NoGUIRun()
		return
	}
	// core.RenderTrace = true
	b := ev.ConfigGUI()
	b.RunMainWindow()
}

// Env encapsulates the virtual environment
type Env struct { //types:add

	// height of emer
	EmerHt float32

	// how far to move every step
	MoveStep float32

	// how far to rotate every step
	RotStep float32

	// width of room
	Width float32

	// depth of room
	Depth float32

	// height of room
	Height float32

	// thickness of walls of room
	Thick float32

	// current depth map
	DepthVals []float32

	// offscreen render camera settings
	Camera world.Camera

	// color map to use for rendering depth map
	DepthMap core.ColorMapName

	// world
	World *physics.Group `display:"-"`

	// 3D view of world
	View3D *world.View

	// view of world
	View2D *world2d.View

	// 3D visualization of the Scene
	SceneEditor *xyzcore.SceneEditor

	// 2D visualization of the Scene
	Scene2D *core.SVG

	// emer group
	Emer *physics.Group `display:"-"`

	// Right eye of emer
	EyeR physics.Body `display:"-"`

	// contacts from last step, for body
	Contacts physics.Contacts `display:"-"`

	// snapshot image
	EyeRImg *core.Image `display:"-"`

	// depth map image
	DepthImage *core.Image `display:"-"`
}

func (ev *Env) Defaults() {
	ev.Width = 10
	ev.Depth = 15
	ev.Height = 2
	ev.Thick = 0.2
	ev.EmerHt = 1
	ev.MoveStep = ev.EmerHt * .2
	ev.RotStep = 15
	ev.DepthMap = core.ColorMapName("ColdHot")
	ev.Camera.Defaults()
	ev.Camera.FOV = 90
}

func (ev *Env) ConfigScene(se *xyz.Scene) {
	ev.SceneEditor.Styler(func(s *styles.Style) {
		se.Background = colors.Scheme.Select.Container
	})
	xyz.NewAmbient(se, "ambient", 0.3, xyz.DirectSun)

	dir := xyz.NewDirectional(se, "dir", 1, xyz.DirectSun)
	dir.Pos.Set(0, 2, 1) // default: 0,1,1 = above and behind us (we are at 0,0,X)

	// grtx := xyz.NewTextureFileFS(assets.Content, se, "ground", "ground.png")
	// floorp := xyz.NewPlane(se, "floor-plane", 100, 100)
	// floor := xyz.NewSolid(se, "floor").SetMesh(floorp).
	// 	SetColor(colors.Tan).SetTexture(grtx).SetPos(0, -5, 0)
	// floor.Mat.Tiling.Repeat.Set(40, 40)
}

// MakeWorld constructs a new virtual physics world
func (ev *Env) MakeWorld() {
	ev.World = physics.NewGroup()
	ev.World.SetName("RoomWorld")

	MakeRoom(ev.World, "room1", ev.Width, ev.Depth, ev.Height, ev.Thick)
	ev.Emer = MakeEmer(ev.World, ev.EmerHt)
	ev.EyeR = ev.Emer.ChildByName("head", 1).AsTree().ChildByName("eye-r", 2).(physics.Body)

	ev.World.WorldInit()
}

// InitWorld does init on world and re-syncs
func (ev *Env) WorldInit() { //types:add
	ev.World.WorldInit()
	if ev.View3D != nil {
		ev.View3D.Sync()
		ev.GrabEyeImg()
	}
	if ev.View2D != nil {
		ev.View2D.Sync()
	}
}

// ReMakeWorld rebuilds the world and re-syncs with gui
func (ev *Env) ReMakeWorld() { //types:add
	ev.MakeWorld()
	ev.View3D.World = ev.World
	if ev.View3D != nil {
		ev.View3D.Sync()
		ev.GrabEyeImg()
	}
	if ev.View2D != nil {
		ev.View2D.Sync()
	}
}

// ConfigView3D makes the 3D view
func (ev *Env) ConfigView3D(sc *xyz.Scene) {
	// sc.MultiSample = 1 // we are using depth grab so we need this = 1
	wgp := xyz.NewGroup(sc)
	wgp.SetName("world")
	ev.View3D = world.NewView(ev.World, sc, wgp)
	ev.View3D.InitLibrary() // this makes a basic library based on body shapes, sizes
	// at this point the library can be updated to configure custom visualizations
	// for any of the named bodies.
	ev.View3D.Sync()
}

// ConfigView2D makes the 2D view
func (ev *Env) ConfigView2D(sc *svg.SVG) {
	wgp := svg.NewGroup(sc.Root)
	wgp.SetName("world")
	ev.View2D = world2d.NewView(ev.World, sc, wgp)
	ev.View2D.InitLibrary() // this makes a basic library based on body shapes, sizes
	// at this point the library can be updated to configure custom visualizations
	// for any of the named bodies.
	ev.View2D.Sync()
}

// RenderEyeImg returns a snapshot from the perspective of Emer's right eye
func (ev *Env) RenderEyeImg() (*image.RGBA, error) {
	err := ev.View3D.RenderOffNode(ev.EyeR, &ev.Camera)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ev.View3D.Image()
}

// GrabEyeImg takes a snapshot from the perspective of Emer's right eye
func (ev *Env) GrabEyeImg() { //types:add
	img, err := ev.RenderEyeImg()
	if err == nil && img != nil {
		ev.EyeRImg.SetImage(img)
		ev.EyeRImg.NeedsRender()
	} else {
		if err != nil {
			log.Println(err)
		}
	}

	depth, err := ev.View3D.DepthImage()
	if err == nil && depth != nil {
		ev.DepthVals = depth
		ev.ViewDepth(depth)
	}
}

// ViewDepth updates depth bitmap with depth data
func (ev *Env) ViewDepth(depth []float32) {
	cmap := colormap.AvailableMaps[string(ev.DepthMap)]
	img := image.NewRGBA(image.Rectangle{Max: ev.Camera.Size})
	ev.DepthImage.SetImage(img)
	world.DepthImage(img, depth, cmap, &ev.Camera)
	ev.DepthImage.NeedsRender()
}

// UpdateViews updates the 2D and 3D views of the scene
func (ev *Env) UpdateViews() {
	if ev.SceneEditor.IsVisible() {
		ev.SceneEditor.NeedsRender()
	}
	if ev.Scene2D.IsVisible() {
		ev.Scene2D.NeedsRender()
	}
}

// WorldStep does one step of the world
func (ev *Env) WorldStep() {
	ev.World.WorldRelToAbs()
	cts := ev.World.WorldCollide(physics.DynsTopGps)
	ev.Contacts = nil
	for _, cl := range cts {
		if len(cl) > 1 {
			for _, c := range cl {
				if c.A.AsTree().Name == "body" {
					ev.Contacts = cl
				}
				fmt.Printf("A: %v  B: %v\n", c.A.AsTree().Name, c.B.AsTree().Name)
			}
		}
	}
	if len(ev.Contacts) > 1 { // turn around
		fmt.Printf("hit wall: turn around!\n")
		rot := 100.0 + 90.0*rand.Float32()
		ev.Emer.Rel.RotateOnAxis(0, 1, 0, rot)
	}
	ev.View3D.UpdatePose()
	ev.View2D.UpdatePose()
	ev.UpdateViews()
	// ev.GrabEyeImg() // todo: this is not working in the first place,
	// and with goroutine rendering in core/renderwindow, it crashes because
	// the main render texture for the view is stale.
}

// StepForward moves Emer forward in current facing direction one step, and takes GrabEyeImg
func (ev *Env) StepForward() { //types:add
	ev.Emer.Rel.MoveOnAxis(0, 0, 1, -ev.MoveStep)
	ev.WorldStep()
}

// StepBackward moves Emer backward in current facing direction one step, and takes GrabEyeImg
func (ev *Env) StepBackward() { //types:add
	ev.Emer.Rel.MoveOnAxis(0, 0, 1, ev.MoveStep)
	ev.WorldStep()
}

// RotBodyLeft rotates emer left and takes GrabEyeImg
func (ev *Env) RotBodyLeft() { //types:add
	ev.Emer.Rel.RotateOnAxis(0, 1, 0, ev.RotStep)
	ev.WorldStep()
}

// RotBodyRight rotates emer right and takes GrabEyeImg
func (ev *Env) RotBodyRight() { //types:add
	ev.Emer.Rel.RotateOnAxis(0, 1, 0, -ev.RotStep)
	ev.WorldStep()
}

// RotHeadLeft rotates emer left and takes GrabEyeImg
func (ev *Env) RotHeadLeft() { //types:add
	hd := ev.Emer.ChildByName("head", 1).(*physics.Group)
	hd.Rel.RotateOnAxis(0, 1, 0, ev.RotStep)
	ev.WorldStep()
}

// RotHeadRight rotates emer right and takes GrabEyeImg
func (ev *Env) RotHeadRight() { //types:add
	hd := ev.Emer.ChildByName("head", 1).(*physics.Group)
	hd.Rel.RotateOnAxis(0, 1, 0, -ev.RotStep)
	ev.WorldStep()
}

// MakeRoom constructs a new room in given parent group with given params
func MakeRoom(par *physics.Group, name string, width, depth, height, thick float32) *physics.Group {
	rm := physics.NewGroup(par)
	rm.SetName(name)
	physics.NewBox(rm).SetSize(math32.Vec3(width, thick, depth)).
		SetColor("grey").SetInitPos(math32.Vec3(0, -thick/2, 0)).SetName("floor")

	physics.NewBox(rm).SetSize(math32.Vec3(width, height, thick)).
		SetColor("blue").SetInitPos(math32.Vec3(0, height/2, -depth/2)).SetName("back-wall")
	physics.NewBox(rm).SetSize(math32.Vec3(thick, height, depth)).
		SetColor("red").SetInitPos(math32.Vec3(-width/2, height/2, 0)).SetName("left-wall")
	physics.NewBox(rm).SetSize(math32.Vec3(thick, height, depth)).
		SetColor("green").SetInitPos(math32.Vec3(width/2, height/2, 0)).SetName("right-wall")
	physics.NewBox(rm).SetSize(math32.Vec3(width, height, thick)).
		SetColor("yellow").SetInitPos(math32.Vec3(0, height/2, depth/2)).SetName("front-wall")
	return rm
}

// MakeEmer constructs a new Emer virtual robot of given height (e.g., 1)
func MakeEmer(par *physics.Group, height float32) *physics.Group {
	emr := physics.NewGroup(par)
	emr.SetName("emer")
	width := height * .4
	depth := height * .15

	physics.NewBox(emr).SetSize(math32.Vec3(width, height, depth)).
		SetColor("purple").SetDynamic(true).
		SetInitPos(math32.Vec3(0, height/2, 0)).SetName("body")
	// body := physics.NewCapsule(emr, "body", math32.Vec3(0, height / 2, 0), height, width/2)
	// body := physics.NewCylinder(emr, "body", math32.Vec3(0, height / 2, 0), height, width/2)

	headsz := depth * 1.5
	hhsz := .5 * headsz
	hgp := physics.NewGroup(emr).SetInitPos(math32.Vec3(0, height+hhsz, 0))
	hgp.SetName("head")

	physics.NewBox(hgp).SetSize(math32.Vec3(headsz, headsz, headsz)).
		SetColor("tan").SetDynamic(true).SetInitPos(math32.Vec3(0, 0, 0)).SetName("head")

	eyesz := headsz * .2
	physics.NewBox(hgp).SetSize(math32.Vec3(eyesz, eyesz*.5, eyesz*.2)).
		SetColor("green").SetDynamic(true).
		SetInitPos(math32.Vec3(-hhsz*.6, headsz*.1, -(hhsz + eyesz*.3))).SetName("eye-l")

	physics.NewBox(hgp).SetSize(math32.Vec3(eyesz, eyesz*.5, eyesz*.2)).
		SetColor("green").SetDynamic(true).
		SetInitPos(math32.Vec3(hhsz*.6, headsz*.1, -(hhsz + eyesz*.3))).SetName("eye-r")

	return emr
}

func (ev *Env) ConfigGUI() *core.Body {
	// vgpu.Debug = true

	b := core.NewBody("virtroom").SetTitle("Emergent Virtual Engine")

	ev.MakeWorld()

	split := core.NewSplits(b)

	tv := core.NewTree(core.NewFrame(split)).SyncTree(ev.World)
	sv := core.NewForm(split).SetStruct(ev)
	imfr := core.NewFrame(split)
	tbvw := core.NewTabs(split)

	scfr, _ := tbvw.NewTab("3D View")
	twofr, _ := tbvw.NewTab("2D View")

	split.SetSplits(.1, .2, .2, .5)

	tv.OnSelect(func(e events.Event) {
		if len(tv.SelectedNodes) > 0 {
			sv.SetStruct(tv.SelectedNodes[0].AsCoreTree().SyncNode)
		}
	})

	//////////////////////////////////////////
	//    3D Scene

	ev.SceneEditor = xyzcore.NewSceneEditor(scfr)
	ev.SceneEditor.UpdateWidget()
	se := ev.SceneEditor.SceneXYZ()
	ev.ConfigScene(se)
	ev.ConfigView3D(se)

	se.Camera.Pose.Pos = math32.Vec3(0, 40, 3.5)
	se.Camera.LookAt(math32.Vec3(0, 5, 0), math32.Vec3(0, 1, 0))
	se.SaveCamera("3")

	se.Camera.Pose.Pos = math32.Vec3(0, 20, 30)
	se.Camera.LookAt(math32.Vec3(0, 5, 0), math32.Vec3(0, 1, 0))
	se.SaveCamera("2")

	se.Camera.Pose.Pos = math32.Vec3(-.86, .97, 2.7)
	se.Camera.LookAt(math32.Vec3(0, .8, 0), math32.Vec3(0, 1, 0))
	se.SaveCamera("1")
	se.SaveCamera("default")

	//////////////////////////////////////////
	//    Image

	imfr.Styler(func(s *styles.Style) {
		s.Direction = styles.Column
	})
	core.NewText(imfr).SetText("Right Eye Image:")
	ev.EyeRImg = core.NewImage(imfr)
	ev.EyeRImg.SetName("eye-r-img")
	ev.EyeRImg.Image = image.NewRGBA(image.Rectangle{Max: ev.Camera.Size})

	core.NewText(imfr).SetText("Right Eye Depth:")
	ev.DepthImage = core.NewImage(imfr)
	ev.DepthImage.SetName("depth-img")
	ev.DepthImage.Image = image.NewRGBA(image.Rectangle{Max: ev.Camera.Size})

	//////////////////////////////////////////
	//    2D Scene

	twov := core.NewSVG(twofr)
	ev.Scene2D = twov
	twov.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
		twov.SVG.Root.ViewBox.Size.Set(ev.Width+4, ev.Depth+4)
		twov.SVG.Root.ViewBox.Min.Set(-0.5*(ev.Width+4), -0.5*(ev.Depth+4))
		twov.SetReadOnly(false)
	})

	ev.ConfigView2D(twov.SVG)

	//////////////////////////////////////////
	//    Toolbar

	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(func(p *tree.Plan) {
			tree.Add(p, func(w *core.Button) {
				w.SetText("Edit Env").SetIcon(icons.Edit).
					SetTooltip("Edit the settings for the environment").
					OnClick(func(e events.Event) {
						sv.SetStruct(ev)
					})
			})
			tree.Add(p, func(w *core.FuncButton) {
				w.SetFunc(ev.WorldInit).SetText("Init").SetIcon(icons.Update)
			})
			tree.Add(p, func(w *core.FuncButton) {
				w.SetFunc(ev.ReMakeWorld).SetText("Make").SetIcon(icons.Update)
			})
			tree.Add(p, func(w *core.FuncButton) {
				w.SetFunc(ev.GrabEyeImg).SetText("Grab Image").SetIcon(icons.Image)
			})
			tree.Add(p, func(w *core.Separator) {})

			tree.Add(p, func(w *core.FuncButton) {
				w.SetFunc(ev.StepForward).SetText("Fwd").SetIcon(icons.SkipNext).
					Styler(func(s *styles.Style) {
						s.SetAbilities(true, abilities.RepeatClickable)
					})
			})
			tree.Add(p, func(w *core.FuncButton) {
				w.SetFunc(ev.StepBackward).SetText("Bkw").SetIcon(icons.SkipPrevious).
					Styler(func(s *styles.Style) {
						s.SetAbilities(true, abilities.RepeatClickable)
					})
			})
			tree.Add(p, func(w *core.FuncButton) {
				w.SetFunc(ev.RotBodyLeft).SetText("Body Left").SetIcon(icons.KeyboardArrowLeft).
					Styler(func(s *styles.Style) {
						s.SetAbilities(true, abilities.RepeatClickable)
					})
			})
			tree.Add(p, func(w *core.FuncButton) {
				w.SetFunc(ev.RotBodyRight).SetText("Body Right").SetIcon(icons.KeyboardArrowRight).
					Styler(func(s *styles.Style) {
						s.SetAbilities(true, abilities.RepeatClickable)
					})
			})
			tree.Add(p, func(w *core.FuncButton) {
				w.SetFunc(ev.RotHeadLeft).SetText("Head Left").SetIcon(icons.KeyboardArrowLeft).
					Styler(func(s *styles.Style) {
						s.SetAbilities(true, abilities.RepeatClickable)
					})
			})
			tree.Add(p, func(w *core.FuncButton) {
				w.SetFunc(ev.RotHeadRight).SetText("Head Right").SetIcon(icons.KeyboardArrowRight).
					Styler(func(s *styles.Style) {
						s.SetAbilities(true, abilities.RepeatClickable)
					})
			})
			tree.Add(p, func(w *core.Separator) {})

			tree.Add(p, func(w *core.Button) {
				w.SetText("README").SetIcon(icons.FileMarkdown).
					SetTooltip("Open browser on README.").
					OnClick(func(e events.Event) {
						core.TheApp.OpenURL("https://github.com/emer/eve/blob/master/examples/virtroom/README.md")
					})
			})
		})
	})
	return b
}

func (ev *Env) NoGUIRun() {
	gp, dev, err := gpu.NoDisplayGPU()
	if err != nil {
		panic(err)
	}
	se := world.NoDisplayScene(gp, dev)
	ev.ConfigScene(se)
	ev.MakeWorld()
	ev.ConfigView3D(se)

	img, err := ev.RenderEyeImg()
	if err == nil {
		imagex.Save(img, "eyer_0.png")
	} else {
		panic(err)
	}
}

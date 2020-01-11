// Copyright (c) 2020, The EFight Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/emer/eve/eve"
	"github.com/emer/eve/evev"
	"github.com/goki/gi/gi"
	"github.com/goki/gi/gi3d"
	"github.com/goki/gi/giv"
	"github.com/goki/gi/mat32"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/oswin/key"
	"github.com/goki/gi/oswin/mouse"
	"github.com/goki/ki/ki"
	"github.com/goki/ki/kit"
)

type Game struct {
	World   *eve.Group
	View    *evev.View
	Scene   *Scene
	Map     Map
	MapObjs map[string]bool
}

// TheGame is the game instance for the current game
var TheGame *Game

type Scene struct {
	gi3d.Scene
	TrackMouse bool
	CamRotUD   float32 // current camera rotation up / down
	CamRotLR   float32 // current camera rotation LR
}

var KiT_Scene = kit.Types.AddType(&Scene{}, nil)

func (gm *Game) BuildMap() {
	for nm, obj := range gm.Map {
		gm.MakeObj(obj, nm)
	}
}

func (gm *Game) MakeObj(obj *MapObj, nm string) *eve.Group {
	var ogp *eve.Group
	switch obj.ObjType {
	case "TheWall":
		ogp = eve.AddNewGroup(gm.World, nm)
		gm.PhysMakeTheWall(ogp, nm)
	case "House":
		ogp = eve.AddNewGroup(gm.World, nm)
		gm.PhysMakeBrickHouse(ogp, nm)
	case "Block":
		ogp = eve.AddNewGroup(gm.World, nm)
		for i := 0; i < 8; i++ {
			house := gm.PhysMakeBrickHouse(ogp, fmt.Sprintf("%v%v", nm, i))
			house.Initial.Pos.Set(float32(20*i), 0, 0)
		}
		/*
			case "Hill":
					ogp = gi3d.AddNewGroup(sc, sc, nm)
					o := gi3d.AddNewObject(sc, ogp, "hill", "Hill")
					o.Pose.Pos.Set(0, 0, 0)
					// o.Pose.Scale.Set(1, 10, 1)
					o.Mat.Color.SetString("green", nil)
				case "House":
					ogp = gi3d.AddNewGroup(sc, sc, nm)
					o := gi3d.AddNewObject(sc, ogp, "house_ground", "HouseFloor")
					o.Pose.Pos.Set(0, 0, 0)
					// o.Pose.Scale.Set(10, 0.01, 10)
					o.Mat.Color.SetString("brown", nil)

					o = gi3d.AddNewObject(sc, ogp, "house_roof", "HouseRoof")
					o.Pose.Pos.Set(0, 4.995, 0)
					// o.Pose.Scale.Set(10, 0.01, 10)
					o.Mat.Color.SetString("brown", nil)

					o = gi3d.AddNewObject(sc, ogp, "house_wall1", "HouseWallOne")
					o.Pose.Pos.Set(-5.25, 0, 0)
					// o.Pose.Scale.Set(0.5, 10, 10)
					o.Mat.Color.SetString("brown", nil)

					o = gi3d.AddNewObject(sc, ogp, "house_wall2", "HouseWallOne")
					o.Pose.Pos.Set(4.75, 0, 0)
					// o.Pose.Scale.Set(0.5, 10, 10)
					o.Mat.Color.SetString("brown", nil)

					o = gi3d.AddNewObject(sc, ogp, "house_wall3", "HouseWallTwo")
					o.Pose.Pos.Set(0, 0, -5)
					// o.Pose.Scale.Set(10, 10, 0.5)
					o.Mat.Color.SetString("brown", nil)

					o = gi3d.AddNewObject(sc, ogp, "house_bed1", "HouseBedOne")
					o.Pose.Pos.Set(-3.5, 0, -4)
					// o.Mat.Color.SetString("green", nil)
					o.Mat.SetTextureName(sc, "HouseBed")

					o = gi3d.AddNewObject(sc, ogp, "house_blanket1", "HouseBlanketOne")
					o.Pose.Pos.Set(-3.5, 1.05, -4)
					o.Mat.SetTextureName(sc, "HouseBlanket")

					o = gi3d.AddNewObject(sc, ogp, "house_pillow1", "HousePillowOne")
					o.Pose.Pos.Set(-4.75, 1.15, -4)
					o.Mat.SetTextureName(sc, "HousePillow")

					o = gi3d.AddNewObject(sc, ogp, "house_couch_base1", "HouseCouchBaseOne")
					o.Pose.Pos.Set(2, 0, -3.5)
					o.Mat.SetTextureName(sc, "HouseCouch")

					o = gi3d.AddNewObject(sc, ogp, "house_couch_top1", "HouseCouchTopOne")
					o.Pose.Pos.Set(2, 1.5, -4.5)
					o.Mat.SetTextureName(sc, "HouseCouch")

					o = gi3d.AddNewObject(sc, ogp, "house_window1", "HouseWindowOne")
					o.Pose.Pos.Set(2, 3.5, -4.8)
					o.Mat.SetTextureName(sc, "HouseWindow")

					o = gi3d.AddNewObject(sc, ogp, "house_window2", "HouseWindowOne")
					o.Pose.Pos.Set(-3, 3.5, -4.8)
					o.Mat.SetTextureName(sc, "HouseWindow")

				case "Center_Blue":
					ogp = gi3d.AddNewGroup(sc, sc, nm)
					o := gi3d.AddNewObject(sc, ogp, "center_blue", "Center_Blue")
					o.Pose.Pos.Set(0, 0, 0)
					o.Mat.Color.SetString("blue", nil)
				case "Table":
					ogp = gi3d.AddNewGroup(sc, sc, nm)
					o := gi3d.AddNewObject(sc, ogp, "table", "Table")
					o.Pose.Pos.Set(0, 0, 0)

					o.Mat.SetTextureName(sc, "Table")
				case "BigComplex":
					ogp = gi3d.AddNewGroup(sc, sc, nm)
					o := gi3d.AddNewObject(sc, ogp, "bigComplexPlaceholder", "BigComplexPlaceholder")
					o.Pose.Pos.Set(0, 20, 0)
					// o.Mat.Color.SetString("red", nil)
					o.Mat.SetTextureName(sc, "Metal1")
				case "TestBed":
					ogp = gi3d.AddNewGroup(sc, sc, nm)
					err := sc.OpenObj([]string{"bed1.obj"}, ogp)
					if err != nil {
						log.Println(err)

					}
		*/

	}
	if ogp != nil {
		ogp.Initial.Pos = obj.Pos
		// ogp.Initial.Scale = obj.Scale
	}
	return ogp
}

// todo: could have a smarter way of figuring out all the stuff you need
// to make, or not..
func (gm *Game) MakeLibrary() {
	gm.LibMakeBrickHouse()
	gm.LibMakeTheWall()
}

func (gm *Game) MakeMeshes() {
	sc := &gm.Scene.Scene
	gi3d.AddNewBox(sc, "Gun", 0.1, 0.1, 1)
	gi3d.AddNewBox(sc, "Person", 0.5, 2, 0.5)
	/*
		gi3d.AddNewBox(sc, "Hill", 1, 10, 1)
		gi3d.AddNewBox(sc, "Table", 5, 2.5, 5)
		gi3d.AddNewBox(sc, "Center_Blue", 3, 2, 3)
		gi3d.AddNewBox(sc, "HouseFloor", 10, 0.01, 10)
		gi3d.AddNewBox(sc, "HouseRoof", 10, 0.01, 10)
		gi3d.AddNewBox(sc, "HouseWallOne", 0.5, 10, 10)
		gi3d.AddNewBox(sc, "HouseWallTwo", 10, 10, 0.5)
		gi3d.AddNewBox(sc, "HouseBedOne", 3, 2, 2)
		gi3d.AddNewBox(sc, "HouseBlanketOne", 3, 0.1, 2)
		gi3d.AddNewBox(sc, "HousePillowOne", 0.5, 0.25, 2)
		gi3d.AddNewBox(sc, "HouseCouchBaseOne", 5, 2, 3)
		gi3d.AddNewBox(sc, "HouseCouchTopOne", 5, 1, 1)
		gi3d.AddNewBox(sc, "HouseWindowOne", 1, 1, 0.1)
		gi3d.AddNewBox(sc, "BigComplexPlaceholder", 40, 40, 40)
	*/
}

// func MakeTextures(sc *gi3d.Scene) {
// 	gi3d.AddNewTextureFile(sc, "Table", "table.jpg")
// 	gi3d.AddNewTextureFile(sc, "HouseBed", "bed.png")
// 	gi3d.AddNewTextureFile(sc, "HouseBlanket", "blanket.png")
// 	gi3d.AddNewTextureFile(sc, "HousePillow", "pillow.png")
// 	gi3d.AddNewTextureFile(sc, "HouseCouch", "couch.jpg")
// 	gi3d.AddNewTextureFile(sc, "HouseWindow", "window.png")
// 	gi3d.AddNewTextureFile(sc, "Metal1", "metal1.jpg")
// 	gi3d.AddNewTextureFile(sc, "Brick1", "brick.jpg")
// }

// MakeWorld constructs a new virtual physics world
func (gm *Game) MakeWorld() {
	gm.World = &eve.Group{}
	gm.World.InitName(gm.World, "World")
	gm.BuildMap()
	gm.World.InitWorld() // key to put things in their places!
}

// MakeView makes the view
func (gm *Game) MakeView() {
	sc := &gm.Scene.Scene
	gm.MakeMeshes()
	wgp := gi3d.AddNewGroup(sc, sc, "world")
	gm.View = evev.NewView(gm.World, sc, wgp)
	//gm.View.InitLibrary() // this makes a basic library based on body shapes, sizes
	gm.MakeLibrary()
	gm.View.Sync()
}

func (gm *Game) Config() {
	gamerow := gi.AddNewLayout(signUpTab, "gamerow", gi.LayoutVert)
	gamerow.SetStretchMaxWidth()
	gamerow.SetStretchMaxHeight()

	sc := AddNewScene(gamerow, "scene")
	gm.Scene = sc
	sc.SetStretchMaxWidth()
	sc.SetStretchMaxHeight()

	// first, add lights, set camera
	sc.BgColor.SetUInt8(230, 230, 255, 255) // sky blue-ish
	gi3d.AddNewAmbientLight(&sc.Scene, "ambient", 0.5, gi3d.DirectSun)

	dir := gi3d.AddNewDirLight(&sc.Scene, "dir", 1, gi3d.DirectSun)
	dir.Pos.Set(0, 1, 1) // default: 0,1,1 = above and behind us (we are at 0,0,X)

	// point := gi3d.AddNewPointLight(sc, "point", 1, gi3d.DirectSun)
	// point.Pos.Set(0, 5, 5)

	// spot := gi3d.AddNewSpotLight(sc, "spot", 1, gi3d.DirectSun)
	// spot.Pose.Pos.Set(0, 0, 2)
	sc.Camera.Pose.Pos.Y = 2
	sc.Camera.Pose.Pos.Z = 20

	gm.Map = FirstMap
	gm.MakeWorld()

	gm.MakeView()
	gi.AddNewSpace(gamerow, "space1")
	epbut := gi.AddNewButton(gamerow, "edit-phys")
	epbut.SetText("Edit Phys")
	epbut.ButtonSig.Connect(gm.World.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			giv.GoGiEditorDialog(gm.World)
		}
	})

	// center_bluem :=
	// cbm.Segs.Set(10, 10, 10) // not clear if any diff really..

	fpobj := gi3d.AddNewGroup(&sc.Scene, &sc.Scene, "TrackCamera")
	rcb := gi3d.AddNewSolid(&sc.Scene, fpobj, "red-cube", "Person")
	rcb.Pose.Pos.Set(0, -1, -8)
	// rcb.Pose.Scale.Set(0.1, 0.1, 1)
	rcb.Mat.Color.SetString("red", nil)

	// center_blue := sc.AddNewObject("center_blue", center_bluem.Name())
	// center_blue.Pose.Pos.Set(0, 0, 0)
	// center_blue.Mat.Color.SetString("blue", nil)
	//
	// green_hill := sc.AddNewObject("green_hill", cbm.Name())
	// green_hill.Pose.Pos.Set(1, 0, -20)
	// green_hill.Pose.Scale.Set(1, 10, 1)
	// green_hill.Mat.Color.SetString("green", nil)
	//
	// tbtx := gi3d.AddNewTextureFile(&sc.Scene, "table", "table.jpg")
	// var posy float32 = 0
	// var posx float32 = -5
	// var posz float32 = -20
	// for r := 0; r < 4; r++ {
	// 	posx = -5
	// 	for c := 0; c < 4; c++ {
	// 		market := sc.AddNewObject(fmt.Sprintf("market%v", c*r), cbm.Name())
	// 		market.Pose.Pos.Set(posx, posy, posz)
	// 		market.Pose.Scale.Set(5, 2.5, 5)
	// 		// market1.Mat.Color.SetString("red", nil)
	// 		market.Mat.SetTexture(&sc.Scene, tbtx.Name())
	// 		posx = posx - 6.5
	// 	}
	// 	posz = posz - 6.5
	//
	// }
	//
	// // market1 := sc.AddNewObject("market1", cbm.Name())
	// // market1.Pose.Pos.Set(-5, 0, -20)
	// // market1.Pose.Scale.Set(5, 2.5, 5)
	// // // market1.Mat.Color.SetString("red", nil)
	// // market1.Mat.SetTexture(&sc.Scene, tbtx.Name())

	floorp := gi3d.AddNewPlane(&sc.Scene, "floor-plane", 200, 200)
	floor := gi3d.AddNewSolid(&sc.Scene, &sc.Scene, "floor", floorp.Name())
	floor.Pose.Pos.Set(0, 0, 0)
	// floor.Mat.Emissive.SetString("green", nil)
	grtx := gi3d.AddNewTextureFile(&sc.Scene, "ground", "ground.jpg")
	floor.Mat.SetTexture(&sc.Scene, grtx)
	floor.Mat.Tiling.Repeat.Set(200, 200)

}

func AddNewScene(parent ki.Ki, name string) *Scene {
	sc := parent.AddNewChild(KiT_Scene, name).(*Scene)
	sc.Defaults()
	return sc
}

func (sc *Scene) Render2D() {
	if sc.PushBounds() {
		if !sc.NoNav {
			sc.NavEvents()
		}
		if gi.Render2DTrace {
			fmt.Printf("3D Render2D: %v\n", sc.PathUnique())
		}
		sc.Render()
		sc.PopBounds()
	} else {
		sc.DisconnectAllEvents(gi.RegPri)
	}
}

func (sc *Scene) NavEvents() {
	sc.ConnectEvent(oswin.MouseMoveEvent, gi.RegPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		if !sc.TrackMouse {
			return
		}
		me := d.(*mouse.MoveEvent)
		me.SetProcessed()
		ssc := recv.Embed(KiT_Scene).(*Scene)
		orbDel := float32(.2)
		orbDels := orbDel * 0.05
		panDel := float32(.05)
		del := me.Where.Sub(me.From)
		dx := float32(-del.X)
		dy := float32(-del.Y)
		// fmt.Printf("pos: %v  fm: %v del: %v\n", me.Where, me.From, del)
		switch {
		case key.HasAllModifierBits(me.Modifiers, key.Shift):
			ssc.Camera.Pan(dx*panDel, -dy*panDel)
		case key.HasAllModifierBits(me.Modifiers, key.Control):
			ssc.Camera.PanAxis(dx*panDel, -dy*panDel)
		case key.HasAllModifierBits(me.Modifiers, key.Alt):
			ssc.Camera.PanTarget(dx*panDel, -dy*panDel, 0)
		default:
			if mat32.Abs(dx) > mat32.Abs(dy) {
				dy = 0
			} else {
				dx = 0
			}
			sc.CamRotUD += dy * orbDels * .2
			sc.CamRotLR += dx * orbDels * 2
			ssc.Camera.Pose.SetAxisRotation(0, 1, 0, sc.CamRotLR)
			ssc.Camera.Pose.RotateOnAxis(1, 0, 0, sc.CamRotUD)
		}
		ssc.UpdateSig()
	})
	sc.ConnectEvent(oswin.MouseScrollEvent, gi.RegPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		me := d.(*mouse.ScrollEvent)
		me.SetProcessed()
		ssc := recv.Embed(KiT_Scene).(*Scene)
		if ssc.SetDragCursor {
			oswin.TheApp.Cursor(ssc.Viewport.Win.OSWin).Pop()
			ssc.SetDragCursor = false
		}
		zoom := float32(me.NonZeroDelta(false))
		zoomPct := float32(.05)
		zoomDel := float32(.05)
		switch {
		case key.HasAllModifierBits(me.Modifiers, key.Alt):
			ssc.Camera.PanTarget(0, 0, zoom*zoomDel)
		default:
			ssc.Camera.Zoom(zoomPct * zoom)
		}
		ssc.UpdateSig()
	})
	sc.ConnectEvent(oswin.MouseEvent, gi.RegPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		me := d.(*mouse.Event)
		me.SetProcessed()
		ssc := recv.Embed(KiT_Scene).(*Scene)
		if ssc.SetDragCursor {
			oswin.TheApp.Cursor(ssc.Viewport.Win.OSWin).Pop()
			ssc.SetDragCursor = false
		}
		if !ssc.IsInactive() && !ssc.HasFocus() {
			ssc.GrabFocus()
		}
		// obj := ssc.FirstContainingPoint(me.Where, true)
		// if me.Action == mouse.Release && me.Button == mouse.Right {
		// 	me.SetProcessed()
		// 	if obj != nil {
		// 		giv.StructViewDialog(ssc.Viewport, obj, giv.DlgOpts{Title: "sc Element View"}, nil, nil)
		// 	}
		// }
	})
	sc.ConnectEvent(oswin.MouseHoverEvent, gi.RegPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		me := d.(*mouse.HoverEvent)
		me.SetProcessed()
		// ssc := recv.Embed(KiT_Scene).(*Scene)
		// obj := ssc.FirstContainingPoint(me.Where, true)
		// if obj != nil {
		// 	pos := me.Where
		// 	ttxt := fmt.Sprintf("element name: %v -- use right mouse click to edit", obj.Name())
		// 	gi.PopupTooltip(obj.Name(), pos.X, pos.Y, sc.Viewport, ttxt)
		// }
	})
	sc.ConnectEvent(oswin.KeyChordEvent, gi.RegPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		ssc := recv.Embed(KiT_Scene).(*Scene)
		kt := d.(*key.ChordEvent)
		ssc.NavKeyEvents(kt)
	})
}

func (sc *Scene) NavKeyEvents(kt *key.ChordEvent) {
	ch := string(kt.Chord())
	// fmt.Printf(ch)
	// orbDeg := float32(5)
	panDel := float32(.1)
	// zoomPct := float32(.05)
	switch ch {
	case "Escape":
		sc.TrackMouse = !sc.TrackMouse
		if sc.TrackMouse {
			sc.Win.OSWin.SetCursorEnabled(false, true) // turn off mouse cursor
		} else {
			sc.Win.OSWin.SetCursorEnabled(true, false)
		}
		kt.SetProcessed()

	// case "UpArrow":
	//
	// 	sc.Camera.Pose.SetEulerRotation(orbDeg, 0, 0)
	// kt.SetProcessed()

	// case "Shift+UpArrow":
	// 	sc.Camera.Pan(0, panDel)
	// 	kt.SetProcessed()
	// case "Control+UpArrow":
	// 	sc.Camera.PanAxis(0, panDel)
	// 	kt.SetProcessed()
	// case "Alt+UpArrow":
	// 	sc.Camera.PanTarget(0, panDel, 0)
	// 	kt.SetProcessed()
	// case "DownArrow":
	// sc.Camera.Orbit(0, -orbDeg)
	// kt.SetProcessed()
	// case "Shift+DownArrow":
	// 	sc.Camera.Pan(0, -panDel)
	// 	kt.SetProcessed()
	// case "Control+DownArrow":
	// 	sc.Camera.PanAxis(0, -panDel)
	// 	kt.SetProcessed()
	// case "Alt+DownArrow":
	// 	sc.Camera.PanTarget(0, -panDel, 0)
	// 	kt.SetProcessed()
	// case "LeftArrow":
	// sc.Camera.Orbit(orbDeg, 0)
	// kt.SetProcessed()
	// case "Shift+LeftArrow":
	// 	sc.Camera.Pan(-panDel, 0)
	// 	kt.SetProcessed()
	// case "Control+LeftArrow":
	// 	sc.Camera.PanAxis(-panDel, 0)
	// 	kt.SetProcessed()
	// case "Alt+LeftArrow":
	// 	sc.Camera.PanTarget(-panDel, 0, 0)
	// 	kt.SetProcessed()
	// case "RightArrow":
	// sc.Camera.Orbit(-orbDeg, 0)
	// kt.SetProcessed()
	// case "Shift+RightArrow":
	// 	sc.Camera.Pan(panDel, 0)
	// 	kt.SetProcessed()
	// case "Control+RightArrow":
	// 	sc.Camera.PanAxis(panDel, 0)
	// 	kt.SetProcessed()
	// case "Alt+RightArrow":
	// 	sc.Camera.PanTarget(panDel, 0, 0)
	// 	kt.SetProcessed()
	// case "Alt++", "Alt+=":
	// 	sc.Camera.PanTarget(0, 0, panDel)
	// 	kt.SetProcessed()
	// case "Alt+-", "Alt+_":
	// 	sc.Camera.PanTarget(0, 0, -panDel)
	// 	kt.SetProcessed()
	// case "+", "=":
	// 	sc.Camera.Zoom(-zoomPct)
	// 	kt.SetProcessed()
	// case "-", "_":
	// 	sc.Camera.Zoom(zoomPct)
	// 	kt.SetProcessed()
	case " ":
		err := sc.SetCamera("default")
		if err != nil {
			sc.Camera.DefaultPose()
		}
	case "w":
		y := sc.Camera.Pose.Pos.Y
		sc.Camera.Pose.MoveOnAxis(0, 0, -0.5, .5)
		kt.SetProcessed()
		sc.Camera.Pose.Pos.Y = y
	case "s":
		y := sc.Camera.Pose.Pos.Y
		sc.Camera.Pose.MoveOnAxis(0, 0, 0.5, .5)
		kt.SetProcessed()
		sc.Camera.Pose.Pos.Y = y
	case "a":
		sc.Camera.Pan(panDel, 0)
		kt.SetProcessed()
	case "d":
		sc.Camera.Pan(-panDel, 0)
		kt.SetProcessed()
	case "t":
		kt.SetProcessed()
		obj := sc.Child(0).(*gi3d.Solid)
		fmt.Printf("updated obj: %v\n", obj.PathUnique())
		obj.UpdateSig()
		return
	}
	sc.UpdateSig()
}

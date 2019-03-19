package main

import (
	"fmt"
	"rapidengine/cmd"
	"rapidengine/geometry"
	"rapidengine/material"
)

var backMat1 *material.BasicMaterial
var backMat2 *material.BasicMaterial
var backMat3 *material.BasicMaterial
var backMat4 *material.BasicMaterial
var backMat5 *material.BasicMaterial
var backMat6 *material.BasicMaterial

func InitializeWorldScene() {

	Engine.TextureControl.NewTexture("assets/backgrounds/gradient.png", "sky", "pixel")

	Engine.TextureControl.NewTexture("assets/backgrounds/og1/1.png", "parallax1", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/og1/2.png", "parallax2", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/og1/3.png", "parallax3", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/og1/4.png", "parallax4", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/og1/5.png", "parallax5", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/og1/6.png", "parallax6", "pixel")

	/*Engine.TextureControl.NewTexture("assets/backgrounds/forest/trees1.png", "parallax5", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/forest/trees2.png", "parallax6", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/forest/trees3.png", "parallax7", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/mountain/parallax4.png", "parallax8", "pixel")

	Engine.TextureControl.NewTexture("assets/backgrounds/snow/snow1.png", "parallax9", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/snow/snow2.png", "parallax10", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/snow/snow3.png", "parallax11", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/snow/snow4.png", "parallax12", "pixel")

	Engine.TextureControl.NewTexture("assets/backgrounds/mountain2/mountain1.png", "parallax13", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/mountain2/mountain2.png", "parallax14", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/mountain2/mountain3.png", "parallax15", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/mountain2/mountain4.png", "parallax16", "pixel")*/

	backgroundMaterial := Engine.MaterialControl.NewBasicMaterial()
	backgroundMaterial.DiffuseLevel = 1
	backgroundMaterial.DiffuseMap = Engine.TextureControl.GetTexture("sky")
	backgroundMaterial.DiffuseMapScale = 1

	backMat1 = Engine.MaterialControl.NewBasicMaterial()
	backMat1.DiffuseLevel = 1
	backMat1.Blending = true
	backMat1.DiffuseMap = Engine.TextureControl.GetTexture("parallax1")
	backMat1.DiffuseMapScale = float32(Config.ScreenWidth) / float32(WorldWidth*BlockSize)

	backMat2 = Engine.MaterialControl.NewBasicMaterial()
	backMat2.DiffuseLevel = 1
	backMat2.Blending = true
	backMat2.DiffuseMap = Engine.TextureControl.GetTexture("parallax2")
	backMat2.DiffuseMapScale = float32(Config.ScreenWidth) / float32(WorldWidth*BlockSize)

	backMat3 = Engine.MaterialControl.NewBasicMaterial()
	backMat3.DiffuseLevel = 1
	backMat3.Blending = true
	backMat3.DiffuseMap = Engine.TextureControl.GetTexture("parallax3")
	backMat3.DiffuseMapScale = float32(Config.ScreenWidth) / float32(WorldWidth*BlockSize)

	backMat4 = Engine.MaterialControl.NewBasicMaterial()
	backMat4.DiffuseLevel = 1
	backMat4.Blending = true
	backMat4.DiffuseMap = Engine.TextureControl.GetTexture("parallax4")
	backMat4.DiffuseMapScale = float32(Config.ScreenWidth) / float32(WorldWidth*BlockSize)

	backMat5 = Engine.MaterialControl.NewBasicMaterial()
	backMat5.DiffuseLevel = 1
	backMat5.Blending = true
	backMat5.DiffuseMap = Engine.TextureControl.GetTexture("parallax5")
	backMat5.DiffuseMapScale = float32(Config.ScreenWidth) / float32(WorldWidth*BlockSize)

	backMat6 = Engine.MaterialControl.NewBasicMaterial()
	backMat6.DiffuseLevel = 1
	backMat6.Blending = true
	backMat6.DiffuseMap = Engine.TextureControl.GetTexture("parallax6")
	backMat6.DiffuseMapScale = float32(Config.ScreenWidth) / float32(WorldWidth*BlockSize)

	Back1Child = Engine.ChildControl.NewChild2D()
	Back1Child.AttachMaterial(backMat1)
	Back1Child.AttachMesh(geometry.NewRectangle())
	Back1Child.ScaleX = float32(WorldWidth * BlockSize)
	Back1Child.ScaleY = float32(Config.ScreenHeight)

	Back2Child = Engine.ChildControl.NewChild2D()
	Back2Child.AttachMaterial(backMat2)
	Back2Child.AttachMesh(geometry.NewRectangle())
	Back2Child.ScaleX = float32(WorldWidth * BlockSize)
	Back2Child.ScaleY = float32(Config.ScreenHeight)

	Back3Child = Engine.ChildControl.NewChild2D()
	Back3Child.AttachMaterial(backMat3)
	Back3Child.AttachMesh(geometry.NewRectangle())
	Back3Child.ScaleX = float32(WorldWidth * BlockSize)
	Back3Child.ScaleY = float32(Config.ScreenHeight)

	Back4Child = Engine.ChildControl.NewChild2D()
	Back4Child.AttachMaterial(backMat4)
	Back4Child.AttachMesh(geometry.NewRectangle())
	Back4Child.ScaleX = float32(WorldWidth * BlockSize)
	Back4Child.ScaleY = float32(Config.ScreenHeight)

	Back5Child = Engine.ChildControl.NewChild2D()
	Back5Child.AttachMaterial(backMat5)
	Back5Child.AttachMesh(geometry.NewRectangle())
	Back5Child.ScaleX = float32(WorldWidth * BlockSize)
	Back5Child.ScaleY = float32(Config.ScreenHeight)

	Back6Child = Engine.ChildControl.NewChild2D()
	Back6Child.AttachMaterial(backMat6)
	Back6Child.AttachMesh(geometry.NewRectangle())
	Back6Child.ScaleX = float32(WorldWidth * BlockSize)
	Back6Child.ScaleY = float32(Config.ScreenHeight)

	SkyChild = Engine.ChildControl.NewChild2D()
	SkyChild.AttachMaterial(backgroundMaterial)
	SkyChild.AttachMesh(geometry.NewRectangle())
	SkyChild.ScaleX = float32(ScreenWidth)
	SkyChild.ScaleY = float32(ScreenHeight)

	InitializePlayer()

	m := Engine.MaterialControl.NewBasicMaterial()
	m.Hue = [4]float32{200, 200, 200, 0.5}

	BlockSelect = Engine.ChildControl.NewChild2D()
	BlockSelect.AttachMaterial(m)
	BlockSelect.AttachMesh(geometry.NewRectangle())
	BlockSelect.ScaleX = 32
	BlockSelect.ScaleY = 32

	WorldChild = Engine.ChildControl.NewChild2D()
	WorldChild.AttachMesh(geometry.NewRectangle())
	WorldChild.ScaleX = BlockSize
	WorldChild.ScaleY = BlockSize
	WorldChild.EnableCopying()
	WorldChild.AttachCollider(0, 0, BlockSize, BlockSize)

	NoCollisionChild = Engine.ChildControl.NewChild2D()
	NoCollisionChild.AttachMesh(geometry.NewRectangle())
	NoCollisionChild.ScaleX = BlockSize
	NoCollisionChild.ScaleY = BlockSize
	NoCollisionChild.EnableCopying()

	NatureChild = Engine.ChildControl.NewChild2D()
	NatureChild.AttachMesh(geometry.NewRectangle())
	NatureChild.ScaleX = BlockSize
	NatureChild.ScaleY = BlockSize
	NatureChild.EnableCopying()

	GrassChild = Engine.ChildControl.NewChild2D()
	GrassChild.AttachMesh(geometry.NewRectangle())
	GrassChild.ScaleX = BlockSize
	GrassChild.ScaleY = BlockSize / 1.5
	GrassChild.EnableCopying()

	Engine.TextureControl.NewTexture("./assets/cloud1.png", "cloud1", "pixel")
	cloudMaterial = Engine.MaterialControl.NewBasicMaterial()
	cloudMaterial.DiffuseLevel = 1
	cloudMaterial.DiffuseMap = Engine.TextureControl.GetTexture("cloud1")
	CloudChild = Engine.ChildControl.NewChild2D()
	CloudChild.AttachMaterial(cloudMaterial)
	CloudChild.AttachMesh(geometry.NewRectangle())
	CloudChild.ScaleX = 300
	CloudChild.ScaleY = 145
	CloudChild.EnableCopying()
	CloudChild.SetSpecificRenderDistance(float32(ScreenWidth/2) + 300)

	Engine.TextureControl.NewTexture("./assets/sun.png", "sun", "pixel")
	sunMat := Engine.MaterialControl.NewBasicMaterial()
	sunMat.Shader = Engine.ShaderControl.GetShader("sun")
	sunMat.DiffuseLevel = 1
	sunMat.DiffuseMap = Engine.TextureControl.GetTexture("sun")
	sunMat.ScatterLevel = 1
	SunChild = Engine.ChildControl.NewChild2D()
	SunChild.AttachMesh(geometry.NewRectangle())
	SunChild.AttachMaterial(sunMat)
	SunChild.ScaleX = 200
	SunChild.ScaleY = 200
	SunChild.Static = true
	SunChild.SetPosition(1500, 700)
	cmd.SunX = (1500.0 + 100) / 1920.0
	cmd.SunY = (700.0 + 100) / 1080.0

	initializeWorldTree()

	//   --------------------------------------------------
	//   Instancing
	//   --------------------------------------------------

	WorldScene = Engine.SceneControl.NewScene("world")
	WorldScene.DisableAutomaticRendering()

	WorldScene.InstanceChild(SkyChild)

	WorldScene.InstanceChild(SunChild)

	WorldScene.InstanceChild(Back6Child)
	WorldScene.InstanceChild(Back5Child)
	WorldScene.InstanceChild(Back4Child)
	WorldScene.InstanceChild(Back3Child)
	WorldScene.InstanceChild(Back2Child)
	WorldScene.InstanceChild(Back1Child)

	WorldScene.InstanceChild(CloudChild)
	WorldScene.InstanceChild(NoCollisionChild)
	WorldScene.InstanceChild(NatureChild)
	WorldScene.InstanceChild(WorldChild)
	WorldScene.InstanceChild(GrassChild)
	WorldScene.InstanceChild(Player1.PlayerChild)
	WorldScene.InstanceChild(BlockSelect)

	Engine.UIControl.InstanceElement(Player1Health, WorldScene)
}

var CurrentParallax = int(1)
var JustParallax = false

func ChangeParallax() {
	if JustParallax {
		return
	}

	JustParallax = true
	println(CurrentParallax)

	if CurrentParallax == 13 {
		CurrentParallax = 1
	} else {
		CurrentParallax += 4
	}

	backMat1.DiffuseMap = Engine.TextureControl.GetTexture("parallax" + fmt.Sprint(CurrentParallax))
	backMat2.DiffuseMap = Engine.TextureControl.GetTexture("parallax" + fmt.Sprint(CurrentParallax+1))
	backMat3.DiffuseMap = Engine.TextureControl.GetTexture("parallax" + fmt.Sprint(CurrentParallax+2))
	backMat4.DiffuseMap = Engine.TextureControl.GetTexture("parallax" + fmt.Sprint(CurrentParallax+3))
}

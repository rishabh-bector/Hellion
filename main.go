package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"rapidengine"
	"rapidengine/configuration"
	"rapidengine/input"
	"runtime"
)

var engine rapidengine.Engine
var config configuration.EngineConfig

var WorldChild rapidengine.Child2D
var Player rapidengine.Child2D
var SkyChild rapidengine.Child2D

var blocks []string

func init() {
	runtime.LockOSThread()
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()
	config = rapidengine.NewEngineConfig(1920, 1080, 2)
	engine = rapidengine.NewEngine(config, render)

	engine.Renderer.SetRenderDistance(1000)
	engine.Renderer.MainCamera.SetPosition(100, 100, 0)
	engine.Renderer.MainCamera.SetSpeed(0.2)

	engine.TextureControl.NewTexture("./assets/player/player.png", "player")
	engine.TextureControl.NewTexture("./assets/blocks/dirt.png", "dirt")
	engine.TextureControl.NewTexture("./assets/blocks/grass.png", "grass")
	engine.TextureControl.NewTexture("./assets/back.png", "back")

	blocks = append(blocks, "player")
	blocks = append(blocks, "dirt")
	blocks = append(blocks, "grass")

	WorldChild = engine.NewChild2D()
	WorldChild.AttachPrimitive(rapidengine.NewRectangle(BlockSize, BlockSize, &config))
	WorldChild.AttachTexturePrimitive(engine.TextureControl.GetTexture("grass"))
	WorldChild.EnableCopying()
	WorldChild.AttachCollider(0, 0, BlockSize, BlockSize)

	engine.Config.Logger.Info("Generating world...")
	generateWorld()
	createCopies()

	engine.CollisionControl.CreateGroup("ground")
	engine.CollisionControl.AddChildToGroup(&WorldChild, "ground")

	Player = engine.NewChild2D()
	Player.AttachPrimitive(rapidengine.NewRectangle(30, 50, &config))
	Player.AttachTexturePrimitive(engine.TextureControl.GetTexture("player"))
	Player.SetPosition(3000, 20000)
	Player.AttachCollider(0, 0, 30, 50)
	Player.SetGravity(1)

	Player.EnableAnimation()
	Player.SetAnimationSpeed(1)
	Player.AddFrame(engine.TextureControl.GetTexture("dirt"))
	Player.AddFrame(engine.TextureControl.GetTexture("player"))

	SkyChild = engine.NewChild2D()
	SkyChild.AttachPrimitive(rapidengine.NewRectangle(2000, 1150, &config))
	SkyChild.AttachTexturePrimitive(engine.TextureControl.GetTexture("back"))

	engine.Instance(&SkyChild)
	engine.Instance(&WorldChild)
	engine.Instance(&Player)

	engine.Initialize()

	engine.StartRenderer()
	<-engine.Done()
}

func render(renderer *rapidengine.Renderer, inputs *input.Input) {
	renderer.MainCamera.SetPosition(Player.X, Player.Y, 0)
	SkyChild.SetPosition(Player.X-1950/2, Player.Y-1110/2)
	movePlayer(inputs.Keys)
}

func movePlayer(keys map[string]bool) {
	if keys["w"] {
		Player.SetVelocityY(10)
	}
	if keys["a"] {
		Player.SetVelocityX(5)
	} else if keys["d"] {
		Player.SetVelocityX(-5)
	} else {
		Player.SetVelocityX(0)
	}
}

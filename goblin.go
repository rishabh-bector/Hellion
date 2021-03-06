package main

//"rapidengine/child"
//"math"

import (
	"fmt"
	"rapidengine/child"

	"rapidengine/material"
)

type Goblin struct {
	common    *Common
	activator Activator
}

func (g *Goblin) Update() {
	g.common.Update()
	Engine.Renderer.RenderChild(g.common.MonsterChild)
}

func (g *Goblin) GetChild() *child.Child2D {
	return g.common.MonsterChild
}

func (g *Goblin) GetCommon() *Common {
	return g.common
}

func (g *Goblin) Damage(amount float32) {
	g.common.Health -= amount
	fmt.Printf("Goblin hit! Health: %v \n", g.common.Health)
}

func (g *Goblin) Activator() *Activator {
	return &g.activator
}

func LoadGoblinTextures() {
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/idle/1.png", "goblin_i1", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/idle/2.png", "goblin_i2", "pixel")

	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/1.png", "goblin_a1", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/2.png", "goblin_a2", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/3.png", "goblin_a3", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/4.png", "goblin_a4", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/5.png", "goblin_a5", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/6.png", "goblin_a6", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/7.png", "goblin_a7", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/8.png", "goblin_a8", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/9.png", "goblin_a9", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/10.png", "goblin_a10", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/11.png", "goblin_a11", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/12.png", "goblin_a12", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/13.png", "goblin_a13", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/14.png", "goblin_a14", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/15.png", "goblin_a15", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/16.png", "goblin_a16", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/attack/17.png", "goblin_a17", "pixel")

	Engine.TextureControl.NewTexture("./assets/enemies/goblin/walk/1.png", "goblin_w1", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/walk/2.png", "goblin_w2", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/walk/3.png", "goblin_w3", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/goblin/walk/4.png", "goblin_w4", "pixel")
}

func NewGoblinMaterial() *material.BasicMaterial {
	goblinMaterial := Engine.MaterialControl.NewBasicMaterial()
	goblinMaterial.DiffuseLevel = 1
	goblinMaterial.DiffuseMap = Engine.TextureControl.GetTexture("goblin_w1")

	goblinMaterial.EnableAnimation()

	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_i1"), "idle")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_i2"), "idle")

	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_i1"), "jump")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_i2"), "jump")

	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a1"), "attack")
	goblinMaterial.AddHitFrame(Engine.TextureControl.GetTexture("goblin_a2"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a3"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a4"), "attack")
	goblinMaterial.AddHitFrame(Engine.TextureControl.GetTexture("goblin_a5"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a6"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a7"), "attack")
	goblinMaterial.AddHitFrame(Engine.TextureControl.GetTexture("goblin_a8"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a9"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a10"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a11"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a12"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a13"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a14"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a15"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a16"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a17"), "attack")

	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w1"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w2"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w3"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w4"), "walk")

	goblinMaterial.SetAnimationFPS("walk", 20)
	goblinMaterial.SetAnimationFPS("idle", 5)
	goblinMaterial.SetAnimationFPS("jump", 5)
	goblinMaterial.SetAnimationFPS("hit", 5)
	goblinMaterial.SetAnimationFPS("attack", 10)

	goblinMaterial.PlayAnimation("idle")

	return goblinMaterial
}

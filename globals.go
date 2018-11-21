package main

import (
	"rapidengine/child"
	"rapidengine/cmd"
	"rapidengine/configuration"
	"rapidengine/lighting"
	"rapidengine/material"
)

//  --------------------------------------------------
//  Globals.go contains all the global variables in the project.
//  This is not good practice.
//  --------------------------------------------------

// Rapid Engine
var Engine *cmd.Engine
var Config configuration.EngineConfig

// Screen Size
var ScreenWidth = 1920
var ScreenHeight = 1080

//  --------------------------------------------------
//  Children
//  --------------------------------------------------

var BlockSelect *child.Child2D

// World
var WorldChild *child.Child2D
var SkyChild *child.Child2D
var NoCollisionChild *child.Child2D
var NatureChild *child.Child2D
var CloudChild *child.Child2D

var l lighting.PointLight

//  --------------------------------------------------
//  World Generation
//  --------------------------------------------------

var Seed = int64(0)

// Size
const WorldWidth = 3000
const WorldHeight = 2000
const BlockSize = 32

// Height
const Flatness = 0.25
const GrassMinimum = 1500

// Cave generation
const CaveStartingThreshold = 0.27
const CaveEndingThreshold = 0.42
const CaveThresholdDelta = 0.002

const CaveIterations = 20
const CaveBirthLimit = 4
const CaveDeathLimit = 3

const SecondCaveIterations = 5
const SecondCaveBirthLimit = 3
const SecondCaveDeathLimit = 2

// Stone generation
const StoneFrequencyDelta = 0.001
const StoneStartingFrequency = 0.32
const StoneEndingFrequency = 0.77
const StoneTopDeviation = 10

// Data
var WorldMap WorldTree
var HeightMap [WorldWidth]int
var CaveMap [][]bool

//  --------------------------------------------------
//  Data
//  --------------------------------------------------

var TransparentBlocks = []string{"backdirt", "torch"} //"topGrass1", "topGrass2", "topGrass3", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot", "treeBranchR1", "treeBranchL1", "flower1", "flower2", "flower3", "pebble"}
var natureBlocks = []string{"leaves", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot", "treeBranchR1", "treeBranchL1", "topGrass1", "topGrass2", "topGrass3", "flower1", "flower2", "flower3", "pebble"}

var cloudMaterial *material.BasicMaterial

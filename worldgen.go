package main

import (
	"math/rand"
	"rapidengine/child"
	"time"
	"fmt"

	perlin "github.com/aquilax/go-perlin"
)

func generateWorldTree() {
	// Make sure all blocks are loaded
	loadBlocks()

	// Make sure all children are loaded
	loadWorldChildren()

	// Randomize seed
	randomizeSeed()

	// Create a blank world tree
	WorldMap = NewWorldTree()

	// Generate heightmap and place grass
	generateHeights()

	// Fill everything underneath grass with dirt
	fillHeights()

	// Generate stone based on height
	fillStone()

	// Clean up stone above ground
	cleanStone()

	// Generate caves
	generateCaves()

	// Clean back dirt
	cleanBackDirt()

	// Put grass and some top grass on dirt with air above it
	growGrass()

	// Create clouds
	generateClouds()

	// Place flowers and pebbles above grass
	generateNature()

	// Generate structure
	generateStructures()

	// Fix the orientation of blocks in the world
	orientBlocks("dirt", true)
	orientBlocks("grass", true)
	orientBlocks("stone", true)
	orientBlocks("leaves", true)

	// Fix backdirt
	createAllExtraBackdirt()
	orientBlocks("backdirt", true)

	// Light up all blocks
	CreateLighting(WorldWidth/2, HeightMap[WorldWidth/2]+5, 0.9)

	// Set player starting position
	Player.SetPosition(float32(WorldWidth*BlockSize/2), float32((HeightMap[WorldWidth/2]+25)*BlockSize))
}

//  --------------------------------------------------
//  World Generation Functions
//  --------------------------------------------------

func generateHeights() {
	for x := 0; x < WorldWidth; x++ {
		HeightMap[x] = GrassMinimum + int(Flatness*noise1D(float64(x)/(WorldWidth/2))*WorldHeight)
	}
	for x := 0; x < WorldWidth; x++ {
		createWorldBlock(x, HeightMap[x], "grass")
	}
}

func fillHeights() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight-1; y++ {
			createWorldBlock(x, y, "dirt")
			if WorldMap.GetWorldBlockName(x, y+1) == "grass" {
				break
			}
		}
	}
}

func fillStone() {
	Generator = perlin.NewPerlin(1.2, 2, 2, int64(rand.Int()))
	stoneFrequency := StoneStartingFrequency
	for y := 0; y < WorldHeight; y++ {
		for x := 0; x < WorldWidth; x++ {
			n := noise2D(StoneNoiseScalar*float64(x)/WorldWidth*2, StoneNoiseScalar*float64(y)/WorldHeight*4)
			if n > stoneFrequency {
				createWorldBlock(x, y, "stone")
			}
		}
		stoneFrequency += (1 / StoneTop)
	}
}

func cleanStone() {
	for x := 0; x < WorldWidth; x++ {
		grassHeight := HeightMap[x]
		if WorldMap.GetWorldBlockName(x, grassHeight) == "stone" {
			for y := grassHeight + StoneTopDeviation; y < WorldHeight; y++ {
				createWorldBlock(x, y, "sky")
			}
		} else {
			for y := grassHeight + 1; y < WorldHeight; y++ {
				createWorldBlock(x, y, "sky")
			}
		}
	}
}

func generateCaves() {
	Generator = perlin.NewPerlin(1.5, 2, 3, int64(rand.Int()))
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			n := noise2D(CaveNoiseScalar*float64(x)/WorldWidth*2, CaveNoiseScalar*float64(y)/WorldHeight*4)
			if n > CaveNoiseThreshold && y <= HeightMap[x] {
				WorldMap.RemoveWorldBlock(x, y)
				createBackBlock(x, y, "backdirt")
			}
		}
	}
}

func cleanBackDirt() {

}

func growGrass() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if WorldMap.GetWorldBlockName(x, y) == "dirt" && (WorldMap.GetWorldBlockName(x, y+1) == "sky" || WorldMap.GetBackBlockName(x, y+1) == "backdirt") {
				createWorldBlock(x, y, "grass")
			}
		}
	}
}

func generateClouds() {
	for x := 0; x < WorldWidth; x++ {
		if rand.Float32() < 0.4 {
			CloudChild.AddCopy(
				child.ChildCopy{
					X:        float32(x * BlockSize),
					Y:        float32((rand.Intn(20) + HeightMap[x] + 15) * BlockSize),
					Material: &cloudMaterial,
					Darkness: 1,
				},
			)
			x += 400 / BlockSize
		}
	}
}

func generateNature() {
	for x := 1; x < WorldWidth-1; x++ {
		if WorldMap.GetWorldBlockName(x, HeightMap[x]) == "grass" && WorldMap.GetWorldBlockName(x, HeightMap[x] + 1) == "sky" || WorldMap.GetWorldBlockName(x, HeightMap[x] + 1) == "backdirt" {
			natureRand := rand.Intn(16)
			if natureRand == 15 && WorldMap.GetWorldBlockName(x-1, HeightMap[x] + 2) != "treetrunk" {
				createNatureBlock(x, HeightMap[x] + 1, "treeBottomRoot")
				height := 4 + rand.Intn(8)
				for i := 0; i < height; i++ {
					if rand.Intn(4) == 0 && i < height-2 && i > 0 {
						createNatureBlock(x-1, HeightMap[x] + i + 2, "treeBranchL1")
					}
					if rand.Intn(4) == 0 && i < height-2 && i > 0 {
						createNatureBlock(x+1, HeightMap[x] + i + 2, "treeBranchR1")
					}
					createNatureBlock(x, HeightMap[x] + i + 2, "treeTrunk")
				}
				
				createNatureBlock(x-1, HeightMap[x] + height + 1, "leaves") // TL
				createNatureBlock(x, HeightMap[x] + height + 1, "leaves") // TM
				createNatureBlock(x+1, HeightMap[x] + height + 1, "leaves") // TR
				createNatureBlock(x-1, HeightMap[x] + height, "leaves")   // ML
				createNatureBlock(x, HeightMap[x] + height, "leaves")    // MM
				createNatureBlock(x+1, HeightMap[x] + height, "leaves")   // MR
				createNatureBlock(x-1, HeightMap[x] + height - 1, "leaves")  //BL
				createNatureBlock(x, HeightMap[x] + height - 1, "leaves")    // BM
				createNatureBlock(x+1, HeightMap[x] + height - 1, "leaves")  //BL
			} else if natureRand > 13 {
				floraRand := rand.Intn(4) + 1
				floraType := fmt.Sprintf("flower%d", floraRand)
				if floraRand != 4 {
					createNatureBlock(x, HeightMap[x] + 1, floraType)
				} else {
					createNatureBlock(x, HeightMap[x] + 1, "pebble")
				}
			} else if natureRand > 9 {
				grassRand := rand.Intn(3) + 1
				grassType := fmt.Sprintf("topGrass%d", grassRand)
				createNatureBlock(x, HeightMap[x] + 1, grassType)
			}
		}
	}

}

//  --------------------------------------------------
//  World Generation Helpers
//  --------------------------------------------------

func isBackBlock(name string) bool {
	for _, transparent := range TransparentBlocks {
		if name == transparent {
			return true
		}
	}
	return false
}

/*func blockType(name string) string {
	for _, green := range natureBlocks {
		if NameMap[name] == NameMap[green] {
			return "nature"
		}
	}
	return "shit spelling"
}*/

func noise2D(x, y float64) float64 {
	return (Generator.Noise2D(x, y) + 0.4) / 0.8
}

func noise1D(x float64) float64 {
	return (Generator.Noise1D(x) + 0.4) / 0.8
}

func randomizeSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
	Generator = perlin.NewPerlin(2, 2, 10, int64(rand.Int()))
}

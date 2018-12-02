package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"rapidengine/child"
	"strconv"
)

//  --------------------------------------------------
//  Storage.go contains the WorldTree, which stores the
//  entire world, and can save/load worlds.
//
//  ID System:
//  Each block in the map has a 5 digit ID, in which the
//  first 3 digits specify the Block ID (see blocks.go)
//  and the last 2 specify the orientation of the block.
//
//  This allows for simple serialization of the world
//	in and out of a file on the disk.
//  --------------------------------------------------

// WorldTree contains the entire world map
type WorldTree struct {
	blockNodes [WorldWidth][WorldHeight]BlockNode
}

// BlockNode contains all the data for one tile on the map
type BlockNode struct {
	worldBlock  *child.ChildCopy
	backBlock   *child.ChildCopy
	natureBlock *child.ChildCopy
	lightBlock  *child.ChildCopy
}

// NewWorldTree returns an empty WorldTree
func NewWorldTree() WorldTree {
	w := WorldTree{}
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			w.AddWorldBlock(x, y, &child.ChildCopy{
				ID:       "00000",
				Darkness: 0,
			})
			w.AddBackBlock(x, y, &child.ChildCopy{
				ID:       "00000",
				Darkness: 0,
			})
			w.AddNatureBlock(x, y, &child.ChildCopy{
				ID:       "00000",
				Darkness: 0,
			})
			w.AddLightBlock(x, y, &child.ChildCopy{
				ID:       "00000",
				Darkness: 0,
			})
		}
	}
	return w
}

//  --------------------------------------------------
//  Node Creation
//  --------------------------------------------------

func (tree *WorldTree) AddNode(x, y int, node BlockNode) {
	tree.blockNodes[x][y] = node
}

func (tree *WorldTree) AddWorldBlock(x, y int, cpy *child.ChildCopy) {
	tree.blockNodes[x][y].worldBlock = cpy
}

func (tree *WorldTree) AddBackBlock(x, y int, cpy *child.ChildCopy) {
	tree.blockNodes[x][y].backBlock = cpy
}

func (tree *WorldTree) AddNatureBlock(x, y int, cpy *child.ChildCopy) {
	tree.blockNodes[x][y].natureBlock = cpy
}

func (tree *WorldTree) AddLightBlock(x, y int, cpy *child.ChildCopy) {
	tree.blockNodes[x][y].lightBlock = cpy
}

func (tree *WorldTree) RemoveWorldBlock(x, y int) {
	tree.blockNodes[x][y].worldBlock = &child.ChildCopy{
		ID: "00000",
	}
}

func (tree *WorldTree) RemoveBackBlock(x, y int) {
	tree.blockNodes[x][y].backBlock = &child.ChildCopy{
		ID: "00000",
	}
}

func (tree *WorldTree) RemoveNatureBlock(x, y int) {
	tree.blockNodes[x][y].natureBlock = &child.ChildCopy{
		ID: "00000",
	}
}

//  --------------------------------------------------
//  Node Retrieval
//  --------------------------------------------------

func (tree *WorldTree) GetWorldBlock(x, y int) *child.ChildCopy {
	return tree.blockNodes[x][y].worldBlock
}

func (tree *WorldTree) GetBackBlock(x, y int) *child.ChildCopy {
	return tree.blockNodes[x][y].backBlock
}

func (tree *WorldTree) GetNatureBlock(x, y int) *child.ChildCopy {
	return tree.blockNodes[x][y].natureBlock
}

func (tree *WorldTree) GetLightBlock(x, y int) *child.ChildCopy {
	return tree.blockNodes[x][y].lightBlock
}

func (tree *WorldTree) GetWorldBlockName(x, y int) string {
	return GetNameFromID(tree.blockNodes[x][y].worldBlock.ID[:3])
}

func (tree *WorldTree) GetBackBlockName(x, y int) string {
	return GetNameFromID(tree.blockNodes[x][y].backBlock.ID[:3])
}

func (tree *WorldTree) GetNatureBlockName(x, y int) string {
	return GetNameFromID(tree.blockNodes[x][y].natureBlock.ID[:3])
}

func (tree *WorldTree) GetWorldBlockID(x, y int) string {
	return tree.blockNodes[x][y].worldBlock.ID
}

func (tree *WorldTree) GetBackBlockID(x, y int) string {
	return tree.blockNodes[x][y].backBlock.ID
}

func (tree *WorldTree) GetNatureBlockID(x, y int) string {
	return tree.blockNodes[x][y].natureBlock.ID
}

func (tree *WorldTree) GetLightBlockID(x, y int) string {
	return tree.blockNodes[x][y].lightBlock.ID
}

func (tree *WorldTree) GetWorldBlockOrientation(x, y int) string {
	return GetOrientationFromID(tree.blockNodes[x][y].worldBlock.ID[3:])
}

func (tree *WorldTree) GetBackBlockOrientation(x, y int) string {
	return GetOrientationFromID(tree.blockNodes[x][y].backBlock.ID[3:])
}

func (tree *WorldTree) GetDarkness(x, y int) float32 {
	if tree.blockNodes[x][y].worldBlock.ID == "00000" {
		if back := tree.blockNodes[x][y].backBlock; back.ID != "00000" {
			return back.Darkness
		}
	}
	return tree.blockNodes[x][y].worldBlock.Darkness
}

//  --------------------------------------------------
//  Node Modification
//  --------------------------------------------------

// Updates node materials, for when orientations change
func (tree *WorldTree) UpdateWorldBlockMaterial(x, y int) {
	tree.blockNodes[x][y].worldBlock.Material = GetBlock(tree.GetWorldBlockName(x, y)).GetMaterial(tree.GetWorldBlockOrientation(x, y))
}

func (tree *WorldTree) UpdateBackBlockMaterial(x, y int) {
	tree.blockNodes[x][y].backBlock.Material = GetBlock(tree.GetBackBlockName(x, y)).GetMaterial(tree.GetBackBlockOrientation(x, y))
}

func (tree *WorldTree) UpdateNatureBlockMaterial(x, y int) {
	tree.blockNodes[x][y].natureBlock.Material = GetBlock(tree.GetNatureBlockName(x, y)).GetMaterial(tree.GetWorldBlockOrientation(x, y))
}

func (tree *WorldTree) SetWorldBlockOrientation(x, y int, orient string) {
	tree.blockNodes[x][y].worldBlock.ID = tree.blockNodes[x][y].worldBlock.ID[:3] + OrientationsMap[orient]
}
func (tree *WorldTree) SetBackBlockOrientation(x, y int, orient string) {
	tree.blockNodes[x][y].backBlock.ID = tree.blockNodes[x][y].backBlock.ID[:3] + OrientationsMap[orient]
}

func (tree *WorldTree) SetDarkness(x, y int, darkness float32) {
	tree.blockNodes[x][y].worldBlock.Darkness = darkness
	tree.blockNodes[x][y].backBlock.Darkness = darkness
	tree.blockNodes[x][y].natureBlock.Darkness = darkness
	tree.blockNodes[x][y].lightBlock.Darkness = darkness
}

//  --------------------------------------------------
//  Block Helpers
//  --------------------------------------------------

func createWorldBlock(x, y int, name string) {
	WorldMap.AddWorldBlock(x, y, &child.ChildCopy{
		X:        float32(x * BlockSize),
		Y:        float32(y * BlockSize),
		Material: GetBlock(name).GetMaterial("NN"),
		Darkness: 0,
		ID:       GetIDFromName(name) + "00",
	})
}

func createBackBlock(x, y int, name string) {
	WorldMap.AddBackBlock(x, y, &child.ChildCopy{
		X:        float32(x * BlockSize),
		Y:        float32(y * BlockSize),
		Material: GetBlock(name).GetMaterial("NN"),
		Darkness: 0,
		ID:       GetIDFromName(name) + "00",
	})
}

func createNatureBlock(x, y int, name string) {
	WorldMap.AddNatureBlock(x, y, &child.ChildCopy{
		X:        float32(x * BlockSize),
		Y:        float32(y*BlockSize) - 5,
		Material: GetBlock(name).GetMaterial("NN"),
		Darkness: 0,
		ID:       GetIDFromName(name) + "00",
	})
}

func createLightBlock(x, y int, name string) {
	WorldMap.AddLightBlock(x, y, &child.ChildCopy{
		X:        float32(x * BlockSize),
		Y:        float32(y * BlockSize),
		Material: GetBlock(name).GetMaterial("NN"),
		Darkness: 0.8,
		ID:       GetIDFromName(name) + "00",
	})
}

//  --------------------------------------------------
//  World Serialization
//  --------------------------------------------------

func (tree *WorldTree) WriteToFile(path string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	point := float32(100) / float32(WorldHeight)
	SaveProgressBar.SetPercentage(0)

	for x := 0; x < WorldWidth; x++ {
		f.WriteString(fmt.Sprint(HeightMap[x]))
		f.WriteString("\n")
	}

	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			f.WriteString(
				tree.GetWorldBlockID(x, y) +
					tree.GetBackBlockID(x, y) +
					tree.GetNatureBlockID(x, y) +
					tree.GetLightBlockID(x, y) +
					fmt.Sprint(tree.GetDarkness(x, y)) + "\n",
			)
		}
		SaveProgressBar.IncrementPercentage(point)
		updateSaveScreen()
	}

	if err != nil {
		panic(err)
	}
}

func (tree *WorldTree) LoadFromFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cx := 0
	cy := 0

	point := 100.0 / float32(WorldHeight)
	ProgressBar.SetPercentage(0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ht, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			panic(err)
		}

		HeightMap[cx] = int(ht)

		cx++
		if cx >= WorldWidth {
			break
		}
	}

	cx = 0
	cy = 0

	for scanner.Scan() {
		block := scanner.Text()

		if block[3:5] == "" {
			return
		}

		if block[:5] != "00000" {
			WorldMap.AddWorldBlock(cx, cy, &child.ChildCopy{
				ID:       block[:5],
				X:        float32(cx * BlockSize),
				Y:        float32(cy * BlockSize),
				Material: GetBlock(GetNameFromID(block[:3])).GetMaterial(InverseOrientationMap[block[3:5]]),
			})
		}

		if block[5:10] != "00000" {
			WorldMap.AddBackBlock(cx, cy, &child.ChildCopy{
				ID:       block[5:10],
				X:        float32(cx * BlockSize),
				Y:        float32(cy * BlockSize),
				Material: GetBlock(GetNameFromID(block[5:8])).GetMaterial(InverseOrientationMap[block[8:10]]),
			})
		}

		if block[10:15] != "00000" {
			WorldMap.AddNatureBlock(cx, cy, &child.ChildCopy{
				ID:       block[10:15],
				X:        float32(cx * BlockSize),
				Y:        float32(cy * BlockSize),
				Material: GetBlock(GetNameFromID(block[10:13])).GetMaterial(InverseOrientationMap[block[13:15]]),
			})
		}

		if block[15:20] != "00000" {
			WorldMap.AddLightBlock(cx, cy, &child.ChildCopy{
				ID:       block[15:20],
				X:        float32(cx * BlockSize),
				Y:        float32(cy * BlockSize),
				Material: GetBlock(GetNameFromID(block[15:18])).GetMaterial(InverseOrientationMap[block[18:20]]),
			})
		}

		darkness, err := strconv.ParseFloat(block[20:], 32)
		if err != nil {
			panic(err)
		}
		WorldMap.SetDarkness(cx, cy, float32(darkness))

		cy++
		if cy >= WorldHeight {
			cy = 0
			cx++

			ProgressBar.IncrementPercentage(point)
			updateLoadingScreen()
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (tree *WorldTree) writeToImage() {
	img := image.NewRGBA(image.Rect(0, 0, len(tree.blockNodes), len(tree.blockNodes[0])))

	width := len(tree.blockNodes)
	height := len(tree.blockNodes[0])

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			b := tree.GetWorldBlockName(x, y)

			if b == "sky" {
				if back := tree.GetBackBlockName(x, y); back != "sky" {
					b = back
				}
				if nature := tree.GetNatureBlockName(x, y); nature != "sky" {
					b = nature
				}
			}

			block := GetBlock(b)

			img.Set(x, height-y, color.RGBA{
				uint8(block.SaveColor[0]),
				uint8(block.SaveColor[1]),
				uint8(block.SaveColor[2]), 255})
		}
	}

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

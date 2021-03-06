package main

import (
	"math/rand"
	// "fmt"
)

type Dungeon struct {
	rooms     []Room
	corridors []Corridor
}

type Point struct {
	x int
	y int
}

type Room struct {
	x int
	y int
	w int
	h int
}

type Corridor struct {
	points []Point
}

func generateAllDungeons() {
	// Number of dungeons to be generated
	numDungeons := 20

	// Dimensions of total dungeon
	dungeonWidth := 100
	dungeonHeight := 60

	// Max and min dimensions of a single room
	maxRoomWidth := 30
	maxRoomHeight := 10
	minRoomWidth := 15
	minRoomHeight := 6

	// Maximum number of rooms per dungeon
	maxNumRooms := 5

	for currentDungeon := 0; currentDungeon < numDungeons; currentDungeon++ {
		// Slice of Rooms
		//Engine.Logger.Info("Making a Dungeon")
		rooms := make([]Room, 0)

		//Slice of Corridors
		corridors := make([]Corridor, 0)

		startx := rand.Intn(WorldWidth - dungeonWidth - maxRoomWidth)
		starty := HeightMap[startx] - rand.Intn(HeightMap[startx]-dungeonHeight-maxRoomHeight)

		numRooms := 1 + rand.Intn(maxNumRooms)

		// for currRooms := 0; currRooms < numRooms; {
		// 	tempRoom := generateRoom(dungeonWidth, dungeonHeight, maxRoomWidth, maxRoomHeight, startx, starty)
		// 	for _, currentRoom := range rooms {
		// 		if len(rooms) == 0 {
		// 			currRooms += 1
		// 			rooms = append(rooms, tempRoom)
		// 			Engine.Logger.Info("Added 1 YEEEEEEEEEEEEEEEEEEEEEEEEEEEEEET")
		// 			break
		// 		} else if roomIntersects(tempRoom, currentRoom) {
		// 			tempRoom = generateRoom(dungeonWidth, dungeonHeight, maxRoomWidth, maxRoomHeight, startx, starty)
		// 			Engine.Logger.Info("rip")
		// 			break
		// 		}
		// 		currRooms += 1
		// 		rooms = append(rooms, tempRoom)
		// 		Engine.Logger.Info("Added 1 YEEEEEEEEEEEEEEEEEEEEEEEEEEEEEET")
		// 	}
		// }

		// Generates Rooms
		for r := 0; r < numRooms; r++ {
			tempRoom := generateRoom(dungeonWidth, dungeonHeight, maxRoomWidth, maxRoomHeight, minRoomWidth, minRoomHeight, startx, starty)
			//Engine.Logger.Info("Starting to make a room")
			//checks if intersecting with all rooms
			//intersecting := true
			for _, currentRoom := range rooms {
				if !roomIntersects(tempRoom, currentRoom) || len(rooms) == 0 {
					break
				} else {
					tempRoom = generateRoom(dungeonWidth, dungeonHeight, maxRoomWidth, maxRoomHeight, minRoomWidth, minRoomHeight, startx, starty)
					continue
				}
			}
			rooms = append(rooms, tempRoom)
			//Engine.Logger.Info("Made A Room")
			//Engine.Logger.Info(fmt.Sprintf("X: %d", tempRoom.x) + fmt.Sprintf(" Y: %d", tempRoom.y))
		}

		// Generates Corridors
		for c := 1; c < numRooms; c++ {
			corridors = append(corridors, generateCorridor(rooms[c-1], rooms[c]))
		}

		// Places dungeon in world
		generateDungeon(Dungeon{rooms, corridors})
	}
}

func generateDungeon(dungeon Dungeon) {
	// Places rooms
	for _, room := range dungeon.rooms {
		for x := room.x; x < room.x+room.w; x++ {
			for y := room.y; y < room.y+room.h; y++ {
				if x == room.x || x == room.x+room.w-1 || y == room.y || y == room.y+room.h-1 {
					WorldMap.RemoveWorldBlock(x, y)
					createWorldBlock(x, y, "stoneBrick")
				} else {
					WorldMap.RemoveWorldBlock(x, y)
					createWorldBlock(x, y, "backdirt")
				}
			}
		}
	}

	// Places corridors
	for _, corridor := range dungeon.corridors {

		// Makes sures looping from the left-most x to the right-most x
		startx := 0
		endx := 0
		if corridor.points[0].x < corridor.points[1].x {
			startx = corridor.points[0].x
			endx = corridor.points[1].x
		} else {
			startx = corridor.points[1].x
			endx = corridor.points[0].x
		}

		//Creates Horizontal Corridors
		for x := startx; x < endx; x++ {
			y := corridor.points[0].y
			WorldMap.RemoveWorldBlock(x, y-2)
			WorldMap.RemoveWorldBlock(x, y-1)
			WorldMap.RemoveWorldBlock(x, y)
			WorldMap.RemoveWorldBlock(x, y+1)
			WorldMap.RemoveWorldBlock(x, y+2)
			createWorldBlock(x, y-2, "stoneBrick")
			createWorldBlock(x, y-1, "backdirt")
			createWorldBlock(x, y, "backdirt")
			createWorldBlock(x, y+1, "backdirt")
			createWorldBlock(x, y+2, "stoneBrick")
		}

		// Makes sures looping from the lowest y to the highest y
		starty := 0
		endy := 0
		if corridor.points[1].y < corridor.points[2].y {
			starty = corridor.points[1].y
			endy = corridor.points[2].y
		} else {
			starty = corridor.points[2].y
			endy = corridor.points[1].y
		}

		// Creates Vertical Corridors
		for y := starty; y < endy; y++ {
			x := corridor.points[2].x
			WorldMap.RemoveWorldBlock(x-1, y)
			WorldMap.RemoveWorldBlock(x, y)
			WorldMap.RemoveWorldBlock(x+1, y)
			createWorldBlock(x-1, y, "stoneBrick")
			createWorldBlock(x, y, "backdirt")
			createWorldBlock(x+1, y, "stoneBrick")
		}
	}
}

func generateRoom(dungeonWidth, dungeonHeight, maxRoomWidth, maxRoomHeight, minRoomWidth, minRoomHeight, startx, starty int) Room {
	roomX := startx + rand.Intn(dungeonWidth)
	roomY := starty - rand.Intn(dungeonHeight)
	roomW := minRoomWidth + rand.Intn(maxRoomWidth-minRoomWidth+1)
	roomH := minRoomHeight + rand.Intn(maxRoomHeight-minRoomHeight+1)
	return Room{roomX, roomY, roomW, roomH}
}

func roomIntersects(r1, r2 Room) bool {
	if r1.x < r2.x+r2.w &&
		r1.x+r1.w > r2.x &&
		r1.y < r2.y+r2.h &&
		r1.y+r1.h > r2.y {
		return true
	}
	return false
}

func generateCorridor(r1, r2 Room) Corridor {
	center1x := r1.x + r1.w/2
	center1y := r1.y + r1.h/2
	center2x := r2.x + r2.w/2
	center2y := r2.y + r2.h/2
	points := make([]Point, 3)
	points[0] = Point{center1x, center1y}
	points[1] = Point{center2x, center1y}
	points[2] = Point{center2x, center2y}
	return Corridor{points}
}

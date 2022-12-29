package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"log"
)

type Square uint8

const (
	Void Square = iota
	Open
	Wall
)

func (s Square) ToChar() byte {
	switch s {
	case Open:
		return '.'
	case Wall:
		return '#'
	default:
		return ' '
	}
}

type Direction uint8

const (
	Right Direction = iota
	Down
	Left
	Up
)

func (d Direction) Turn(by Direction, times int) Direction {
	var add int
	if by == Left {
		add = -times
	} else if by == Right {
		add = times
	} else {
		log.Panic("Invalid direction to turn by")
	}
	return Direction(lib.Mod(int(d)+add, 4))
}

func (d Direction) ToVec() lib.Vector2 {
	switch d {
	case Right:
		return lib.Vector2{X: 1, Y: 0}
	case Down:
		return lib.Vector2{X: 0, Y: 1}
	case Left:
		return lib.Vector2{X: -1, Y: 0}
	default:
		return lib.Vector2{X: 0, Y: -1}
	}
}

func (d Direction) ToChar() byte {
	switch d {
	case Right:
		return '>'
	case Down:
		return 'v'
	case Left:
		return '<'
	default:
		return '^'
	}
}

type Operation struct {
	steps   int
	turnDir Direction
}

type Face uint8

const (
	FNone   Face = 0
	FTop    Face = 1
	FBack   Face = 2
	FLeft   Face = 3
	FFront  Face = 4
	FBottom Face = 5
	FRight  Face = 6
)

type Character struct {
	pos lib.Vector2
	dir Direction
}

type Cubemap struct {
	sideLen    int
	faces      [][]Face
	faceRot    map[Face]int
	facePosMap map[Face]lib.Vector2
}

type World struct {
	sqrs        [][]Square
	path        []Operation
	char        Character
	isCube      bool
	cube        Cubemap
	breadCrumbs map[lib.Vector2]Direction
}

func (w *World) Height() int {
	return len(w.sqrs)
}

func (w *World) Width() int {
	return len(w.sqrs[0])
}

func (w *World) Get(v lib.Vector2) Square {
	return w.sqrs[v.Y][v.X]
}

func (w *World) GetInitialPos() lib.Vector2 {
	for x := 0; x < w.Width(); x++ {
		test := lib.Vector2{X: x, Y: 0}
		if w.Get(test) == Open {
			return test
		}
	}
	log.Panic("Failed to find initial pos")
	return lib.Vector2{}
}

var faceMap = map[Face]map[Direction]Face{
	FTop: {
		Right: FRight,
		Down:  FFront,
		Left:  FLeft,
		Up:    FBack,
	},
	FFront: {
		Right: FRight,
		Down:  FBottom,
		Left:  FLeft,
		Up:    FTop,
	},
	FBottom: {
		Right: FRight,
		Down:  FBack,
		Left:  FLeft,
		Up:    FFront,
	},
	FBack: {
		Right: FLeft,
		Down:  FBottom,
		Left:  FRight,
		Up:    FTop,
	},
	FLeft: {
		Right: FFront,
		Down:  FBottom,
		Left:  FBack,
		Up:    FTop,
	},
	FRight: {
		Right: FBack,
		Down:  FBottom,
		Left:  FFront,
		Up:    FTop,
	},
}

func getDirFromFaceToFace(face1 Face, face2 Face) Direction {
	for dir, otherFace := range faceMap[face1] {
		if otherFace == face2 {
			return dir
		}
	}
	log.Panic("Face2 can't be reached from Face1")
	return 0
}

func (w *World) CubeStep(curFace Face, dir Direction) Face {
	dir = dir.Turn(Left, w.cube.faceRot[curFace])
	return faceMap[curFace][dir]
}

func (w *World) posToRelative(pos lib.Vector2) (rel lib.Vector2) {
	rel.X = pos.X % w.cube.sideLen
	rel.Y = pos.Y % w.cube.sideLen
	return
}

func (w *World) posToAbsolute(rel lib.Vector2, face Face) (abs lib.Vector2) {
	facePos := w.cube.facePosMap[face]
	abs = facePos.Add(rel)
	return
}

func (w *World) GetFace(v lib.Vector2) Face {
	fx := v.X / w.cube.sideLen
	fy := v.Y / w.cube.sideLen
	return w.cube.faces[fy][fx]
}

func (w *World) GetEdgeRelativePos(pos lib.Vector2, dir Direction) int {
	relPos := w.posToRelative(pos)
	switch dir {
	case Right:
		return relPos.Y
	case Down:
		return relPos.X
	case Left:
		return relPos.Y
	case Up:
		return relPos.X
	}
	log.Panic("Unknown direction")
	return 0
}

func (w *World) GetEdgeAbsolutePos(edgeRelPos int, face Face, dir Direction) lib.Vector2 {
	var rel lib.Vector2
	switch dir {
	case Right:
		rel.X = w.cube.sideLen - 1
		rel.Y = edgeRelPos
	case Down:
		rel.X = edgeRelPos
		rel.Y = w.cube.sideLen - 1
	case Left:
		rel.X = 0
		rel.Y = edgeRelPos
	case Up:
		rel.X = edgeRelPos
		rel.Y = 0
	}
	return w.posToAbsolute(rel, face)
}

var mapEdgeInvert = map[Direction]map[Direction]bool{
	Left: {
		Left:  true,
		Up:    false,
		Right: false,
		Down:  true,
	},
	Up: {
		Left:  false,
		Up:    true,
		Right: true,
		Down:  false,
	},
	Right: {
		Left:  false,
		Up:    true,
		Right: true,
		Down:  false,
	},
	Down: {
		Left:  true,
		Up:    false,
		Right: false,
		Down:  true,
	},
}

func (w *World) mapEdgePos(fromEdge Direction, toEdge Direction, edgeRelPos int) int {
	invert := mapEdgeInvert[fromEdge][toEdge]
	if invert {
		return (w.cube.sideLen - 1) - edgeRelPos
	} else {
		return edgeRelPos
	}
}

func (w *World) InterpretCube(sideLen int, faces [][]Face, rot map[Face]int) {
	w.isCube = true
	w.cube = Cubemap{
		sideLen: sideLen,
		faces:   faces,
		faceRot: rot,
	}

	facePosMap := make(map[Face]lib.Vector2)
	for y := 0; y < len(faces); y++ {
		for x := 0; x < len(faces[y]); x++ {
			if faces[y][x] != 0 {
				facePosMap[Face(faces[y][x])] = lib.Vector2{
					X: x * sideLen,
					Y: y * sideLen,
				}
			}
		}
	}
	w.cube.facePosMap = facePosMap
}

func (w *World) IsOOB(pos lib.Vector2) bool {
	return pos.X < 0 || pos.Y < 0 || pos.X >= w.Width() || pos.Y >= w.Height() || w.Get(pos) == Void
}

func (w *World) Step(pos lib.Vector2, dir Direction) (lib.Vector2, Direction) {
	// First see if this step is trivial and not out of bounds
	targSqr := pos.Add(dir.ToVec())
	if !w.IsOOB(targSqr) {
		return targSqr, dir
	}

	// If not, figure out the step destination
	if w.isCube {
		posAlongEdge := w.GetEdgeRelativePos(pos, dir)
		curFace := w.GetFace(pos)
		nextFace := w.CubeStep(curFace, dir)
		newDirOpposite := getDirFromFaceToFace(nextFace, curFace)
		newDirOpposite = newDirOpposite.Turn(Left, -w.cube.faceRot[nextFace])
		newPosAlongEdge := w.mapEdgePos(dir, newDirOpposite, posAlongEdge)
		pos = w.GetEdgeAbsolutePos(newPosAlongEdge, nextFace, newDirOpposite)
		dir = newDirOpposite.Turn(Left, 2)
	} else {
		oppositeDir := dir.Turn(Left, 2)
		for true {
			next := pos.Add(oppositeDir.ToVec())
			if w.IsOOB(next) {
				break
			} else {
				pos = next
			}
		}
	}
	return pos, dir
}

func (w *World) Walk(steps int) {
	for i := 0; i < steps; i++ {
		w.breadCrumbs[w.char.pos] = w.char.dir

		newPos, newDir := w.Step(w.char.pos, w.char.dir)
		if w.Get(newPos) == Wall {
			break
		}

		w.char.pos = newPos
		w.char.dir = newDir
	}
}

func (w *World) Traverse() {
	w.char.pos = w.GetInitialPos()
	w.char.dir = Right

	for _, op := range w.path {
		if op.steps == 0 {
			w.char.dir = w.char.dir.Turn(op.turnDir, 1)
		} else {
			w.Walk(op.steps)
			// w.Print()
		}
	}
}

func (w *World) GetPassword() int {
	return (w.char.pos.Y+1)*1000 + (w.char.pos.X+1)*4 + int(w.char.dir)
}

func (w *World) Print() {
	for y := 0; y < w.Height(); y++ {
		for x := 0; x < w.Width(); x++ {
			p := lib.Vector2{X: x, Y: y}
			dir, ok := w.breadCrumbs[p]
			if ok {
				fmt.Printf("%c", dir.ToChar())
			} else {
				fmt.Printf("%c", w.Get(p).ToChar())
			}
		}
		fmt.Println()
	}
}

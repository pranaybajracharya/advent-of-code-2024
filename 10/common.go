package main

type Pos struct {
	x     int
	y     int
	value int
}

var threshold = 1

func validatePos(current Pos, area [][]int, x int, y int) *Pos {
	if !(x >= 0 && x < len(area) && y >= 0 && y < len(area[0])) {
		return nil
	}

	if area[x][y]-current.value == threshold {
		return &Pos{x: x, y: y, value: area[x][y]}
	}

	return nil
}

func nextMove(area [][]int, current Pos) []Pos {
	nextLeft := validatePos(current, area, current.x-1, current.y)
	nextRight := validatePos(current, area, current.x+1, current.y)
	nextUp := validatePos(current, area, current.x, current.y-1)
	nextDown := validatePos(current, area, current.x, current.y+1)

	next := []Pos{}
	if nextLeft != nil {
		next = append(next, *nextLeft)
	}
	if nextRight != nil {
		next = append(next, *nextRight)
	}
	if nextUp != nil {
		next = append(next, *nextUp)
	}
	if nextDown != nil {
		next = append(next, *nextDown)
	}
	return next
}

func uniquePositions(positions []Pos) []Pos {
	uniqueMap := make(map[Pos]bool)
	uniqueList := []Pos{}

	for _, pos := range positions {
		if !uniqueMap[pos] {
			uniqueMap[pos] = true
			uniqueList = append(uniqueList, pos)
		}
	}

	return uniqueList
}

func findPaths(area [][]int, current Pos, dest []Pos) []Pos {
	next := nextMove(area, current)

	if len(next) == 0 {
		return dest
	}

	for _, n := range next {
		if n.value == 9 {
			dest = append(dest, n)
		} else {
			dest = findPaths(area, n, dest)
		}
	}

	return dest
}

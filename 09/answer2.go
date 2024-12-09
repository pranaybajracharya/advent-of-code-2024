package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func answer2() {
	data, err := os.ReadFile("input.txt")
	// data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := string(data)
	diskmap := strings.Split(input, "")

	var dm []int
	idx := 0
	for i, disk := range diskmap {
		count, err := strconv.Atoi(disk)
		if err != nil {
			fmt.Println(err)
			return
		}
		if i%2 == 0 {
			for i := 0; i < count; i++ {
				dm = append(dm, idx)
			}
			idx++
		} else {
			for i := 0; i < count; i++ {
				dm = append(dm, -1)
			}
		}
	}

	ptr := len(dm) - 1
	for {
		// Find free space each time from start
		startPtr := 0
		if startPtr > ptr || ptr < 1 {
			break
		}

		current := dm[ptr]

		// Store indexes of ID number to be moved
		var idxArray []int
		if current != -1 {
			idxArray = append(idxArray, ptr)
		}
		for {
			if ptr < 1 {
				break
			}
			next := dm[ptr-1]
			if next == current {
				idxArray = append(idxArray, ptr-1)
				ptr--
			} else {
				break
			}
		}

		// Store indexes of free space
		var freeSpace []int

		for {
			if startPtr > ptr {
				// If free space is unavailable for this ID number, move on to next one
				ptr--
				break
			}
			if dm[startPtr] == -1 {
				freeSpace = append(freeSpace, startPtr)
			} else {
				if len(freeSpace) > 0 {
					if len(freeSpace) >= len(idxArray) {
						// Move files to free space ahead
						for i := 0; i < len(idxArray); i++ {
							destination := freeSpace[i]
							source := idxArray[i]
							dm[destination] = current
							dm[source] = -1
						}
						ptr--
						break
					} else {
						// Clear array when this chunk does not have enough space
						freeSpace = freeSpace[:0]
					}
				}
			}
			startPtr++
		}
	}

	sum := 0

	for i, char := range dm {
		if char == -1 {
			continue
		}
		sum += char * i
	}

	fmt.Println("Answer 2:", sum)
}

package days

import (
	"strconv"
	"strings"
)

type Day9 struct {
	disk     []int
	blocksum []int
}

func (d *Day9) Init(input string) {
	d.blocksum = make([]int, 0, 1024)
	s := strings.TrimSpace(input)
	d.disk = make([]int, len(s))
	for i, v := range s {
		d.disk[i], _ = strconv.Atoi(string(v))
	}
}

func (d *Day9) compact() {
	disk := d.disk
	lastFileID := len(disk) / 2
	// is the last block free space?
	if (len(disk)-1)%2 > 0 {
		lastFileID--
		disk = disk[:len(disk)-1]
	}
	lastFileSize := disk[len(disk)-1]

	index := 0
	for i := range disk {
		if i > len(disk)-1 {
			break
		}
		numBlocks := disk[i]
		if i%2 == 0 {
			for range numBlocks {
				d.blocksum = append(d.blocksum, (i / 2))
				index++
			}
		} else if numBlocks > 0 {
			for range numBlocks {
				if lastFileSize == 0 {
					// remove space from the end
					disk = disk[: len(disk)-2 : len(disk)-2]
					lastFileSize = disk[len(disk)-1]
					lastFileID--
					if lastFileID <= i/2 {
						break
					}
				} else {
					disk[i] -= 1
				}
				d.blocksum = append(d.blocksum, lastFileID)

				lastFileSize--
				disk[len(disk)-1] = lastFileSize
				index++
			}
		}
	}
}

func (d *Day9) checksum() (result int) {
	for i, id := range d.blocksum {
		result += i * id
	}
	return result
}

func (d *Day9) PartOne() (result int) {
	d.compact()
	return d.checksum()
}

func (d *Day9) PartTwo() (result int) {
	return result
}

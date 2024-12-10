package days

import (
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Day9 struct {
	disk     []int
	blocksum []int
}

func (d *Day9) Init(input string) {
	d.blocksum = make([]int, 0, 102400)
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

	for i := range disk {
		if i > len(disk)-1 {
			break
		}
		numBlocks := disk[i]
		if i%2 == 0 {
			d.blocksum = append(d.blocksum, sliceOf(i/2, numBlocks)...)
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
			}
		}
	}
}

func (d *Day9) loadDisk() {
	index := 0
	for i := range d.disk {
		numBlocks := d.disk[i]
		index += numBlocks
		d.blocksum = d.blocksum[:index]
		if i%2 == 0 {
			copy(d.blocksum[index-numBlocks:index], sliceOf(i/2, numBlocks))
		}
	}
}

func (d *Day9) defrag() {
	for i := len(d.blocksum) - 1; i > 0; i-- {
		if d.blocksum[i] == 0 {
			continue
		}
		size := d.disk[d.blocksum[i]*2]
		if size == 0 {
			continue
		}
		oldIndex := i - (size - 1)
		newIndex := oldIndex

		available := 0
		for j, v := range d.blocksum[:oldIndex] {
			if j > d.disk[0]-1 && v == 0 {
				available++
			} else {
				available = 0
			}
			if available == size {
				newIndex = j - (size - 1)
				break
			}
		}
		if available < size {
			continue
		}

		d.disk[d.blocksum[i]*2] = 0
		copy(d.blocksum[newIndex:newIndex+size], d.blocksum[oldIndex:oldIndex+size])
		copy(d.blocksum[oldIndex:oldIndex+size], sliceOf(0, size))
		i = oldIndex
	}
}

// Slice of length n with repeating value v
func sliceOf[E constraints.Integer](v E, n int) []E {
	s := make([]E, n)
	for i := range s {
		s[i] = v
	}
	return s
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
	d.loadDisk()
	d.defrag()
	return d.checksum()
}

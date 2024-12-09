package day9

import (
	"strconv"
	"strings"
)

type BlockSequence struct {
	fileID         int
	sequenceSize   int
	freeSpaceRight int
}

func (b BlockSequence) String() string {
	str := ""
	fileID := strconv.Itoa(b.fileID)
	for i := 0; i < b.sequenceSize; i++ {
		str += fileID
	}
	for i := 0; i < b.freeSpaceRight; i++ {
		str += "."
	}
	return str
}

func Part1(input string) string {
	splat := strings.Split(input, "")
	blocks := make([]BlockSequence, 0)
	for i, char := range splat {
		if strings.TrimSpace(char) == "" {
			continue
		}
		size, err := strconv.Atoi(char)
		if i%2 == 0 {
			if err != nil {
				panic(err)
			}
			blocks = append(blocks, BlockSequence{
				fileID:         i / 2,
				sequenceSize:   size,
				freeSpaceRight: 0,
			})
		} else {
			blocks[len(blocks)-1].freeSpaceRight = size
		}
	}

	for Fragmented(blocks) {
		lastBlock := blocks[len(blocks)-1]
		blocks = blocks[:len(blocks)-1]
		blocks[len(blocks)-1].freeSpaceRight += lastBlock.sequenceSize
		blocks = InsertBlockSequence(blocks, lastBlock)
	}

	return strconv.Itoa(CheckSum(blocks))
}

func InsertBlockSequence(blocks []BlockSequence, block BlockSequence) []BlockSequence {
	if len(blocks) == 0 {
		return []BlockSequence{block}
	}
	newBlocks := make([]BlockSequence, 0)
	done := false
	inserted := 0
	for _, b := range blocks {
		if done {
			newBlocks = append(newBlocks, b)
			continue
		}

		if b.freeSpaceRight == 0 {
			newBlocks = append(newBlocks, b)
			continue
		}

		if b.freeSpaceRight-block.sequenceSize >= 0 {
			newBlocks = append(newBlocks, BlockSequence{
				fileID:         b.fileID,
				sequenceSize:   b.sequenceSize,
				freeSpaceRight: 0,
			})
			block.freeSpaceRight = b.freeSpaceRight - block.sequenceSize
			newBlocks = append(newBlocks, block)
			inserted = block.sequenceSize
			done = true
		}

		if b.freeSpaceRight-block.sequenceSize < 0 {
			newBlocks = append(newBlocks, BlockSequence{
				fileID:         b.fileID,
				sequenceSize:   b.sequenceSize,
				freeSpaceRight: 0,
			})
			newBlocks = append(newBlocks, BlockSequence{
				fileID:         block.fileID,
				sequenceSize:   b.freeSpaceRight,
				freeSpaceRight: 0,
			})
			inserted = b.freeSpaceRight
			done = true
		}
	}

	if inserted != block.sequenceSize {
		newBlocks = append(newBlocks, BlockSequence{
			fileID:         block.fileID,
			sequenceSize:   block.sequenceSize - inserted,
			freeSpaceRight: block.freeSpaceRight,
		})
	}
	return newBlocks
}

func CheckSum(blocks []BlockSequence) int {
	sum := 0
	positionInFS := 0
	for _, b := range blocks {

		for j := positionInFS; j < positionInFS+b.sequenceSize; j++ {
			sum += j * b.fileID
		}

		positionInFS += b.sequenceSize + b.freeSpaceRight
	}
	return sum
}

func Fragmented(blocks []BlockSequence) bool {
	for i, b := range blocks {
		if i == len(blocks)-1 {
			return false
		}

		if b.freeSpaceRight != 0 {
			return true
		}
	}

	return false
}

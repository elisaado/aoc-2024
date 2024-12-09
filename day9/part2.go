package day9

import (
	"strconv"
	"strings"
)

func Part2(input string) string {
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

	// highest file ID is always the last one
	highestFileID := blocks[len(blocks)-1].fileID
	for i := highestFileID; i >= 0; i-- {
		// for _, b := range blocks {
		// 	print(b.String())
		// }
		// println()
		blocks = moveFile(blocks, i)
	}

	// for _, b := range blocks {
	// 	print(b.String())
	// }
	// println()

	return strconv.Itoa(CheckSum(blocks))
}

func moveFile(blocks []BlockSequence, fileID int) []BlockSequence {
	// find the block sequence with the given file ID
	var block BlockSequence
	for _, b := range blocks {
		if b.fileID == fileID {
			block = b
			break
		}
	}

	blockWithEnoughSpace := BlockSequence{
		fileID: -1,
	}
	for _, b := range blocks {
		if b.fileID == block.fileID {
			break
		}
		if b.freeSpaceRight >= block.sequenceSize {
			blockWithEnoughSpace = b
			break
		}
	}

	// can't move the file
	if blockWithEnoughSpace.fileID == -1 {
		return blocks
	}

	// remove the file to be moved from the list
	newBlocks := make([]BlockSequence, 0)
	for i, b := range blocks {
		if b.fileID != block.fileID {
			newBlocks = append(newBlocks, b)
		} else {
			if i > 0 {
				newBlocks[len(newBlocks)-1].freeSpaceRight += block.sequenceSize + block.freeSpaceRight
			}
		}
	}

	// insert the file to be moved
	// man i wish there was a better way to do this /s
	newNewBlocks := make([]BlockSequence, 0)
	for _, b := range newBlocks {
		if b.fileID == blockWithEnoughSpace.fileID {
			newNewBlocks = append(newNewBlocks, BlockSequence{
				fileID:         b.fileID,
				sequenceSize:   b.sequenceSize,
				freeSpaceRight: 0,
			})
			block.freeSpaceRight = b.freeSpaceRight - block.sequenceSize
			newNewBlocks = append(newNewBlocks, block)
		} else {
			newNewBlocks = append(newNewBlocks, b)
		}
	}

	return newNewBlocks

}

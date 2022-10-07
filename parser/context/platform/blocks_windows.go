//go:build windows
// +build windows

package platform

import (
	"io/fs"
)

type BlockSize = int64
type Blocks = int64

func FileBlocks(file fs.FileInfo) (BlockSize, Blocks) {
	return -1, -1
}

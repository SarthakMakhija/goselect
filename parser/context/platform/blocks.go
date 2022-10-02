//go:build !windows
// +build !windows

package platform

import (
	"io/fs"
	"syscall"
)

type BlockSize = int64
type Blocks = int64

func FileBlocks(file fs.FileInfo) (BlockSize, Blocks) {
	stat := file.Sys().(*syscall.Stat_t)
	return int64(stat.Blksize), stat.Blocks
}

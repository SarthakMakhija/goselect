//go:build windows
// +build windows

package platform

import (
	"io/fs"
	"syscall"
	"time"
)

type CreatedTime = time.Time
type ModifiedTime = time.Time
type AccessTime = time.Time

func FileTimes(file fs.FileInfo) (CreatedTime, ModifiedTime, AccessTime) {
	toTime := func(ft syscall.Filetime) time.Time {
		return time.Unix(0, ft.Nanoseconds())
	}

	stat := file.Sys().(*syscall.Win32FileAttributeData)
	return toTime(stat.CreationTime), toTime(stat.LastWriteTime), toTime(stat.LastAccessTime)
}

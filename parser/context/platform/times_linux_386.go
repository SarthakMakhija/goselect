//go:build 386 && linux
// +build 386,linux

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
	toTime := func(ts syscall.Timespec) time.Time {
		return time.Unix(int64(ts.Sec), int64(ts.Nsec))
	}
	stat := file.Sys().(*syscall.Stat_t)
	return toTime(stat.Ctim), toTime(stat.Mtim), toTime(stat.Atim)
}

//go:build linux
// +build linux

package platform

import (
	"io/fs"
	"syscall"
	"time"
)

func FileTimes(file fs.FileInfo) (CreatedTime, ModifiedTime, AccessTime) {
	toTime := func(ts syscall.Timespec) time.Time {
		return time.Unix(ts.Sec, ts.Nsec)
	}
	stat := file.Sys().(*syscall.Stat_t)
	return toTime(stat.Ctim), toTime(stat.Mtim), toTime(stat.Atim)
}

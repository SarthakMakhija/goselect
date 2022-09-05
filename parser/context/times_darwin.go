//go:build !linux
// +build !linux

package context

import (
	"io/fs"
	"syscall"
	"time"
)

type createdTime = time.Time
type modifiedTime = time.Time
type accessTime = time.Time

func fileTimes(file fs.FileInfo) (createdTime, modifiedTime, accessTime) {
	toTime := func(ts syscall.Timespec) time.Time {
		return time.Unix(ts.Sec, ts.Nsec)
	}
	stat := file.Sys().(*syscall.Stat_t)
	return toTime(stat.Ctimespec), toTime(stat.Mtimespec), toTime(stat.Atimespec)
}

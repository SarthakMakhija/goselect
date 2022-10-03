//go:build !windows
// +build !windows

package platform

import (
	"io/fs"
	"os/user"
	"strconv"
	"syscall"
)

type UserId = string
type UserName = string
type GroupId = string
type GroupName = string

func UserGroup(file fs.FileInfo) (UserId, UserName, GroupId, GroupName) {
	stat := file.Sys().(*syscall.Stat_t)
	userId := strconv.FormatUint(uint64(stat.Uid), 10)

	lookedUpUser, err := user.LookupId(userId)
	if err != nil {
		return "", "", "", ""
	}
	group, err := user.LookupGroupId(lookedUpUser.Gid)
	if err != nil {
		return "", "", "", ""
	}
	return lookedUpUser.Uid, lookedUpUser.Username, group.Gid, group.Name
}

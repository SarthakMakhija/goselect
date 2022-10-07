//go:build windows
// +build windows

package platform

import (
	"io/fs"
)

type UserId = string
type UserName = string
type GroupId = string
type GroupName = string

func UserGroup(file fs.FileInfo) (UserId, UserName, GroupId, GroupName) {
	return "", "", "", ""
}

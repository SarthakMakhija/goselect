package projection

import (
	"io/fs"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

type FileAttributes struct {
	attributes map[string]interface{}
}

func FromFile(file fs.FileInfo) *FileAttributes {
	fileAttributes := newFileAttributes()

	fileAttributes.SetName(file.Name())
	fileAttributes.SetExtension(filepath.Ext(file.Name()))
	fileAttributes.SetSize(file.Size())
	fileAttributes.SetPermission(file.Mode().Perm().String())
	fileAttributes.SetUserGroup(file)

	return fileAttributes
}

func newFileAttributes() *FileAttributes {
	return &FileAttributes{attributes: make(map[string]interface{})}
}

func (fileAttributes *FileAttributes) SetName(name string) {
	fileAttributes.attributes["name"] = name
	fileAttributes.attributes["fName"] = name
	fileAttributes.attributes["fname"] = name
	fileAttributes.attributes["Name"] = name
}

func (fileAttributes *FileAttributes) SetSize(size int64) {
	fileAttributes.attributes["size"] = size
	fileAttributes.attributes["fSize"] = size
	fileAttributes.attributes["fsize"] = size
	fileAttributes.attributes["Size"] = size
}

func (fileAttributes *FileAttributes) SetModifiedTime(time time.Time) {
	fileAttributes.attributes["modifiedTime"] = time
	fileAttributes.attributes["mTime"] = time
	fileAttributes.attributes["mtime"] = time
}

func (fileAttributes *FileAttributes) SetExtension(extension string) {
	fileAttributes.attributes["extension"] = extension
	fileAttributes.attributes["ext"] = extension
}

func (fileAttributes *FileAttributes) SetPermission(permission string) {
	fileAttributes.attributes["perm"] = permission
	fileAttributes.attributes["permission"] = permission
}

func (fileAttributes *FileAttributes) SetUserGroup(file fs.FileInfo) {
	stat := file.Sys().(*syscall.Stat_t)
	userId := strconv.FormatUint(uint64(stat.Uid), 10)

	lookedUpUser, err := user.LookupId(userId)
	if err != nil {
		fileAttributes.SetBlankUserGroup()
		return
	}

	group, err := user.LookupGroupId(lookedUpUser.Gid)
	if err != nil {
		fileAttributes.SetBlankUserGroup()
		return
	}
	fileAttributes.SetUserId(lookedUpUser.Uid)
	fileAttributes.SetUserName(lookedUpUser.Username)
	fileAttributes.SetGroupId(lookedUpUser.Gid)
	fileAttributes.SetGroupName(group.Name)
}

func (fileAttributes *FileAttributes) SetBlankUserGroup() {
	fileAttributes.SetUserId("")
	fileAttributes.SetUserName("")
	fileAttributes.SetGroupId("")
	fileAttributes.SetGroupName("")
}

func (fileAttributes *FileAttributes) SetUserId(userId string) {
	fileAttributes.attributes["userid"] = userId
	fileAttributes.attributes["userId"] = userId
	fileAttributes.attributes["uId"] = userId
	fileAttributes.attributes["uid"] = userId
}

func (fileAttributes *FileAttributes) SetUserName(userName string) {
	fileAttributes.attributes["userName"] = userName
	fileAttributes.attributes["username"] = userName
	fileAttributes.attributes["uName"] = userName
	fileAttributes.attributes["uname"] = userName
}

func (fileAttributes *FileAttributes) SetGroupId(groupId string) {
	fileAttributes.attributes["groupid"] = groupId
	fileAttributes.attributes["groupId"] = groupId
	fileAttributes.attributes["gId"] = groupId
	fileAttributes.attributes["gid"] = groupId
}

func (fileAttributes *FileAttributes) SetGroupName(groupName string) {
	fileAttributes.attributes["groupName"] = groupName
	fileAttributes.attributes["groupname"] = groupName
	fileAttributes.attributes["gName"] = groupName
	fileAttributes.attributes["gname"] = groupName
}

func (fileAttributes FileAttributes) Get(attribute string) interface{} {
	return fileAttributes.attributes[attribute]
}

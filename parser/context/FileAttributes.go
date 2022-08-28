package context

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

func ToFileAttributes(file fs.FileInfo, ctx *Context) *FileAttributes {
	fileAttributes := newFileAttributes()

	fileAttributes.setName(file.Name(), ctx.allAttributes)
	fileAttributes.setExtension(filepath.Ext(file.Name()), ctx.allAttributes)
	fileAttributes.setSize(file.Size(), ctx.allAttributes)
	fileAttributes.setPermission(file.Mode().Perm().String(), ctx.allAttributes)
	fileAttributes.setUserGroup(file, ctx.allAttributes)

	return fileAttributes
}

func newFileAttributes() *FileAttributes {
	return &FileAttributes{attributes: make(map[string]interface{})}
}

func (fileAttributes *FileAttributes) setName(name string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeName, name, attributes)
}

func (fileAttributes *FileAttributes) setSize(size int64, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeSize, size, attributes)
}

func (fileAttributes *FileAttributes) setModifiedTime(time time.Time, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeModifiedTime, time, attributes)
}

func (fileAttributes *FileAttributes) setExtension(extension string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeExtension, extension, attributes)
}

func (fileAttributes *FileAttributes) setPermission(permission string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributePermission, permission, attributes)
}

func (fileAttributes *FileAttributes) setUserGroup(file fs.FileInfo, attributes *AllAttributes) {
	stat := file.Sys().(*syscall.Stat_t)
	userId := strconv.FormatUint(uint64(stat.Uid), 10)

	lookedUpUser, err := user.LookupId(userId)
	if err != nil {
		fileAttributes.setBlankUserGroup(attributes)
		return
	}

	group, err := user.LookupGroupId(lookedUpUser.Gid)
	if err != nil {
		fileAttributes.setBlankUserGroup(attributes)
		return
	}
	fileAttributes.setUserId(lookedUpUser.Uid, attributes)
	fileAttributes.setUserName(lookedUpUser.Username, attributes)
	fileAttributes.setGroupId(lookedUpUser.Gid, attributes)
	fileAttributes.setGroupName(group.Name, attributes)
}

func (fileAttributes *FileAttributes) setBlankUserGroup(attributes *AllAttributes) {
	fileAttributes.setUserId("", attributes)
	fileAttributes.setUserName("", attributes)
	fileAttributes.setGroupId("", attributes)
	fileAttributes.setGroupName("", attributes)
}

func (fileAttributes *FileAttributes) setUserId(userId string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeUserId, userId, attributes)
}

func (fileAttributes *FileAttributes) setUserName(userName string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeUserName, userName, attributes)
}

func (fileAttributes *FileAttributes) setGroupId(groupId string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeGroupId, groupId, attributes)
}

func (fileAttributes *FileAttributes) setGroupName(groupName string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeGroupName, groupName, attributes)
}

func (fileAttributes *FileAttributes) setAllAliasesForAttribute(attribute string, value interface{}, attributes *AllAttributes) {
	for _, alias := range attributes.aliasesFor(attribute) {
		fileAttributes.attributes[alias] = value
	}
}

func (fileAttributes FileAttributes) Get(attribute string) interface{} {
	return fileAttributes.attributes[attribute]
}

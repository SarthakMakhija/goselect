package context

import (
	"github.com/dustin/go-humanize"
	"io/fs"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type FileAttributes struct {
	attributes map[string]Value
}

func ToFileAttributes(directory string, file fs.FileInfo, ctx *ParsingApplicationContext) *FileAttributes {
	fileAttributes := newFileAttributes()

	fileName := file.Name()
	extension := filepath.Ext(fileName)

	fileAttributes.setName(fileName, ctx.allAttributes)
	fileAttributes.setExtension(extension, ctx.allAttributes)
	fileAttributes.setBaseName(strings.Replace(fileName, extension, "", 1), ctx.allAttributes)
	fileAttributes.setSize(file.Size(), ctx.allAttributes)
	fileAttributes.setFormattedSize(file.Size(), ctx.allAttributes)
	fileAttributes.setFileType(file, ctx.allAttributes)
	path := directory + "/" + fileName
	absolutePath, err := filepath.Abs(path)
	if err == nil {
		fileAttributes.setAbsolutePath(absolutePath, ctx.allAttributes)
	}
	fileAttributes.setPath(path, ctx.allAttributes)
	fileAttributes.setPermission(file, ctx.allAttributes)
	fileAttributes.setUserGroup(file, ctx.allAttributes)

	return fileAttributes
}

func newFileAttributes() *FileAttributes {
	return &FileAttributes{attributes: make(map[string]Value)}
}

func (fileAttributes *FileAttributes) setName(name string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeName, StringValue(name), attributes)
}

func (fileAttributes *FileAttributes) setBaseName(name string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeBaseName, StringValue(name), attributes)
}

func (fileAttributes *FileAttributes) setSize(size int64, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeSize, Int64Value(size), attributes)
}

func (fileAttributes *FileAttributes) setFormattedSize(size int64, attributes *AllAttributes) {
	formattedSize := humanize.Bytes(uint64(size))
	fileAttributes.setAllAliasesForAttribute(AttributeFormattedSize, StringValue(formattedSize), attributes)
}

func (fileAttributes *FileAttributes) setFileType(file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeNameIsDir, booleanValueUsing(file.IsDir()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeNameIsFile, booleanValueUsing(file.Mode().IsRegular()), attributes)

	hiddenFile, _ := isHiddenFile(file.Name())
	fileAttributes.setAllAliasesForAttribute(AttributeNameIsHidden, booleanValueUsing(hiddenFile), attributes)
}

func (fileAttributes *FileAttributes) setPath(path string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributePath, StringValue(path), attributes)
}

func (fileAttributes *FileAttributes) setAbsolutePath(path string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeAbsolutePath, StringValue(path), attributes)
}

func (fileAttributes *FileAttributes) setModifiedTime(time time.Time, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeModifiedTime, DateTimeValue(time), attributes)
}

func (fileAttributes *FileAttributes) setExtension(extension string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeExtension, StringValue(extension), attributes)
}

func (fileAttributes *FileAttributes) setPermission(file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributePermission, StringValue(file.Mode().Perm().String()), attributes)

	perm := filePermission(file.Mode().Perm())
	fileAttributes.setAllAliasesForAttribute(AttributeUserRead, booleanValueUsing(perm.userRead()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeUserWrite, booleanValueUsing(perm.userWrite()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeUserExecute, booleanValueUsing(perm.userExecute()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeGroupRead, booleanValueUsing(perm.groupRead()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeGroupWrite, booleanValueUsing(perm.groupWrite()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeGroupExecute, booleanValueUsing(perm.groupExecute()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeOthersRead, booleanValueUsing(perm.othersRead()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeOthersWrite, booleanValueUsing(perm.othersWrite()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeOthersExecute, booleanValueUsing(perm.othersExecute()), attributes)
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
	fileAttributes.setAllAliasesForAttribute(AttributeUserId, StringValue(userId), attributes)
}

func (fileAttributes *FileAttributes) setUserName(userName string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeUserName, StringValue(userName), attributes)
}

func (fileAttributes *FileAttributes) setGroupId(groupId string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeGroupId, StringValue(groupId), attributes)
}

func (fileAttributes *FileAttributes) setGroupName(groupName string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeGroupName, StringValue(groupName), attributes)
}

func (fileAttributes *FileAttributes) setAllAliasesForAttribute(
	attribute string,
	value Value,
	attributes *AllAttributes,
) {
	for _, alias := range attributes.aliasesFor(attribute) {
		fileAttributes.attributes[alias] = value
	}
}

func (fileAttributes *FileAttributes) Get(attribute string) Value {
	v, ok := fileAttributes.attributes[strings.ToLower(attribute)]
	if ok {
		return v
	}
	return EmptyValue
}

type filePermission uint32

const (
	userRead      filePermission = 1 << 8
	userWrite                    = 1 << 7
	userExecute                  = 1 << 6
	groupRead                    = 1 << 5
	groupWrite                   = 1 << 4
	groupExecute                 = 1 << 3
	othersRead                   = 1 << 2
	othersWrite                  = 1 << 1
	othersExecute                = 1 << 0
)

func (p filePermission) userRead() bool {
	return p&userRead != 0
}

func (p filePermission) userWrite() bool {
	return p&userWrite != 0
}

func (p filePermission) userExecute() bool {
	return p&userExecute != 0
}

func (p filePermission) groupRead() bool {
	return p&groupRead != 0
}

func (p filePermission) groupWrite() bool {
	return p&groupWrite != 0
}

func (p filePermission) groupExecute() bool {
	return p&groupExecute != 0
}

func (p filePermission) othersRead() bool {
	return p&othersRead != 0
}

func (p filePermission) othersWrite() bool {
	return p&othersWrite != 0
}

func (p filePermission) othersExecute() bool {
	return p&othersExecute != 0
}

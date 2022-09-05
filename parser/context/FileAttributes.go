package context

import (
	"github.com/dustin/go-humanize"
	"goselect/parser/context/platform"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

type FileAttributes struct {
	attributes map[string]Value
}

func ToFileAttributes(directory string, file fs.FileInfo, ctx *ParsingApplicationContext) *FileAttributes {
	fileAttributes := newFileAttributes()

	fileAttributes.setName(file, ctx.allAttributes)
	fileAttributes.setExtension(file, ctx.allAttributes)
	fileAttributes.setSize(file, ctx.allAttributes)
	fileAttributes.setFileType(file, ctx.allAttributes)
	fileAttributes.setTimes(file, ctx.allAttributes)
	fileAttributes.setPath(directory, file, ctx.allAttributes)
	fileAttributes.setPermission(file, ctx.allAttributes)
	fileAttributes.setBlock(file, ctx.allAttributes)
	fileAttributes.setUserGroup(file, ctx.allAttributes)

	return fileAttributes
}

func newFileAttributes() *FileAttributes {
	return &FileAttributes{attributes: make(map[string]Value)}
}

func (fileAttributes *FileAttributes) setName(file fs.FileInfo, attributes *AllAttributes) {
	baseName := strings.Replace(file.Name(), filepath.Ext(file.Name()), "", 1)
	fileAttributes.setAllAliasesForAttribute(AttributeName, StringValue(file.Name()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeBaseName, StringValue(baseName), attributes)
}

func (fileAttributes *FileAttributes) setSize(file fs.FileInfo, attributes *AllAttributes) {
	formattedSize := humanize.Bytes(uint64(file.Size()))
	fileAttributes.setAllAliasesForAttribute(AttributeSize, Int64Value(file.Size()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeFormattedSize, StringValue(formattedSize), attributes)
}

func (fileAttributes *FileAttributes) setFileType(file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeNameIsDir, booleanValueUsing(file.IsDir()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeNameIsFile, booleanValueUsing(file.Mode().IsRegular()), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeNameIsSymbolicLink, booleanValueUsing(file.Mode()&os.ModeSymlink == os.ModeSymlink), attributes)
	if file.Mode().IsDir() {
		files, _ := ioutil.ReadDir(file.Name())
		fileAttributes.setAllAliasesForAttribute(AttributeNameIsEmpty, booleanValueUsing(len(files) == 0), attributes)
	} else {
		fileAttributes.setAllAliasesForAttribute(AttributeNameIsEmpty, booleanValueUsing(file.Size() == 0), attributes)
	}
	hiddenFile, _ := platform.IsHiddenFile(file.Name())
	fileAttributes.setAllAliasesForAttribute(AttributeNameIsHidden, booleanValueUsing(hiddenFile), attributes)
}

func (fileAttributes *FileAttributes) setTimes(file fs.FileInfo, attributes *AllAttributes) {
	created, modified, accessed := platform.FileTimes(file)
	fileAttributes.setAllAliasesForAttribute(AttributeCreatedTime, DateTimeValue(created), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeModifiedTime, DateTimeValue(modified), attributes)
	fileAttributes.setAllAliasesForAttribute(AttributeAccessedTime, DateTimeValue(accessed), attributes)
}

func (fileAttributes *FileAttributes) setPath(directory string, file fs.FileInfo, attributes *AllAttributes) {
	path := directory + "/" + file.Name()
	absolutePath, err := filepath.Abs(path)
	if err == nil {
		fileAttributes.setAllAliasesForAttribute(AttributeAbsolutePath, StringValue(absolutePath), attributes)
	}
	fileAttributes.setAllAliasesForAttribute(AttributePath, StringValue(path), attributes)
}

func (fileAttributes *FileAttributes) setExtension(file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForAttribute(AttributeExtension, StringValue(filepath.Ext(file.Name())), attributes)
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

func (fileAttributes *FileAttributes) setBlock(file fs.FileInfo, attributes *AllAttributes) {
	stat := file.Sys().(*syscall.Stat_t)
	if stat != nil {
		fileAttributes.setAllAliasesForAttribute(AttributeBlockSize, Int64Value(int64(stat.Blksize)), attributes)
		fileAttributes.setAllAliasesForAttribute(AttributeBlocks, Int64Value(stat.Blocks), attributes)
	} else {
		fileAttributes.setAllAliasesForAttribute(AttributeBlockSize, StringValue("NA"), attributes)
		fileAttributes.setAllAliasesForAttribute(AttributeBlocks, StringValue("NA"), attributes)
	}
}

func (fileAttributes *FileAttributes) setUserGroup(file fs.FileInfo, attributes *AllAttributes) {
	stat := file.Sys().(*syscall.Stat_t)
	if stat != nil {
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

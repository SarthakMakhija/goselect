package context

import (
	"github.com/dustin/go-humanize"
	"goselect/parser/context/platform"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

type EvaluatingValue struct {
	value           Value
	filePath        string
	isEvaluated     bool
	evaluationBlock AttributeLazyEvaluationBlock
}

type FileAttributes struct {
	attributes map[string]EvaluatingValue
}

func ToFileAttributes(directory string, file fs.FileInfo, ctx *ParsingApplicationContext) *FileAttributes {
	fileAttributes := newFileAttributes()

	fileAttributes.setName(file, ctx.allAttributes)
	fileAttributes.setExtension(file, ctx.allAttributes)
	fileAttributes.setSize(file, ctx.allAttributes)
	fileAttributes.setFileType(directory, file, ctx.allAttributes)
	fileAttributes.setTimes(file, ctx.allAttributes)
	fileAttributes.setPath(directory, file, ctx.allAttributes)
	fileAttributes.setPermission(file, ctx.allAttributes)
	fileAttributes.setBlock(file, ctx.allAttributes)
	fileAttributes.setUserGroup(file, ctx.allAttributes)
	fileAttributes.setMimeType(directory, file, ctx.allAttributes)

	return fileAttributes
}

func (fileAttributes *FileAttributes) Get(attribute string) Value {
	evaluatingValue, ok := fileAttributes.attributes[strings.ToLower(attribute)]
	if ok {
		if evaluatingValue.isEvaluated {
			return evaluatingValue.value
		}
		if value, err := evaluatingValue.evaluationBlock.evaluate(evaluatingValue.filePath); err != nil {
			return EmptyValue
		} else {
			return value
		}
	}
	return EmptyValue
}

func newFileAttributes() *FileAttributes {
	return &FileAttributes{attributes: make(map[string]EvaluatingValue)}
}

func (fileAttributes *FileAttributes) setName(file fs.FileInfo, attributes *AllAttributes) {
	baseName := strings.Replace(file.Name(), filepath.Ext(file.Name()), "", 1)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeName, StringValue(file.Name()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeBaseName, StringValue(baseName), attributes)
}

func (fileAttributes *FileAttributes) setSize(file fs.FileInfo, attributes *AllAttributes) {
	formattedSize := humanize.Bytes(uint64(file.Size()))
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeSize, Int64Value(file.Size()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeFormattedSize, StringValue(formattedSize), attributes)
}

func (fileAttributes *FileAttributes) setFileType(directory string, file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeNameIsDir, booleanValueUsing(file.IsDir()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeNameIsFile, booleanValueUsing(file.Mode().IsRegular()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeNameIsSymbolicLink, booleanValueUsing(file.Mode()&os.ModeSymlink == os.ModeSymlink), attributes)
	if file.Mode().IsDir() {
		newPath := fileAttributes.filePath(directory, file)
		entries, _ := os.ReadDir(newPath)
		fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeNameIsEmpty, booleanValueUsing(len(entries) == 0), attributes)
	} else {
		fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeNameIsEmpty, booleanValueUsing(file.Size() == 0), attributes)
	}
	hiddenFile, _ := platform.IsHiddenFile(file.Name())
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeNameIsHidden, booleanValueUsing(hiddenFile), attributes)
}

func (fileAttributes *FileAttributes) setTimes(file fs.FileInfo, attributes *AllAttributes) {
	created, modified, accessed := platform.FileTimes(file)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeCreatedTime, DateTimeValue(created), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeModifiedTime, DateTimeValue(modified), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeAccessedTime, DateTimeValue(accessed), attributes)
}

func (fileAttributes *FileAttributes) setPath(directory string, file fs.FileInfo, attributes *AllAttributes) {
	newPath := fileAttributes.filePath(directory, file)
	absolutePath, err := filepath.Abs(newPath)
	if err == nil {
		fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeAbsolutePath, StringValue(absolutePath), attributes)
	}
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributePath, StringValue(newPath), attributes)
}

func (fileAttributes *FileAttributes) setExtension(file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeExtension, StringValue(filepath.Ext(file.Name())), attributes)
}

func (fileAttributes *FileAttributes) setPermission(file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributePermission, StringValue(file.Mode().Perm().String()), attributes)

	perm := filePermission(file.Mode().Perm())
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeUserRead, booleanValueUsing(perm.userRead()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeUserWrite, booleanValueUsing(perm.userWrite()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeUserExecute, booleanValueUsing(perm.userExecute()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeGroupRead, booleanValueUsing(perm.groupRead()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeGroupWrite, booleanValueUsing(perm.groupWrite()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeGroupExecute, booleanValueUsing(perm.groupExecute()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeOthersRead, booleanValueUsing(perm.othersRead()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeOthersWrite, booleanValueUsing(perm.othersWrite()), attributes)
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeOthersExecute, booleanValueUsing(perm.othersExecute()), attributes)
}

func (fileAttributes *FileAttributes) setBlock(file fs.FileInfo, attributes *AllAttributes) {
	stat := file.Sys().(*syscall.Stat_t)
	if stat != nil {
		fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeBlockSize, Int64Value(int64(stat.Blksize)), attributes)
		fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeBlocks, Int64Value(stat.Blocks), attributes)
	} else {
		fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeBlockSize, StringValue("NA"), attributes)
		fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeBlocks, StringValue("NA"), attributes)
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
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeUserId, StringValue(userId), attributes)
}

func (fileAttributes *FileAttributes) setUserName(userName string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeUserName, StringValue(userName), attributes)
}

func (fileAttributes *FileAttributes) setGroupId(groupId string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeGroupId, StringValue(groupId), attributes)
}

func (fileAttributes *FileAttributes) setGroupName(groupName string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(AttributeGroupName, StringValue(groupName), attributes)
}

func (fileAttributes *FileAttributes) setMimeType(directory string, file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForUnevaluatedAttribute(AttributeMimeType, fileAttributes.filePath(directory, file), attributes)
}

func (fileAttributes *FileAttributes) setAllAliasesForEvaluatedAttribute(
	attribute string,
	value Value,
	attributes *AllAttributes,
) {
	for _, alias := range attributes.aliasesFor(attribute) {
		fileAttributes.attributes[alias] = EvaluatingValue{value: value, isEvaluated: true}
	}
}

func (fileAttributes *FileAttributes) setAllAliasesForUnevaluatedAttribute(
	attribute string,
	filePath string,
	attributes *AllAttributes,
) {
	aliasesFor := attributes.aliasesFor(attribute)
	definition := attributes.attributeDefinitionFor(attribute)

	for _, alias := range aliasesFor {
		fileAttributes.attributes[alias] = EvaluatingValue{
			isEvaluated:     false,
			filePath:        filePath,
			evaluationBlock: definition.lazyEvaluationBlock,
		}
	}
}

func (fileAttributes *FileAttributes) filePath(directory string, file fs.FileInfo) string {
	pathSeparator := string(os.PathSeparator)
	newPath := directory + pathSeparator + file.Name()
	if strings.HasSuffix(directory, pathSeparator) {
		newPath = directory + file.Name()
	}
	return newPath
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

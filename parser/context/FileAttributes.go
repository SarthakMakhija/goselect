package context

import (
	"goselect/parser/context/platform"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type EvaluatingValue struct {
	value           Value
	filePath        string
	isEvaluated     bool
	aliases         []string
	evaluationBlock AttributeLazyEvaluationBlock
}

type FileAttributes struct {
	attributes map[string]EvaluatingValue
}

func ToFileAttributes(directory string, file fs.FileInfo, ctx *ParsingApplicationContext) *FileAttributes {
	fileAttributes := newFileAttributes()
	hiddenFile, _ := platform.IsHiddenFile(file.Name())

	fileAttributes.setName(file, hiddenFile, ctx.allAttributes)
	fileAttributes.setExtension(file, hiddenFile, ctx.allAttributes)
	fileAttributes.setSize(file, ctx.allAttributes)
	fileAttributes.setFileType(directory, file, ctx.allAttributes)
	fileAttributes.setTimes(file, ctx.allAttributes)
	fileAttributes.setPath(directory, file, ctx.allAttributes)
	fileAttributes.setPermission(file, ctx.allAttributes)
	fileAttributes.setBlock(file, ctx.allAttributes)
	fileAttributes.setUserGroup(file, ctx.allAttributes)
	fileAttributes.setMimeType(directory, file, ctx.allAttributes)
	fileAttributes.setContents(directory, file, ctx.allAttributes)

	return fileAttributes
}

func (fileAttributes *FileAttributes) Get(attribute string) Value {
	evaluatingValue, ok := fileAttributes.attributes[strings.ToLower(attribute)]
	if ok {
		if evaluatingValue.isEvaluated {
			return evaluatingValue.value
		}
		value := evaluatingValue.evaluationBlock.evaluate(evaluatingValue.filePath)
		fileAttributes.setAllAliasesForEvaluatedAttribute(value, evaluatingValue.aliases)
		return value
	}
	return EmptyValue
}

func newFileAttributes() *FileAttributes {
	return &FileAttributes{attributes: make(map[string]EvaluatingValue)}
}

func (fileAttributes *FileAttributes) setName(file fs.FileInfo, hiddenFile bool, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(file.Name()), attributes.aliasesFor(AttributeName))
	if hiddenFile {
		fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(file.Name()), attributes.aliasesFor(AttributeBaseName))
	} else {
		baseName := strings.Replace(file.Name(), filepath.Ext(file.Name()), "", 1)
		fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(baseName), attributes.aliasesFor(AttributeBaseName))
	}
}

func (fileAttributes *FileAttributes) setSize(file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(Int64Value(file.Size()), attributes.aliasesFor(AttributeSize))
}

func (fileAttributes *FileAttributes) setFileType(directory string, file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(file.IsDir()), attributes.aliasesFor(AttributeNameIsDir))
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(file.Mode().IsRegular()), attributes.aliasesFor(AttributeNameIsFile))
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(file.Mode()&os.ModeSymlink == os.ModeSymlink), attributes.aliasesFor(AttributeNameIsSymbolicLink))
	if file.Mode().IsDir() {
		newPath := fileAttributes.filePath(directory, file)
		entries, _ := os.ReadDir(newPath)
		fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(len(entries) == 0), attributes.aliasesFor(AttributeNameIsEmpty))
	} else {
		fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(file.Size() == 0), attributes.aliasesFor(AttributeNameIsEmpty))
	}
	hiddenFile, _ := platform.IsHiddenFile(file.Name())
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(hiddenFile), attributes.aliasesFor(AttributeNameIsHidden))
}

func (fileAttributes *FileAttributes) setTimes(file fs.FileInfo, attributes *AllAttributes) {
	created, modified, accessed := platform.FileTimes(file)
	fileAttributes.setAllAliasesForEvaluatedAttribute(DateTimeValue(created), attributes.aliasesFor(AttributeCreatedTime))
	fileAttributes.setAllAliasesForEvaluatedAttribute(DateTimeValue(modified), attributes.aliasesFor(AttributeModifiedTime))
	fileAttributes.setAllAliasesForEvaluatedAttribute(DateTimeValue(accessed), attributes.aliasesFor(AttributeAccessedTime))
}

func (fileAttributes *FileAttributes) setPath(directory string, file fs.FileInfo, attributes *AllAttributes) {
	newPath := fileAttributes.filePath(directory, file)
	absolutePath, err := filepath.Abs(newPath)
	if err == nil {
		fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(absolutePath), attributes.aliasesFor(AttributeAbsolutePath))
	}
	fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(newPath), attributes.aliasesFor(AttributePath))
}

func (fileAttributes *FileAttributes) setExtension(file fs.FileInfo, hiddenFile bool, attributes *AllAttributes) {
	if hiddenFile {
		fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(""), attributes.aliasesFor(AttributeExtension))
	} else {
		fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(filepath.Ext(file.Name())), attributes.aliasesFor(AttributeExtension))
	}
}

func (fileAttributes *FileAttributes) setPermission(file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(file.Mode().Perm().String()), attributes.aliasesFor(AttributePermission))

	perm := filePermission(file.Mode().Perm())
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(perm.userRead()), attributes.aliasesFor(AttributeUserRead))
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(perm.userWrite()), attributes.aliasesFor(AttributeUserWrite))
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(perm.userExecute()), attributes.aliasesFor(AttributeUserExecute))
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(perm.groupRead()), attributes.aliasesFor(AttributeGroupRead))
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(perm.groupWrite()), attributes.aliasesFor(AttributeGroupWrite))
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(perm.groupExecute()), attributes.aliasesFor(AttributeGroupExecute))
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(perm.othersRead()), attributes.aliasesFor(AttributeOthersRead))
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(perm.othersWrite()), attributes.aliasesFor(AttributeOthersWrite))
	fileAttributes.setAllAliasesForEvaluatedAttribute(booleanValueUsing(perm.othersExecute()), attributes.aliasesFor(AttributeOthersExecute))
}

func (fileAttributes *FileAttributes) setBlock(file fs.FileInfo, attributes *AllAttributes) {
	blockSize, blocks := platform.FileBlocks(file)
	fileAttributes.setAllAliasesForEvaluatedAttribute(Int64Value(blockSize), attributes.aliasesFor(AttributeBlockSize))
	fileAttributes.setAllAliasesForEvaluatedAttribute(Int64Value(blocks), attributes.aliasesFor(AttributeBlocks))
}

func (fileAttributes *FileAttributes) setUserGroup(file fs.FileInfo, attributes *AllAttributes) {
	userId, userName, groupId, groupName := platform.UserGroup(file)
	fileAttributes.setUserId(userId, attributes)
	fileAttributes.setUserName(userName, attributes)
	fileAttributes.setGroupId(groupId, attributes)
	fileAttributes.setGroupName(groupName, attributes)
}

func (fileAttributes *FileAttributes) setUserId(userId string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(userId), attributes.aliasesFor(AttributeUserId))
}

func (fileAttributes *FileAttributes) setUserName(userName string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(userName), attributes.aliasesFor(AttributeUserName))
}

func (fileAttributes *FileAttributes) setGroupId(groupId string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(groupId), attributes.aliasesFor(AttributeGroupId))
}

func (fileAttributes *FileAttributes) setGroupName(groupName string, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForEvaluatedAttribute(StringValue(groupName), attributes.aliasesFor(AttributeGroupName))
}

func (fileAttributes *FileAttributes) setMimeType(directory string, file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForUnevaluatedAttribute(AttributeMimeType, fileAttributes.filePath(directory, file), attributes)
}

func (fileAttributes *FileAttributes) setContents(directory string, file fs.FileInfo, attributes *AllAttributes) {
	fileAttributes.setAllAliasesForUnevaluatedAttribute(AttributeContents, fileAttributes.filePath(directory, file), attributes)
}

func (fileAttributes *FileAttributes) setAllAliasesForEvaluatedAttribute(value Value, aliases []string) {
	for _, alias := range aliases {
		fileAttributes.attributes[alias] = EvaluatingValue{value: value, isEvaluated: true}
	}
}

func (fileAttributes *FileAttributes) setAllAliasesForUnevaluatedAttribute(
	attribute string,
	filePath string,
	attributes *AllAttributes,
) {
	aliases := attributes.aliasesFor(attribute)
	definition := attributes.attributeDefinitionFor(attribute)

	for _, alias := range aliases {
		fileAttributes.attributes[alias] = EvaluatingValue{
			isEvaluated:     false,
			filePath:        filePath,
			aliases:         aliases,
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

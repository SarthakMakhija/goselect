package context

import "strings"

type AttributeDefinition struct {
	aliases     []string
	description string
}

const (
	AttributeName               = "name"
	AttributeBaseName           = "basename"
	AttributePath               = "path"
	AttributeAbsolutePath       = "absolutepath"
	AttributeSize               = "size"
	AttributeFormattedSize      = "fmtsize"
	AttributeNameIsDir          = "isdirectory"
	AttributeNameIsFile         = "isfile"
	AttributeNameIsHidden       = "ishidden"
	AttributeNameIsEmpty        = "isempty"
	AttributeNameIsSymbolicLink = "issymboliclink"
	AttributeCreatedTime        = "createdtime"
	AttributeModifiedTime       = "modifiedtime"
	AttributeAccessedTime       = "accessedtime"
	AttributeExtension          = "extension"
	AttributePermission         = "permission"
	AttributeUserRead           = "userread"
	AttributeUserWrite          = "userwrite"
	AttributeUserExecute        = "userexecute"
	AttributeGroupRead          = "groupread"
	AttributeGroupWrite         = "groupwrite"
	AttributeGroupExecute       = "groupexecute"
	AttributeOthersRead         = "otherread"
	AttributeOthersWrite        = "otherwrite"
	AttributeOthersExecute      = "otherexecute"
	AttributeBlockSize          = "blocksize"
	AttributeBlocks             = "blocks"
	AttributeUserId             = "userid"
	AttributeUserName           = "username"
	AttributeGroupId            = "groupid"
	AttributeGroupName          = "groupname"
)

var attributeDefinitions = map[string]*AttributeDefinition{
	AttributeName: {
		aliases:     []string{"filename", "name", "fname"},
		description: "Returns the file name.",
	},
	AttributeBaseName: {
		aliases:     []string{"basename", "bname"},
		description: "Returns the basename of the file. For example, basename for 'sample.log' is 'sample'.",
	},
	AttributePath: {
		aliases:     []string{"filepath", "path", "fpath"},
		description: "Returns the relative path of the file.",
	},
	AttributeAbsolutePath: {
		aliases:     []string{"absolutepath", "apath", "abspath"},
		description: "Returns the absolute path of the file. ",
	},
	AttributeSize: {
		aliases:     []string{"filesize", "size", "fsize"},
		description: "Returns the file size in bytes.",
	},
	AttributeFormattedSize: {
		aliases:     []string{"fmtsize", "hsize"},
		description: "Returns the human readable file size.",
	},
	AttributeNameIsDir: {
		aliases:     []string{"isdir", "isdirectory"},
		description: "Returns true if the file is a directory, false otherwise.",
	},
	AttributeNameIsFile: {
		aliases:     []string{"isfile"},
		description: "Returns true if the file is a file, false otherwise.",
	},
	AttributeNameIsHidden: {
		aliases:     []string{"ishidden"},
		description: "Returns true if the file is hidden, false otherwise.",
	},
	AttributeNameIsEmpty: {
		aliases:     []string{"isempty"},
		description: "Returns true if the file is empty, false otherwise. If the file is a directory, 'isempty' returns true if there are no entries, false otherwise.",
	},
	AttributeNameIsSymbolicLink: {
		aliases:     []string{"issymboliclink", "issymlink"},
		description: "Returns true if the file is a symbolic link.",
	},
	AttributeCreatedTime: {
		aliases:     []string{"createdtime", "ctime"},
		description: "Returns the created time of the file.",
	},
	AttributeModifiedTime: {
		aliases:     []string{"modifiedtime", "mtime", "modtime"},
		description: "Returns the modified time of the file.",
	},
	AttributeAccessedTime: {
		aliases:     []string{"accessedtime", "accesstime", "atime"},
		description: "Returns the access time of the file.",
	},
	AttributeExtension: {
		aliases:     []string{"extension", "ext"},
		description: "Return the file extension. For example, extension of the file 'sample.log' is '.log'.",
	},
	AttributePermission: {
		aliases:     []string{"permission", "perm"},
		description: "Returns the file permission.",
	},
	AttributeUserRead: {
		aliases:     []string{"userread", "uread"},
		description: "Returns true if the user can read the file.",
	},
	AttributeUserWrite: {
		aliases:     []string{"userwrite", "uwrite"},
		description: "Returns true if the user can write to the file.",
	},
	AttributeUserExecute: {
		aliases:     []string{"userexecute", "uexecute"},
		description: "Returns true if the user can execute the file.",
	},
	AttributeGroupRead: {
		aliases:     []string{"groupread", "gread"},
		description: "Returns true if the group can read the file.",
	},
	AttributeGroupWrite: {
		aliases:     []string{"groupwrite", "gwrite"},
		description: "Returns true if the group can write to the file.",
	},
	AttributeGroupExecute: {
		aliases:     []string{"groupexecute", "gexecute"},
		description: "Returns true if the group can execute the file.",
	},
	AttributeOthersRead: {
		aliases:     []string{"otherread", "oread"},
		description: "Returns true if others can read the file.",
	},
	AttributeOthersWrite: {
		aliases:     []string{"otherwrite", "owrite"},
		description: "Returns true if others can write to the file.",
	},
	AttributeOthersExecute: {
		aliases:     []string{"otherexecute", "oexecute"},
		description: "Returns true if others can execute the file.",
	},
	AttributeBlockSize: {
		aliases:     []string{"blocksize", "bsize", "blksize"},
		description: "Returns the block size, usually 4096 bytes.",
	},
	AttributeBlocks: {
		aliases:     []string{"blocks", "blks"},
		description: "Returns the total number of blocks allocated to the file.",
	},
	AttributeUserId: {
		aliases:     []string{"userid", "uid"},
		description: "Returns the user id.",
	},
	AttributeUserName: {
		aliases:     []string{"username", "uname"},
		description: "Returns the user name.",
	},
	AttributeGroupId: {
		aliases:     []string{"groupid", "gid"},
		description: "Returns the group id.",
	},
	AttributeGroupName: {
		aliases:     []string{"groupname", "gname"},
		description: "Returns the group name.",
	},
}

type AllAttributes struct {
	supportedAttributes map[string]*AttributeDefinition
}

func NewAttributes() *AllAttributes {
	supportedAttributes := make(map[string]*AttributeDefinition)
	for _, definition := range attributeDefinitions {
		for _, alias := range definition.aliases {
			supportedAttributes[alias] = definition
		}
	}
	return &AllAttributes{supportedAttributes: supportedAttributes}
}

func (attributes *AllAttributes) IsASupportedAttribute(attribute string) bool {
	_, ok := attributes.supportedAttributes[strings.ToLower(attribute)]
	return ok
}

func (attributes *AllAttributes) AllAttributeWithAliases() map[string][]string {
	supportedAttributes := make(map[string][]string)
	for _, definition := range attributeDefinitions {
		supportedAttributes[definition.aliases[0]] = definition.aliases
	}
	return supportedAttributes
}

func (attributes *AllAttributes) aliasesFor(attribute string) []string {
	definition, ok := attributeDefinitions[strings.ToLower(attribute)]
	if ok {
		return definition.aliases
	}
	return []string{}
}

func IsAWildcardAttribute(attribute string) bool {
	return attribute == "*"
}

func AttributesOnWildcard() []string {
	return []string{AttributeBaseName, AttributeExtension, AttributeFormattedSize, AttributeAbsolutePath}
}

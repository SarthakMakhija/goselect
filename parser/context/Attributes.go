package context

import "strings"

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
	AttributeModifiedTime       = "modifiedtime"
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

var attributeAliases = map[string][]string{
	AttributeName:               {"filename", "name", "fname"},
	AttributeBaseName:           {"basename", "bname"},
	AttributePath:               {"filepath", "path", "fpath"},
	AttributeAbsolutePath:       {"absolutepath", "apath", "abspath"},
	AttributeSize:               {"filesize", "size", "fsize"},
	AttributeFormattedSize:      {"fmtsize", "hsize"},
	AttributeNameIsDir:          {"isdir", "isdirectory"},
	AttributeNameIsFile:         {"isfile"},
	AttributeNameIsHidden:       {"ishidden"},
	AttributeNameIsEmpty:        {"isempty"},
	AttributeNameIsSymbolicLink: {"issymboliclink", "issymlink"},
	AttributeModifiedTime:       {"modifiedTime", "mtime"},
	AttributeExtension:          {"extension", "ext"},
	AttributePermission:         {"permission", "perm"},
	AttributeUserRead:           {"userread", "uread"},
	AttributeUserWrite:          {"userwrite", "uwrite"},
	AttributeUserExecute:        {"userexecute", "uexecute"},
	AttributeGroupRead:          {"groupread", "gread"},
	AttributeGroupWrite:         {"groupwrite", "gwrite"},
	AttributeGroupExecute:       {"groupexecute", "gexecute"},
	AttributeOthersRead:         {"otherread", "oread"},
	AttributeOthersWrite:        {"otherwrite", "owrite"},
	AttributeOthersExecute:      {"otherexecute", "oexecute"},
	AttributeBlockSize:          {"blocksize", "bsize", "blksize"},
	AttributeBlocks:             {"blocks", "blks"},
	AttributeUserId:             {"userid", "uid"},
	AttributeUserName:           {"username", "uname"},
	AttributeGroupId:            {"groupid", "gid"},
	AttributeGroupName:          {"groupname", "gname"},
}

type AllAttributes struct {
	supportedAttributes map[string]bool
}

func NewAttributes() *AllAttributes {
	supportedAttributes := make(map[string]bool)
	for _, aliases := range attributeAliases {
		for _, alias := range aliases {
			supportedAttributes[alias] = true
		}
	}
	return &AllAttributes{supportedAttributes: supportedAttributes}
}

func (attributes *AllAttributes) IsASupportedAttribute(attribute string) bool {
	return attributes.supportedAttributes[strings.ToLower(attribute)]
}

func (attributes *AllAttributes) AllAttributeWithAliases() map[string][]string {
	return attributeAliases
}

func (attributes *AllAttributes) aliasesFor(attribute string) []string {
	return attributeAliases[attribute]
}

func IsAWildcardAttribute(attribute string) bool {
	return attribute == "*"
}

func AttributesOnWildcard() []string {
	return []string{AttributeBaseName, AttributeExtension, AttributeFormattedSize, AttributeAbsolutePath}
}

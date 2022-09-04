package context

import "strings"

const (
	AttributeName         = "name"
	AttributePath         = "path"
	AttributeAbsolutePath = "absolutepath"
	AttributeSize         = "size"
	AttributeModifiedTime = "modifiedtime"
	AttributeExtension    = "extension"
	AttributePermission   = "permission"
	AttributeUserId       = "userid"
	AttributeUserName     = "username"
	AttributeGroupId      = "groupid"
	AttributeGroupName    = "groupname"
)

var attributeAliases = map[string][]string{
	AttributeName:         {"filename", "name", "fname"},
	AttributePath:         {"filepath", "path", "fpath"},
	AttributeAbsolutePath: {"absolutepath", "apath"},
	AttributeSize:         {"filesize", "size", "fsize"},
	AttributeModifiedTime: {"modifiedTime", "mtime"},
	AttributeExtension:    {"extension", "ext"},
	AttributePermission:   {"permission", "perm"},
	AttributeUserId:       {"userid", "uid"},
	AttributeUserName:     {"username", "uname"},
	AttributeGroupId:      {"groupid", "gid"},
	AttributeGroupName:    {"groupname", "gname"},
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
	return []string{AttributeName, AttributeSize}
}

package context

import "strings"

const (
	AttributeName         = "name"
	AttributeSize         = "size"
	AttributeModifiedTime = "modifiedTime"
	AttributeExtension    = "extension"
	AttributePermission   = "permission"
	AttributeUserId       = "userId"
	AttributeUserName     = "userName"
	AttributeGroupId      = "groupId"
	AttributeGroupName    = "groupName"
)

var attributeAliases = map[string][]string{
	AttributeName:         {"name", "fname"},
	AttributeSize:         {"size", "fsize"},
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

func (attributes *AllAttributes) aliasesFor(attribute string) []string {
	return attributeAliases[attribute]
}

func IsAWildcardAttribute(attribute string) bool {
	return attribute == "*"
}

func AttributesOnWildcard() []string {
	return []string{AttributeName, AttributeSize}
}

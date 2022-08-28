package context

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
	AttributeName:         {"name", "Name", "fName", "fname"},
	AttributeSize:         {"size", "Size", "fSize", "fsize"},
	AttributeModifiedTime: {"modifiedTime", "mtime", "mTime"},
	AttributeExtension:    {"extension", "ext"},
	AttributePermission:   {"permission", "perm"},
	AttributeUserId:       {"userid", "userId", "uid", "uId"},
	AttributeUserName:     {"userName", "username", "uname", "uName"},
	AttributeGroupId:      {"groupid", "groupId", "gid", "gId"},
	AttributeGroupName:    {"groupName", "groupname", "gname", "gName"},
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
	return attributes.supportedAttributes[attribute]
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

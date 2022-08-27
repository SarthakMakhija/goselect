package projection

var supportedColumns = map[string]bool{
	"name":  true,
	"Name":  true,
	"fName": true,
	"size":  true,
	"fSize": true,
	"uid":   true,
	"gid":   true,
}

func IsASupportedColumn(column string) bool {
	return supportedColumns[column]
}

func isAWildcard(column string) bool {
	return column == "*"
}

func columnsOnWildcard() []string {
	return []string{"name", "size"}
}

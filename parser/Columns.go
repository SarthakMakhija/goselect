package parser

var supportedColumns = map[string]bool{
	"name":  true,
	"fName": true,
	"size":  true,
	"fSize": true,
	"uid":   true,
	"gid":   true,
}

func isAWildcard(column string) bool {
	return column == "*"
}

func isASupportedColumn(column string) bool {
	return supportedColumns[column]
}

func columnsOnWildcard() []string {
	return []string{"name", "size"}
}

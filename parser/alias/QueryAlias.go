package alias

import (
	"encoding/json"
	"fmt"
	"goselect/parser/error/messages"
	"os"
)

const fileName = "goselect.query.alias"

type QueryAliasReference struct {
	FilePath string
}

type Aliases = map[string]string

var preDefinedAliases = Aliases{
	"lsCurrent":                  "select * from .",
	"lsCurrentOrderedBySizeDesc": "select name, ext, size, abspath from . order by 3 desc",
	"lsCurrentFormattedSize":     "select name, ext, size, fmtSize(size), abspath from . order by 3 desc",
	"fileWithMaxSizeInCurrent":   "select name, size, fmtSize(size) from . order by 2 desc limit 1",
	"fileWithMinSizeInCurrent":   "select name, size, fmtSize(size) from . order by 2 limit 1",
	"totalSizeInCurrent":         "select fmtSize(sum(size)) from .",
	"minMaxSizeInCurrent":        "select fmtSize(min(size)), fmtSize(max(size)) from .",
	"textsInCurrent":             "select name, size, fmtSize(size) from . where isText(mime) order by 2 desc",
	"imagesInCurrent":            "select name, size, fmtSize(size) from . where isImage(mime) order by 2 desc",
	"pdfsInCurrent":              "select name, size, fmtSize(size) from . where isPdf(mime) order by 2 desc",
	"audiosInCurrent":            "select name, size, fmtSize(size) from . where isAudio(mime) order by 2 desc",
	"videosInCurrent":            "select name, size, fmtSize(size) from . where isVideo(mime) order by 2 desc",
	"archivesInCurrent":          "select name, size, fmtSize(size) from . where isArchive(mime) order by 2 desc",
}

type Alias struct {
	Query string `json:"query"`
	Alias string `json:"alias"`
}

func NewQueryAlias() *QueryAliasReference {
	filePath := "." + string(os.PathSeparator) + fileName
	return &QueryAliasReference{FilePath: filePath}
}

func (queryAlias *QueryAliasReference) Add(alias Alias) error {
	aliases, err := queryAlias.readAndUnmarshal()
	if err != nil {
		return fmt.Errorf(messages.ErrorMessageQueryAliasAddPrefixWithExistingError, err)
	}
	if queryAlias.isAliasPresent(aliases, alias) {
		return fmt.Errorf(messages.ErrorMessageQueryAliasAlreadyExists, alias, queryAlias.FilePath)
	}
	aliases[alias.Alias] = alias.Query
	bytes, err := queryAlias.marshal(aliases)
	if err != nil {
		return fmt.Errorf(messages.ErrorMessageQueryAliasAddPrefixWithExistingError, err)
	}
	return queryAlias.write(bytes)
}

func (queryAlias *QueryAliasReference) GetQueryBy(alias string) (string, bool, error) {
	aliases, err := queryAlias.readAndUnmarshal()
	if err != nil {
		return "", false, fmt.Errorf(messages.ErrorMessageQueryAliasGetPrefixWithExistingError, err)
	}
	query, ok := aliases[alias]
	return query, ok, nil
}

func (queryAlias QueryAliasReference) All() (Aliases, error) {
	aliases, err := queryAlias.readAndUnmarshal()
	if err != nil {
		return nil, fmt.Errorf(messages.ErrorMessageQueryAliasGetAllPrefixWithExistingError, err)
	}
	return aliases, nil
}

func (queryAlias QueryAliasReference) PredefinedAliases() Aliases {
	return preDefinedAliases
}

func (queryAlias QueryAliasReference) isAliasPresent(aliases Aliases, alias Alias) bool {
	if _, ok := aliases[alias.Alias]; ok {
		return true
	}
	return false
}

func (queryAlias QueryAliasReference) readAndUnmarshal() (Aliases, error) {
	contents, err := queryAlias.readAll()
	if err != nil {
		return nil, err
	}
	aliases, err := queryAlias.unMarshal(contents)
	if err != nil {
		return nil, err
	}
	return aliases, nil
}

func (queryAlias QueryAliasReference) readAll() ([]byte, error) {
	_, err := os.Stat(queryAlias.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []byte{}, nil
		}
		return nil, err
	}
	return os.ReadFile(queryAlias.FilePath)
}

func (queryAlias *QueryAliasReference) marshal(aliases Aliases) ([]byte, error) {
	return json.Marshal(aliases)
}

func (queryAlias *QueryAliasReference) unMarshal(contents []byte) (Aliases, error) {
	if len(contents) != 0 {
		var aliases Aliases
		err := json.Unmarshal(contents, &aliases)
		if err != nil {
			return Aliases{}, err
		}
		return queryAlias.merge(aliases, preDefinedAliases), nil
	}
	return preDefinedAliases, nil
}

func (queryAlias *QueryAliasReference) write(bytes []byte) error {
	return os.WriteFile(queryAlias.FilePath, bytes, 0644)
}

func (queryAlias *QueryAliasReference) merge(aliases, preDefinedAliases Aliases) Aliases {
	merged := make(Aliases)
	for alias, query := range aliases {
		merged[alias] = query
	}
	for alias, query := range preDefinedAliases {
		merged[alias] = query
	}
	return merged
}

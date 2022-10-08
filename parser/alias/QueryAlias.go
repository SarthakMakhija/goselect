package alias

import (
	"encoding/json"
	"fmt"
	"goselect/parser/error/messages"
	"os"
)

const fileName = "goselect.alias"

type QueryAliasReference struct {
	filePath string
}

type Aliases = map[string]string

type Alias struct {
	Query string `json:"query"`
	Alias string `json:"alias"`
}

func NewQueryAlias() *QueryAliasReference {
	filePath := "." + string(os.PathSeparator) + fileName
	return &QueryAliasReference{filePath: filePath}
}

func (queryAlias *QueryAliasReference) Add(alias Alias) error {
	aliases, err := queryAlias.readAndUnmarshal()
	if err != nil {
		return err
	}
	if queryAlias.isAliasPresent(aliases, alias) {
		return fmt.Errorf(messages.ErrorMessageQueryAliasAlreadyExists, alias, queryAlias.filePath)
	}
	aliases[alias.Alias] = alias.Query
	bytes, err := queryAlias.marshal(aliases)
	if err != nil {
		return err
	}
	return queryAlias.write(bytes)
}

func (queryAlias *QueryAliasReference) GetQueryBy(alias string) (string, bool, error) {
	aliases, err := queryAlias.readAndUnmarshal()
	if err != nil {
		return "", false, err
	}
	query, ok := aliases[alias]
	return query, ok, nil
}

func (queryAlias QueryAliasReference) All() (Aliases, error) {
	return queryAlias.readAndUnmarshal()
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
	_, err := os.Stat(queryAlias.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []byte{}, nil
		}
		return nil, err
	}
	return os.ReadFile(queryAlias.filePath)
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
		return aliases, nil
	}
	return Aliases{}, nil
}

func (queryAlias *QueryAliasReference) write(bytes []byte) error {
	return os.WriteFile(queryAlias.filePath, bytes, 0644)
}

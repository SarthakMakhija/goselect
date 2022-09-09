package context

import (
	"errors"
	"goselect/parser/error/messages"
	"strings"
	"time"
)

const (
	layoutDate              = "2006-01-02"
	layoutDateTimestamp     = "2006-01-02T15:04:05"
	layoutDateTimestampFull = "2006-01-02T15:04:05.000Z"
)

type FormatDefinition struct {
	Format string
	Id     string
}

var formatDefinitions = map[string]FormatDefinition{
	"dt": {
		Format: layoutDate,
		Id:     "dt",
	},
	"ts": {
		Format: layoutDateTimestamp,
		Id:     "ts",
	},
	"tsfull": {
		Format: layoutDateTimestampFull,
		Id:     "tsfull",
	},
}

func parse(str, id string) (time.Time, error) {
	idToLower := strings.ToLower(id)
	definition, ok := formatDefinitions[idToLower]
	if !ok {
		return time.Time{}, errors.New(messages.ErrorMessageUnsupportedDateTimeFormat)
	}
	return time.Parse(definition.Format, str)
}

func SupportedFormats() map[string]FormatDefinition {
	return formatDefinitions
}

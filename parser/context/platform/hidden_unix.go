//go:build !windows
// +build !windows

package platform

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func IsHiddenFile(_, filename string) (bool, error) {
	return filename[0:1] == ".", nil
}

func BaseName(hidden bool, file fs.FileInfo) string {
	if hidden {
		return file.Name()
	}
	return strings.Replace(file.Name(), filepath.Ext(file.Name()), "", 1)
}

func Extension(hidden bool, file fs.FileInfo) string {
	if hidden {
		return ""
	}
	return filepath.Ext(file.Name())
}

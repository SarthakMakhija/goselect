//go:build windows
// +build windows

package platform

import (
	"io/fs"
	"path/filepath"
	"strings"
	"syscall"
)

func IsHiddenFile(path, filename string) (bool, error) {
	pointer, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return false, err
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}
	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}

func BaseName(_ bool, file fs.FileInfo) string {
	return strings.Replace(file.Name(), filepath.Ext(file.Name()), "", 1)
}

func Extension(_ bool, file fs.FileInfo) string {
	return filepath.Ext(file.Name())
}

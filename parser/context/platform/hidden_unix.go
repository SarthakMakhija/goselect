//go:build !windows
// +build !windows

package platform

func IsHiddenFile(path, filename string) (bool, error) {
	return filename[0:1] == ".", nil
}

//go:build !windows
// +build !windows

package platform

func IsHiddenFile(_, filename string) (bool, error) {
	return filename[0:1] == ".", nil
}

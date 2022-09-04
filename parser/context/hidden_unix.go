//go:build !windows
// +build !windows

package context

func isHiddenFile(filename string) (bool, error) {
	return filename[0:1] == ".", nil
}

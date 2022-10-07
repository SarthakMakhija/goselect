//go:build unit && windows
// +build unit,windows

package platform

import (
	"os"
	"syscall"
	"testing"
)

func TestIsAHiddenFile(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "hidden")
	path := directory + string(os.PathSeparator) + "hidden.txt"
	hiddenFile, _ := os.Create(path)
	_ = setHidden(hiddenFile.Name())

	defer func() {
		hiddenFile.Close()
		os.RemoveAll(directory)
	}()

	isHidden, _ := IsHiddenFile(path, "hidden.txt")
	if !isHidden {
		t.Fatalf("Expected file to be hidden but was not")
	}
}

func TestIsNotAHiddenFile(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "notHidden")
	path := directory + string(os.PathSeparator) + "notHidden.txt"
	file, _ := os.Create(path)

	defer func() {
		file.Close()
		os.RemoveAll(directory)
	}()

	isHidden, _ := IsHiddenFile(path, "notHidden.txt")
	if isHidden {
		t.Fatalf("Expected file to not be hidden but was hidden")
	}
}

func TestBaseName(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "hello")
	path := directory + string(os.PathSeparator) + "hello.txt"
	file, _ := os.Create(path)

	defer func() {
		file.Close()
		os.RemoveAll(directory)
	}()

	fileInfo, _ := os.Stat(file.Name())
	baseName := BaseName(false, fileInfo)

	if baseName != "hello" {
		t.Fatalf("Expected basename to be %v received %v", "hello", baseName)
	}
}

func TestExtension1(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "hello")
	path := directory + string(os.PathSeparator) + "hello.txt"
	file, _ := os.Create(path)

	defer func() {
		file.Close()
		os.RemoveAll(directory)
	}()

	fileInfo, _ := os.Stat(file.Name())
	extension := Extension(false, fileInfo)

	if extension != ".txt" {
		t.Fatalf("Expected extension to be %v received %v", ".txt", extension)
	}
}

func TestExtension2(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "temp")
	path := directory + string(os.PathSeparator) + "hello"
	file, _ := os.Create(path)

	defer func() {
		file.Close()
		os.RemoveAll(directory)
	}()

	fileInfo, _ := os.Stat(file.Name())
	extension := Extension(false, fileInfo)

	if extension != "" {
		t.Fatalf("Expected extension to be %v received %v", "", extension)
	}
}

func setHidden(path string) error {
	filenameW, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	err = syscall.SetFileAttributes(filenameW, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		return err
	}
	return nil
}

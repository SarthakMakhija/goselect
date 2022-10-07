//go:build unit && !windows
// +build unit,!windows

package platform

import (
	"os"
	"syscall"
	"testing"
)

func TestIsAHiddenFile(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "hidden")
	path := directory + string(os.PathSeparator) + ".Make"
	hiddenFile, _ := os.Create(path)

	defer func() {
		hiddenFile.Close()
		os.RemoveAll(directory)
	}()

	isHidden, _ := IsHiddenFile(path, ".Make")
	if !isHidden {
		t.Fatalf("Expected file to be hidden but was not")
	}
}

func TestIsNotAHiddenFile(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "notHidden")
	path := directory + string(os.PathSeparator) + "test"
	file, _ := os.Create(path)

	defer func() {
		file.Close()
		os.RemoveAll(directory)
	}()

	isHidden, _ := IsHiddenFile(path, "test")
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

func TestBaseNameForHiddenFile(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "hidden")
	path := directory + string(os.PathSeparator) + ".Make"
	file, _ := os.Create(path)

	defer func() {
		file.Close()
		os.RemoveAll(directory)
	}()

	fileInfo, _ := os.Stat(file.Name())
	baseName := BaseName(false, fileInfo)

	if baseName != ".Make" {
		t.Fatalf("Expected basename to be %v received %v", ".Make", baseName)
	}
}

func TestExtension(t *testing.T) {
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

func TestExtensionForHiddenFile(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "hidden")
	path := directory + string(os.PathSeparator) + ".Make"
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

//go:build unit && windows
// +build unit,windows

package context

import (
	"os"
	"reflect"
	"syscall"
	"testing"
)

func TestFileExtensionForHiddenFile(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "hidden")
	hiddenFile, _ := os.Create(directory + string(os.PathSeparator) + "hidden.txt")
	defer os.RemoveAll(directory)

	_ = setHidden(hiddenFile.Name())

	file, err := os.Stat(hiddenFile.Name())
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes(directory, file, context)
	extension := fileAttributes.Get(AttributeExtension).GetAsString()

	if extension != ".txt" {
		t.Fatalf("Expected file extension to be %v, received %v", ".txt", extension)
	}
}

func TestFileBaseNameForHiddenFile(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "hidden")
	hiddenFile, _ := os.Create(directory + string(os.PathSeparator) + "hidden.txt")
	defer os.RemoveAll(directory)

	file, err := os.Stat(hiddenFile.Name())
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes(directory, file, context)
	baseName := fileAttributes.Get(AttributeBaseName).GetAsString()

	if baseName != "hidden" {
		t.Fatalf("Expected file baseName to be %v, received %v", "hidden", baseName)
	}
}

func TestFilePath(t *testing.T) {
	file, err := os.Stat("..\\test\\resources\\TestResultsWithProjections\\single\\TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("..\\test\\resources\\TestResultsWithProjections\\single\\", file, context)

	path := fileAttributes.Get(AttributePath).GetAsString()
	expected := "..\\test\\resources\\TestResultsWithProjections\\single\\TestResultsWithProjections_A.txt"

	if path != expected {
		t.Fatalf("Expected file path to be %v, received %v", expected, path)
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

func TestFilePermission(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/empty/", file, context)
	permission := fileAttributes.Get(AttributePermission).GetAsString()
	expected := "-rw-rw-rw-"

	if permission != expected {
		t.Fatalf("Expected permission to be %v, received %v", expected, permission)
	}
}

func TestFilePermissionForGroup(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/empty/", file, context)

	groupRead, _ := fileAttributes.Get(AttributeGroupRead).GetBoolean()
	groupWrite, _ := fileAttributes.Get(AttributeGroupWrite).GetBoolean()
	groupExecute, _ := fileAttributes.Get(AttributeGroupExecute).GetBoolean()

	expected := []bool{true, true, false}
	received := []bool{groupRead, groupWrite, groupExecute}

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf("Expected permissions for group to be %v, received %v", expected, received)
	}
}

func TestFilePermissionForOthers(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/empty/", file, context)

	othersRead, _ := fileAttributes.Get(AttributeOthersRead).GetBoolean()
	othersWrite, _ := fileAttributes.Get(AttributeOthersWrite).GetBoolean()
	othersExecute, _ := fileAttributes.Get(AttributeOthersExecute).GetBoolean()

	expected := []bool{true, true, false}
	received := []bool{othersRead, othersWrite, othersExecute}

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf("Expected permissions for others to be %v, received %v", expected, received)
	}
}

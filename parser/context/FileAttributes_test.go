package context

import (
	"os"
	"reflect"
	"testing"
)

func TestFileName(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/single/", file, context)
	name := fileAttributes.Get(AttributeName).GetAsString()

	if name != "TestResultsWithProjections_A.txt" {
		t.Fatalf("Expected file name to be %v, received %v", "TestResultsWithProjections_A.txt", name)
	}
}

func TestFileExtension(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/single/", file, context)
	extension := fileAttributes.Get(AttributeExtension).GetAsString()

	if extension != ".txt" {
		t.Fatalf("Expected file extension to be %v, received %v", ".txt", extension)
	}
}

func TestFileBaseName(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/single/", file, context)
	baseName := fileAttributes.Get(AttributeBaseName).GetAsString()

	if baseName != "TestResultsWithProjections_A" {
		t.Fatalf("Expected file baseName to be %v, received %v", "TestResultsWithProjections_A", baseName)
	}
}

func TestFilePath(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/single/", file, context)

	path := fileAttributes.Get(AttributePath).GetAsString()
	expected := "../resource/TestResultsWithProjections/single/TestResultsWithProjections_A.txt"

	if path != expected {
		t.Fatalf("Expected file path to be %v, received %v", expected, path)
	}
}

func TestFileSize(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/single/", file, context)
	size := fileAttributes.Get(AttributeFormattedSize).GetAsString()

	if size != "58 B" {
		t.Fatalf("Expected file baseName to be %v, received %v", "58 B", size)
	}
}

func TestFileType1(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/single/", file, context)
	isFile, _ := fileAttributes.Get(AttributeNameIsFile).GetBoolean()

	if isFile != true {
		t.Fatalf("Expected TestResultsWithProjections_A to be a file but was not")
	}
}

func TestFileType2(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/single/", file, context)
	isDirectory, _ := fileAttributes.Get(AttributeNameIsDir).GetBoolean()

	if isDirectory != false {
		t.Fatalf("Expected TestResultsWithProjections_A to not be a directory but was")
	}
}

func TestFileType3(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/single/", file, context)
	isHidden, _ := fileAttributes.Get(AttributeNameIsHidden).GetBoolean()

	if isHidden != false {
		t.Fatalf("Expected TestResultsWithProjections_A to not be hidden but was")
	}
}

func TestFileType4(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/empty/", file, context)
	isEmpty, _ := fileAttributes.Get(AttributeNameIsEmpty).GetBoolean()

	if isEmpty != true {
		t.Fatalf("Expected Empty.log to be empty but was not")
	}
}

func TestFileType5(t *testing.T) {
	directoryName, _ := os.MkdirTemp(".", "file-type-dir")
	_, _ = os.CreateTemp(directoryName, "file-type-file")

	defer os.RemoveAll(directoryName)

	file, err := os.Stat(directoryName)
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes(".", file, context)
	isEmpty, _ := fileAttributes.Get(AttributeNameIsEmpty).GetBoolean()

	if isEmpty != false {
		t.Fatalf("Expected file-type-dir to be non-empty but was")
	}
}

func TestFileType6(t *testing.T) {
	directoryName, _ := os.MkdirTemp(".", "file-type-dir")
	defer os.RemoveAll(directoryName)

	file, err := os.Stat(directoryName)
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes(".", file, context)
	isEmpty, _ := fileAttributes.Get(AttributeNameIsEmpty).GetBoolean()

	if isEmpty != true {
		t.Fatalf("Expected file-type-dir to be empty but was not")
	}
}

func TestFilePermission(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/empty/", file, context)
	permission := fileAttributes.Get(AttributePermission).GetAsString()
	expected := "-rw-r--r--"

	if permission != expected {
		t.Fatalf("Expected permission to be %v, received %v", expected, permission)
	}
}

func TestFilePermissionForUsers(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/empty/", file, context)

	userRead, _ := fileAttributes.Get(AttributeUserRead).GetBoolean()
	userWrite, _ := fileAttributes.Get(AttributeUserWrite).GetBoolean()
	userExecute, _ := fileAttributes.Get(AttributeUserExecute).GetBoolean()

	expected := []bool{true, true, false}
	received := []bool{userRead, userWrite, userExecute}

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf("Expected permissions for user to be %v, received %v", expected, received)
	}
}

func TestFilePermissionForGroup(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/empty/", file, context)

	groupRead, _ := fileAttributes.Get(AttributeGroupRead).GetBoolean()
	groupWrite, _ := fileAttributes.Get(AttributeGroupWrite).GetBoolean()
	groupExecute, _ := fileAttributes.Get(AttributeGroupExecute).GetBoolean()

	expected := []bool{true, false, false}
	received := []bool{groupRead, groupWrite, groupExecute}

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf("Expected permissions for group to be %v, received %v", expected, received)
	}
}

func TestFilePermissionForOthers(t *testing.T) {
	file, err := os.Stat("../resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../resource/TestResultsWithProjections/empty/", file, context)

	othersRead, _ := fileAttributes.Get(AttributeOthersRead).GetBoolean()
	othersWrite, _ := fileAttributes.Get(AttributeOthersWrite).GetBoolean()
	othersExecute, _ := fileAttributes.Get(AttributeOthersExecute).GetBoolean()

	expected := []bool{true, false, false}
	received := []bool{othersRead, othersWrite, othersExecute}

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf("Expected permissions for others to be %v, received %v", expected, received)
	}
}

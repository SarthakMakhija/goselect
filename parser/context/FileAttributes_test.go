//go:build unit
// +build unit

package context

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestUnsupportedAttribute(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/single/", file, context)
	value := fileAttributes.Get("unknown").GetAsString()

	if value != "" {
		t.Fatalf("Expected value for the unknown attribute to be blank, received %v", value)
	}
}

func TestFileName(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/single/", file, context)
	name := fileAttributes.Get(AttributeName).GetAsString()

	if name != "TestResultsWithProjections_A.txt" {
		t.Fatalf("Expected file name to be %v, received %v", "TestResultsWithProjections_A.txt", name)
	}
}

func TestFileNameForHiddenFile(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/hidden/.Make")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/hidden/", file, context)
	name := fileAttributes.Get(AttributeName).GetAsString()

	if name != ".Make" {
		t.Fatalf("Expected file name to be %v, received %v", ".Make", name)
	}
}

func TestFileExtension(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/single/", file, context)
	extension := fileAttributes.Get(AttributeExtension).GetAsString()

	if extension != ".txt" {
		t.Fatalf("Expected file extension to be %v, received %v", ".txt", extension)
	}
}

func TestFileBaseName(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/single/", file, context)
	baseName := fileAttributes.Get(AttributeBaseName).GetAsString()

	if baseName != "TestResultsWithProjections_A" {
		t.Fatalf("Expected file baseName to be %v, received %v", "TestResultsWithProjections_A", baseName)
	}
}

func TestFileSize(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/single/", file, context)
	size := fileAttributes.Get(AttributeSize).GetAsString()

	if size != "58" {
		t.Fatalf("Expected file size to be %v, received %v", "58", size)
	}
}

func TestFileType1(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/single/", file, context)
	isFile, _ := fileAttributes.Get(AttributeNameIsFile).GetBoolean()

	if isFile != true {
		t.Fatalf("Expected TestResultsWithProjections_A to be a file but was not")
	}
}

func TestFileType2(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/single/", file, context)
	isDirectory, _ := fileAttributes.Get(AttributeNameIsDir).GetBoolean()

	if isDirectory != false {
		t.Fatalf("Expected TestResultsWithProjections_A to not be a directory but was")
	}
}

func TestFileType3(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/single/", file, context)
	isHidden, _ := fileAttributes.Get(AttributeNameIsHidden).GetBoolean()

	if isHidden != false {
		t.Fatalf("Expected TestResultsWithProjections_A to not be hidden but was")
	}
}

func TestFileType4(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/empty/", file, context)
	isEmpty, _ := fileAttributes.Get(AttributeNameIsEmpty).GetBoolean()

	if isEmpty != true {
		t.Fatalf("Expected Empty.log to be empty but was not")
	}
}

func TestFileType5(t *testing.T) {
	directoryName, _ := os.MkdirTemp(".", "file-type-dir")
	tempFile, _ := os.CreateTemp(directoryName, "file-type-file")
	defer func() {
		tempFile.Close()
		os.RemoveAll(directoryName)
	}()

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

func TestFilePermissionForUsers(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/empty/", file, context)

	userRead, _ := fileAttributes.Get(AttributeUserRead).GetBoolean()
	userWrite, _ := fileAttributes.Get(AttributeUserWrite).GetBoolean()
	userExecute, _ := fileAttributes.Get(AttributeUserExecute).GetBoolean()

	expected := []bool{true, true, false}
	received := []bool{userRead, userWrite, userExecute}

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf("Expected permissions for user to be %v, received %v", expected, received)
	}
}

func TestAccessTime(t *testing.T) {
	directoryName, _ := os.MkdirTemp(".", "access-time-dir")
	newFile, _ := os.CreateTemp(directoryName, "access-time-file")
	defer func() {
		newFile.Close()
		os.RemoveAll(directoryName)
	}()

	defer os.RemoveAll(directoryName)

	file, err := os.Stat(fmt.Sprintf("%v", newFile.Name()))
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes(fmt.Sprintf("%v", directoryName), file, context)

	accessTimeValue := fileAttributes.Get(AttributeAccessedTime)
	expected := formatDate(time.Now()).GetAsString()
	actual := formatDate(accessTimeValue.timeValue).GetAsString()

	if expected != actual {
		t.Fatalf("Expected access date/time to be %v, received %v", expected, actual)
	}
}

func TestModifiedTime(t *testing.T) {
	directoryName, _ := os.MkdirTemp(".", "modified-time-dir")
	newFile, _ := os.CreateTemp(directoryName, "modified-time-file")
	defer func() {
		newFile.Close()
		os.RemoveAll(directoryName)
	}()

	defer os.RemoveAll(directoryName)

	file, err := os.Stat(fmt.Sprintf("%v", newFile.Name()))
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes(fmt.Sprintf("%v", directoryName), file, context)

	modifiedTime := fileAttributes.Get(AttributeModifiedTime)
	expected := formatDate(time.Now()).GetAsString()
	actual := formatDate(modifiedTime.timeValue).GetAsString()

	if expected != actual {
		t.Fatalf("Expected modified date/time to be %v, received %v", expected, actual)
	}
}

func TestCreatedTime(t *testing.T) {
	directoryName, _ := os.MkdirTemp(".", "created-time-dir")
	newFile, _ := os.CreateTemp(directoryName, "created-time-file")
	defer func() {
		newFile.Close()
		os.RemoveAll(directoryName)
	}()

	defer os.RemoveAll(directoryName)

	file, err := os.Stat(fmt.Sprintf("%v", newFile.Name()))
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes(fmt.Sprintf("%v", directoryName), file, context)

	createdTime := fileAttributes.Get(AttributeCreatedTime)
	expected := formatDate(time.Now()).GetAsString()
	actual := formatDate(createdTime.timeValue).GetAsString()

	if expected != actual {
		t.Fatalf("Expected created date/time to be %v, received %v", expected, actual)
	}
}

func TestMimeType1(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/empty/", file, context)
	mimeType := fileAttributes.Get(AttributeMimeType).GetAsString()
	expected := "text/plain"

	if mimeType != expected {
		t.Fatalf("Expected mime type to be %v, received %v", expected, mimeType)
	}
}

func TestMimeType2(t *testing.T) {
	file, err := os.Stat("../test/resources/images/where.png")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/images", file, context)
	mimeType := fileAttributes.Get(AttributeMimeType).GetAsString()
	expected := "image/png"

	if mimeType != expected {
		t.Fatalf("Expected mime type to be %v, received %v", expected, mimeType)
	}
}

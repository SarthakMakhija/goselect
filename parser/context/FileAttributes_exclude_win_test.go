//go:build unit && !windows
// +build unit,!windows

package context

import (
	"os"
	"os/user"
	"reflect"
	"testing"
)

func TestFileExtensionForHiddenFile(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/hidden/.Make")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/hidden/", file, context)
	extension := fileAttributes.Get(AttributeExtension).GetAsString()

	if extension != "" {
		t.Fatalf("Expected file extension to be %v, received %v", "", extension)
	}
}

func TestFileBaseNameForHiddenFile(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/hidden/.Make")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/hidden/", file, context)
	baseName := fileAttributes.Get(AttributeBaseName).GetAsString()

	if baseName != ".Make" {
		t.Fatalf("Expected file baseName to be %v, received %v", ".Make", baseName)
	}
}

func TestFilePath(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/single/", file, context)

	path := fileAttributes.Get(AttributePath).GetAsString()
	expected := "../test/resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt"

	if path != expected {
		t.Fatalf("Expected file path to be %v, received %v", expected, path)
	}
}

func TestUserGroup(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/empty/", file, context)

	currentUser, _ := user.Current()

	userId := fileAttributes.Get(AttributeUserId)
	userName := fileAttributes.Get(AttributeUserName)
	groupId := fileAttributes.Get(AttributeGroupId)
	groupName := fileAttributes.Get(AttributeGroupName)

	expectedUserId := currentUser.Uid
	expectedUserName := currentUser.Username

	expectedGroupId := currentUser.Gid
	expectedGroup, _ := user.LookupGroupId(expectedGroupId)
	expectedGroupName := expectedGroup.Name

	if userId.GetAsString() != expectedUserId {
		t.Fatalf("Expected userId to be %v, received %v", expectedUserId, userId.GetAsString())
	}
	if userName.GetAsString() != expectedUserName {
		t.Fatalf("Expected userName to be %v, received %v", expectedUserName, userName.GetAsString())
	}
	if groupId.GetAsString() != expectedGroupId {
		t.Fatalf("Expected groupId to be %v, received %v", expectedGroupId, groupId.GetAsString())
	}
	if groupName.GetAsString() != expectedGroupName {
		t.Fatalf("Expected groupName to be %v, received %v", expectedGroupName, groupName.GetAsString())
	}
}

func TestFilePermission(t *testing.T) {
	file, err := os.Stat("../test/resources/TestResultsWithProjections/empty/Empty.log")
	if err != nil {
		panic(err)
	}
	context := NewContext(nil, NewAttributes())
	fileAttributes := ToFileAttributes("../test/resources/TestResultsWithProjections/empty/", file, context)
	permission := fileAttributes.Get(AttributePermission).GetAsString()
	expected := "-rw-r--r--"

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

	expected := []bool{true, false, false}
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

	expected := []bool{true, false, false}
	received := []bool{othersRead, othersWrite, othersExecute}

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf("Expected permissions for others to be %v, received %v", expected, received)
	}
}

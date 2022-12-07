package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	root := ParseFileSystem("input_test.txt")

	dirList := make([]Directory, 0)
	GetAllSubDirs(root, &dirList)
	dirList = FilterDirsByMaxSize(dirList, 100000)
	totSize := GetTotalDirSize(dirList)

	expected := 95437
	actual := totSize
	if actual != expected {
		t.Errorf("Expected: %d, actual: %d\n", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	root := ParseFileSystem("input_test.txt")

	totalSpace := 70000000
	usedSpace := root.GetSize()
	unusedSpace := totalSpace - usedSpace
	spaceToDelete := 30000000 - unusedSpace

	dirList := make([]Directory, 0)
	GetAllSubDirs(root, &dirList)
	dir := FindSmallestDirByMinSize(dirList, spaceToDelete)

	expected := 24933642
	actual := dir.GetSize()
	if actual != expected {
		t.Errorf("Expected: %d, actual: %d\n", expected, actual)
	}
}

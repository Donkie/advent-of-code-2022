package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

// Directory represents a directory in the filesystem
type Directory struct {
	parent *Directory
	files  map[string]int
	dirs   map[string]Directory
}

func makeDirectory() Directory {
	return Directory{
		parent: nil,
		files:  make(map[string]int),
		dirs:   make(map[string]Directory),
	}
}

// GetSize returns the size of all files and subdirs of this dir
func (dir Directory) GetSize() int {
	sum := 0
	for _, contents := range dir.dirs {
		sum += contents.GetSize()
	}
	for _, size := range dir.files {
		sum += size
	}
	return sum
}

func (dir Directory) toStringBuilder(b *strings.Builder, level int) {
	for name, contents := range dir.dirs {
		b.WriteString(strings.Repeat("  ", level))
		b.WriteString("/")
		b.WriteString(name)
		b.WriteString(" ")
		b.WriteString(fmt.Sprint(contents.GetSize()))
		b.WriteString("\n")
		contents.toStringBuilder(b, level+1)
	}
	for name, size := range dir.files {
		b.WriteString(strings.Repeat("  ", level))
		b.WriteString(name)
		b.WriteString(" ")
		b.WriteString(fmt.Sprint(size))
		b.WriteString("\n")
	}
}

// String outputs a tree-representation of the directory and its subdirs
func (dir Directory) String() string {
	var sb strings.Builder
	dir.toStringBuilder(&sb, 0)
	return sb.String()
}

// FileSystemTraverser represents a file system and a terminal-type CD
type FileSystemTraverser struct {
	root Directory
	cd   *Directory
}

func makeFileSystemTraverser() FileSystemTraverser {
	var fst FileSystemTraverser
	fst.root = makeDirectory()
	fst.cd = &fst.root
	return fst
}

// AddFile adds a new file to the current directory
func (fst *FileSystemTraverser) AddFile(name string, size int) {
	fst.cd.files[name] = size
}

// AddDir adds a new empty directory to the current directory
func (fst *FileSystemTraverser) AddDir(name string) *Directory {
	newDir := makeDirectory()
	newDir.parent = fst.cd
	fst.cd.dirs[name] = newDir
	return &newDir
}

// CdIn changes the current directory to the specified directory. (cd X)
// If the directory doesn't exist, it is created.
func (fst *FileSystemTraverser) CdIn(dirName string) {
	if dir, exists := fst.cd.dirs[dirName]; exists {
		// Directory already exists, simply CD into it
		fst.cd = &dir
	} else {
		// Directory doesn't exist, create it and CD into it
		newDir := fst.AddDir(dirName)
		fst.cd = newDir
	}
}

// CdOut changes the current directory up one level. (cd ..)
// Exits the program if the current directory doesn't have any parent.
func (fst *FileSystemTraverser) CdOut() {
	if fst.cd.parent == nil {
		log.Fatal("Can't CD out, CD has no parent")
		return
	}
	fst.cd = fst.cd.parent
}

// CdRoot changes the current directory to the root dir (cd /)
func (fst *FileSystemTraverser) CdRoot() {
	fst.cd = &fst.root
}

// ParseFileSystem parses the problem input command line structure into a filesystem and returns the root directory
func ParseFileSystem(fileName string) Directory {
	fst := makeFileSystemTraverser()

	readingLS := false

	lib.ParseInputByLine(fileName, func(line string) error {
		if line[0] == '$' {
			cmdline := strings.Split(strings.TrimPrefix(line, "$ "), " ")
			cmd := cmdline[0]
			args := cmdline[1:]

			if cmd == "cd" {
				if args[0] == "/" {
					fst.CdRoot()
				} else if args[0] == ".." {
					fst.CdOut()
				} else {
					dirName := args[0]
					fst.CdIn(dirName)
				}
			} else if cmd == "ls" {
				readingLS = true
			}
		} else if readingLS {
			args := strings.Split(line, " ")
			if args[0] == "dir" {
				fst.AddDir(args[1])
			} else {
				fileSize, err := strconv.Atoi(args[0])
				if err != nil {
					return err
				}

				fileName := args[1]
				fst.AddFile(fileName, fileSize)
			}
		}

		return nil
	})

	return fst.root
}

// GetAllSubDirs recursively collects all subdirs of the supplied dir into the supplied list
func GetAllSubDirs(dir Directory, dirList *[]Directory) {
	for _, contents := range dir.dirs {
		*dirList = append(*dirList, contents)
		GetAllSubDirs(contents, dirList)
	}
}

// FilterDirsByMaxSize returns a filtered version of the input list with all directories smaller than the
// condition removed.
func FilterDirsByMaxSize(dirList []Directory, maxSize int) (ret []Directory) {
	for _, dir := range dirList {
		if dir.GetSize() <= maxSize {
			ret = append(ret, dir)
		}
	}
	return
}

// FindSmallestDirByMinSize returns the smallest directory in the list which fulfills the minimum size condition
func FindSmallestDirByMinSize(dirList []Directory, minSize int) *Directory {
	sort.SliceStable(dirList, func(i, j int) bool {
		return dirList[i].GetSize() < dirList[j].GetSize()
	})
	for i := 0; i < len(dirList); i++ {
		if dirList[i].GetSize() >= minSize {
			return &dirList[i]
		}
	}
	return nil
}

// GetTotalDirSize returns the summed size of all directories in the list
func GetTotalDirSize(dirList []Directory) int {
	sum := 0
	for _, dir := range dirList {
		sum += dir.GetSize()
	}
	return sum
}

func main() {
	// Part 1
	root := ParseFileSystem("input.txt")
	dirList := make([]Directory, 0)
	GetAllSubDirs(root, &dirList)
	dirList = FilterDirsByMaxSize(dirList, 100000)
	totSize := GetTotalDirSize(dirList)

	log.Printf("Part 1 - Total size: %d", totSize)

	// Part 2
	totalSpace := 70000000
	usedSpace := root.GetSize()
	unusedSpace := totalSpace - usedSpace
	spaceToDelete := 30000000 - unusedSpace

	dirList = make([]Directory, 0)
	GetAllSubDirs(root, &dirList)
	dir := FindSmallestDirByMinSize(dirList, spaceToDelete)

	log.Printf("Part 2 - Total size: %d", dir.GetSize())
}

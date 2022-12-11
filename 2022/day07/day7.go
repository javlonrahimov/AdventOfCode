package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"strconv"
	"strings"
)

type Dir struct {
	parent *Dir
	dirs   map[string]Dir
	files  map[string]File
}

type File struct {
	size int
}

type FileSystem struct {
	currentDir *Dir
	root       *Dir
}

func main() {
	lines := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func createFS(lines []string) Dir {
	root := Dir{
		parent: nil,
		dirs:   make(map[string]Dir),
		files:  make(map[string]File),
	}
	fs := FileSystem{
		currentDir: &root,
		root:       &root,
	}

	for i, line := range lines {
		if strings.HasPrefix(line, "$ ls") {
			output := getOutput(lines, i+1)
			ls(output, &fs)
		} else if strings.HasPrefix(line, "$ cd") {
			split := strings.Split(line, " ")
			cd(split[2], &fs)
		}
	}

	return *fs.root
}

func cd(path string, fs *FileSystem) {
	switch path {
	case "/":
		fs.currentDir = fs.root
	case "..":
		fs.currentDir = fs.currentDir.parent
	default:
		dir := fs.currentDir.dirs[path]
		fs.currentDir = &dir
	}
}

func ls(output []string, fs *FileSystem) {
	for _, s := range output {
		isDir, name, size := readLSResult(s)
		if _, ok := fs.currentDir.dirs[name]; isDir && !ok {
			fs.currentDir.dirs[name] = Dir{
				parent: fs.currentDir,
				dirs:   make(map[string]Dir),
				files:  make(map[string]File),
			}
		} else if _, ok := fs.currentDir.files[name]; !isDir && !ok {
			fs.currentDir.files[name] = File{size: size}
		}
	}
}

func readLSResult(line string) (bool, string, int) {
	split := strings.Split(line, " ")

	if split[0] == "dir" {
		return true, split[1], 0
	}

	size, _ := strconv.Atoi(split[0])
	return false, split[1], size
}

func getOutput(lines []string, index int) (output []string) {
	for i := index; i < len(lines); i++ {
		if lines[i][0] == '$' {
			return output
		}
		output = append(output, lines[i])
	}
	return output
}

func sumDirSizes(dir Dir, maxSize int, sum *int64) (size int) {
	for _, v := range dir.files {
		size += v.size
	}

	for _, v := range dir.dirs {
		size += sumDirSizes(v, maxSize, sum)
	}

	if sum != nil && size <= maxSize {
		*sum += int64(size)
	}
	return size
}

func findDirToDelete(dir Dir, spaceToFree int, freedSpace *int) (size int) {
	for _, v := range dir.files {
		size += v.size
	}

	for _, v := range dir.dirs {
		size += findDirToDelete(v, spaceToFree, freedSpace)
	}

	if size > spaceToFree && size < *freedSpace {
		*freedSpace = size
	}
	return size
}

func part1(lines []string) (sum int64) {
	root := createFS(lines)
	sumDirSizes(root, 100_000, &sum)
	return sum
}

func part2(lines []string) int {
	root := createFS(lines)

	totalSpace := 70_000_000
	neededSpace := 30_000_000
	usedSpace := sumDirSizes(root, 0, nil)
	unusedSpace := totalSpace - usedSpace
	freedSpace := usedSpace
	findDirToDelete(root, neededSpace-unusedSpace, &freedSpace)

	return freedSpace
}

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "whether or not this should be a real or dry run")
	flag.Parse()
	toRename := make(map[string][]string)
	walkDir := "sample"
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}
		curDir := filepath.Dir(path)
		if m, err := match(info.Name()); err == nil {
			key := filepath.Join(curDir, fmt.Sprintf("%s.%s", m.base, m.ext))

			toRename[key] = append(toRename[key], info.Name())
		}

		return nil
	})

	for key, files := range toRename {
		dir := filepath.Dir(key)
		n := len(files)
		sort.Strings(files)
		for i, fileName := range files {
			res, _ := match(fileName)
			newFilename := fmt.Sprintf("%s - %d of %d.%s", res.base, (i + 1), n, res.ext)
			oldPath := filepath.Join(dir, fileName)
			newPath := filepath.Join(dir, newFilename)
			fmt.Printf("mv %s => %s\n", oldPath, newPath)
			if !dry {
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Println("Error renaming:", oldPath, err.Error())
				}
			}
		}
	}

}

type matchResult struct {
	base  string
	index int
	ext   string
}

// match returns the new file name, or an error if the file name
// didn't match our pattern.
func match(fileName string) (*matchResult, error) {
	//"birthday", "001", "txt"
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	temp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(temp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return nil, fmt.Errorf("%s didn't match our pattern", fileName)
	}
	return &matchResult{
		strings.Title(name),
		number,
		ext,
	}, nil
	//return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil
}

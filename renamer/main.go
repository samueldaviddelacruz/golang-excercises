package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type file struct {
	path string
	name string
}

func main() {

	dir := "sample"
	var toRename []file
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, file{
				name: info.Name(),
				path: path,
			})
		}

		return nil
	})
	for _, f := range toRename {
		fmt.Printf("%q\n", f)
	}
	for _, orig := range toRename {
		var n file
		var err error
		n.name, err = match(orig.name)
		if err != nil {
			fmt.Println("Error matching:", orig.path, err.Error())
		}
		n.path = filepath.Join(dir, n.name)

		fmt.Printf("mv %s => %s\n", orig.path, n.path)
		err = os.Rename(orig.path, n.path)
		if err != nil {
			fmt.Println("Error renaming:", orig.path, err.Error())
		}
	}

	/*
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			panic(err)
		}
		count := 0

		var toRename []string
		for _, file := range files {
			if file.IsDir() {
				fmt.Println("Dir: ", file.Name())
			} else {
				_, err := match(file.Name(), 0)
				if err == nil {
					count++
					toRename = append(toRename, file.Name())
				}
			}
		}

		for _, origFilename := range toRename {
			origPath := filepath.Join(dir, origFilename)
			newFilename, err := match(origFilename, count)
			if err != nil {
				panic(err)
			}
			newPath := filepath.Join(dir, newFilename)
			fmt.Printf("mv %s => %s\n", origPath, newPath)
			err = os.Rename(origPath, newPath)
			if err != nil {
				panic(err)
			}
		}
	*/

}

// match returns the new file name, or an error if the file name
// didn't match our pattern.
func match(fileName string) (string, error) {
	//"birthday", "001", "txt"
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	temp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(temp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern", fileName)
	}
	return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil
}

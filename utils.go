package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func getFolderContent(folderName string) []string {
	var files []string
	err := filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		files = append(files, filepath.Base(path))
		return nil
	})
	if err != nil {
		panic(err)
	}
	files = files[1:]
	return files
}

func getFileContentInArray(fileName string) []string {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(fileBytes), "\n")
}

func convertToCheat(tldrLines []string) string {
	var cheatLines []string

	for index, line := range tldrLines {
		if !strings.HasPrefix(line, ">") &&
			!(index == 0 && strings.HasPrefix(line, "#")) &&
			!(tldrLines[index] == "" && index > 0 && tldrLines[index-1] == "") {

			line = strings.Replace(line, "`", "", -1)
			line = strings.Replace(line, "`", "", -1)
			line = strings.Replace(line, "{{", "", -1)
			line = strings.Replace(line, "}}", "", -1)

			cheatLines = append(cheatLines, line)
		}
	}

	var tmp []string
	for index, line := range cheatLines {
		if !(index > 0 && line == "" && strings.HasPrefix(cheatLines[index-1], "- ")) {
			tmp = append(tmp, line)
		}
	}

	return strings.Trim(strings.Join(tmp, "\n"), "\n")
}

func xor(a bool, b bool) bool {
	return ((a || b) && !(a && b))
}

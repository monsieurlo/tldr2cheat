package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	inFolderFlag  string
	outFolderFlag string
	fileFlag      string
)

func init() {
	flag.StringVar(&inFolderFlag, "infolder", "", "Input folder to use")
	flag.StringVar(&outFolderFlag, "outfolder", "", "Output folder to use")
	flag.StringVar(&fileFlag, "file", "", "File to convert")

	flag.Parse()

	// Ensure file or infile + outfile flag(s) has/have been provided
	if len(fileFlag) == 0 && len(inFolderFlag) == 0 && len(outFolderFlag) == 0 {
		println("\"file\" or \"infile\" + \"outfolderflag\" must be provided")
		os.Exit(255)
	}

	// Ensure infile and/or outfile haven't been provided if file has also been provided
	if (len(inFolderFlag) > 0 || len(outFolderFlag) > 0) && len(fileFlag) > 0 {
		println("\"file\" flag cannot be used at the same time than \"infile\" + \"outfile\"")
		os.Exit(255)
	}

	// infile and outfile must both be provided, not only one of them
	if xor((len(inFolderFlag) > 0), (len(outFolderFlag) > 0)) {
		println("Both \"infolder\" and \"outfolder\" must be provided")
		os.Exit(255)
	}

	// if set, ensure infolder is a folder
	if len(inFolderFlag) > 0 {
		if fi, err := os.Stat(inFolderFlag); err != nil || !fi.Mode().IsDir() {
			println("\"infolder\" must be a directory")
			os.Exit(255)
		}
	}

	// if set, ensure outfolder is a folder
	if len(outFolderFlag) > 0 {
		if fi, err := os.Stat(outFolderFlag); err != nil || !fi.Mode().IsDir() {
			println("\"outfolder\" must be a directory")
			os.Exit(255)
		}
	}

	// Ensure file is a regular file
	if len(fileFlag) > 0 {
		if fi, err := os.Stat(fileFlag); err != nil || fi.Mode().IsDir() {
			println("\"file\" must be a file (not a folder)")
			os.Exit(255)
		}
	}
}

func main() {
	if len(fileFlag) > 0 {
		// Convert a single file and output in on STDOUT
		tldrLines := getFileContentInArray(fileFlag)
		cheat := convertToCheat(tldrLines)
		fmt.Println(cheat)
	} else {
		// Convert all the file in infile and output converted files in outfolder
		cpt := 0
		for _, file := range getFolderContent(inFolderFlag) {
			tldrLines := getFileContentInArray(fmt.Sprintf("%s/%s", inFolderFlag, file))
			cheat := convertToCheat(tldrLines)

			bytesContent := []byte(cheat)
			err := ioutil.WriteFile(fmt.Sprintf("%s/%s", outFolderFlag, file), bytesContent, 0644)
			if err != nil {
				panic(err)
			} else {
				cpt += 1
			}
		}
		fmt.Printf("%d files converted\n", cpt)
	}
}

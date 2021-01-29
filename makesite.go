package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// TemplateDir path to templates
const TemplateDir = "./templates/*.tmpl"

// OutputDir path to output files
const OutputDir = "./output/"

// DefaultContentDir default name of directory at root level to source content textfiles from
const DefaultContentDir = "content"

func main() {
	//FLAGS
	fileNamePtr := flag.String("file", "", " Name of text file to generate HTML for")
	contentDirPtr := flag.String("dir", DefaultContentDir, " Name of folder to generate HTML from")
	flag.Parse()

	contentDirPath := "./" + *contentDirPtr + "/"

	if *fileNamePtr == "" { //if no file is specified we print all from content directory
		files, err := ioutil.ReadDir(contentDirPath)
		if err != nil {
			panic(err)
		}

		for _, file := range files { //for each file we render it as a post
			fmt.Printf("Rendering: %g \n", file.Name()) //TODO: Fix string formatting error
			filePath := contentDirPath + file.Name()
			firstPost := NewBlogPost(filePath)
			name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) //trims file suffix
			firstPost.Render(TemplateDir, OutputDir+name+".html")
		}

	} else { //otherwise we print file specified
		filePath := contentDirPath + *fileNamePtr + ".txt"
		firstPost := NewBlogPost(filePath)
		firstPost.Render(TemplateDir, OutputDir+*fileNamePtr+".html")
	}

}

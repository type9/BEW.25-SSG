package main

import (
	"flag"
	"os"
)

// TemplateDir path to templates
const TemplateDir = "./templates/*.tmpl"

// OutputDir path to output files
const OutputDir = "./output/"

// ContentDir path to content files
const ContentDir = "./content/"

func main() {
	//FLAGS
	fileNamePtr := flag.String("file", "", " Name of text file to generate HTML for")
	flag.Parse()

	if *fileNamePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	filePath := ContentDir + *fileNamePtr + ".txt"

	firstPost := NewBlogPost(filePath)
	firstPost.Render(TemplateDir, OutputDir+*fileNamePtr+".html")
}

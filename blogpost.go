package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

// BlogPost struct for content rendering
type BlogPost struct {
	Content      string
	templateName string
}

// NewBlogPost Constructor
func NewBlogPost(contentPath string) (blogPost BlogPost) {
	const TemplateName = "blogPost.tmpl"
	blogPost = BlogPost{}

	blogPost.templateName = TemplateName
	blogPost.Content = readContent(contentPath)

	return
}

func readContent(filePath string) string {
	blogContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(blogContents)
}

//Render takes a generic template directory and outputs using object defined template name
func (blogPost BlogPost) Render(templateDir, outputDir string) {
	file, err := os.Create(outputDir)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New(blogPost.templateName).ParseGlob(templateDir))
	err = t.Execute(file, blogPost)
	if err != nil {
		panic(err)
	}
}

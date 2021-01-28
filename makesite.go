package main

// TemplateDir path to templates
const TemplateDir = "./templates/*.tmpl"

// OutputDir path to output files
const OutputDir = "./output/"

func main() {
	firstPost := NewBlogPost("./content/first-post.txt")
	firstPost.Render(TemplateDir, OutputDir+"first-post.html")
}

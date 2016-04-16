package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/russross/blackfriday"
)

// This program grabs all of the markdown files used in the report and concatenates them into a single file for postprocessing.
func main() {

	log.Println("Opening files...")

	filenames := FetchFileNames()
	files := ReadFiles(filenames)
	output := Concat(files)
	DumpToFile(output, "output.md")
	RenderToHTML(output, "output.html")
	RenderToLatex(output, "output.tex", "latex/final_report.tex")
}

func FetchFileNames() []string {
	log.Println("Fetching file names...")
	return []string{
		"introduction.md",
		"background.md",
		"approach.md",
		"experimental_setup.md",
		"future_and_related_works.md",
		"conclusions.md",
	}
}

func ReadFiles(files []string) []string {

	log.Println("Opening and reading the files...")
	fileData := make([]string, 0)

	for _, name := range files {
		content, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}
		fileData = append(fileData, string(content))
	}
	return fileData
}

func Concat(content []string) string {

	var buffer bytes.Buffer
	log.Println("Concatenating the files into one entity...")

	for _, section := range content {
		buffer.WriteString(section)
		buffer.WriteString("\n\n-----------------------------\n\n")
	}

	return buffer.String()
}

func DumpToFile(output, filename string) {

	log.Println("Dumping the markdown output to a file...")
	err := ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success!")
}

func RenderToHTML(content string, filenames ...string) {
	output := blackfriday.MarkdownCommon([]byte(content))

	log.Println("Rendering the HTML...")
	for _, filename := range filenames {
		err := ioutil.WriteFile(filename, output, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Success!")
}
func RenderToLatex(content string, filenames ...string) {
	renderer := blackfriday.LatexRenderer(0)
	output := blackfriday.Markdown([]byte(content), renderer, 0)

	log.Println("Rendering the LaTeX...")
	for _, filename := range filenames {
		err := ioutil.WriteFile(filename, output, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Success!")
}

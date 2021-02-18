// https://blog.gopheracademy.com/advent-2017/using-go-templates/
// https://gobyexample.com/panic
// https://gobyexample.com/writing-files
// https://gobyexample.com/reading-files
// https://golang.org/pkg/flag/
// Reference to Dani's makesite.go : https://gist.github.com/droxey/5984bf42810ad53f03b9c465e1484449

package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	// https://github.com/mind1949/googletrans

	"cloud.google.com/go/translate"
	"golang.org/x/text/language" // List of languages https://pkg.go.dev/golang.org/x/text/language
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(fileContents)
}

func createPageFromTextFile(filePath string) Page {
	// Make sure we can read in the file first!
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	translatedFileContents, err := translateText("es", string(fileContents)) // Translate the fileContents into Spanish
	if err != nil {
		panic(err)
	}

	// Get the name of the file without `.txt` at the end.
	// We'll use this later when naming our new HTML file.
	fileNameWithoutExtension := strings.Split(filePath, ".txt")[0]

	// Instantiate a new Page.
	// Populate each field and return the data.
	return Page{
		TextFilePath: filePath,
		TextFileName: fileNameWithoutExtension,
		HTMLPagePath: fileNameWithoutExtension + ".html",
		Content:      translatedFileContents, // Adding the translated text into the Page
	}
}

func renderTemplateFromPage(templateFilePath string, page Page) {
	// Create a new template in memory named "template.tmpl".
	// When the template is executed, it will parse template.tmpl,
	// looking for {{ }} where we can inject content.
	t := template.Must(template.New(templateFilePath).ParseFiles(templateFilePath))

	// Create a new, blank HTML file.
	newFile, err := os.Create(page.HTMLPagePath)
	if err != nil {
		panic(err)
	}

	// Executing the template injects the Page instance's data,
	// allowing us to render the content of our text file.
	// Furthermore, upon execution, the rendered template will be
	// saved inside the new file we created earlier.
	t.Execute(newFile, page)
	fmt.Println("✅ Generated File: ", page.HTMLPagePath)
}

// Translate text into new language
func translateText(targetLanguage, text string) (string, error) {
	// text := "The Go Gopher is cute"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}

func main() {
	// This flag represents the name of any `.txt` file in the same directory as your program.
	// Run `./makesite --file=latest-post.txt` to test.
	var textFilePath string
	flag.StringVar(&textFilePath, "file", "", "Name or Path to a text file")

	var dir string
	flag.StringVar(&dir, "dir", "", "Directory of files")

	flag.Parse()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	// Make sure the `file` flag isn't blank.
	if textFilePath == "" {
		panic("Missing the --file flag! Please supply one.")
	}

	// Read the provided text file and store it's information in a struct.
	newPage := createPageFromTextFile(textFilePath)

	// Use the struct to generate a new HTML page based on the provided template.
	renderTemplateFromPage("template.tmpl", newPage)
}

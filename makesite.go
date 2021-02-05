// https://blog.gopheracademy.com/advent-2017/using-go-templates/
// https://gobyexample.com/panic
// https://gobyexample.com/writing-files
// https://gobyexample.com/reading-files
// https://golang.org/pkg/flag/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func main() {

	// Add a new flag to your command named file. This flag represents the name of any .txt file in the same directory as your program.
	textFileFlag := flag.String("file", "", "Name of the .txt file")
	flag.Parse()

	save(textFileFlag)
}

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	return string(fileContents)
}

func save(textFileFlag *string) {
	textFile, err := ioutil.ReadFile(*textFileFlag)
	newTextFile := strings.Split(*textFileFlag, ".")[0]

	// convert the text file to HTML
	newHTMLFile, err := os.Create(fmt.Sprintf("%s.html", newTextFile))
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(newHTMLFile, textFile)
	if err != nil {
		panic(err)
	}
}

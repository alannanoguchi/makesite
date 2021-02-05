// https://blog.gopheracademy.com/advent-2017/using-go-templates/
// https://gobyexample.com/panic
// https://gobyexample.com/writing-files
// https://gobyexample.com/reading-files

package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

func main() {
	// the Must function is used to verify that a template is valid during parsing
	// .New("new_template_name") creates a new template with it's new name
	// .ParseFiles() allows you to parse through the newly created template
	// (Dani) looking for {{ }}  in .tmpl file where where we can inject content.
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// Reading a File:
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	// .Execute() injects the Page instance's data, allowing us to render the content of the text file //
	// it also saves the new file
	t.Execute(os.Stdout, string(fileContents))
	if err != nil {
		panic(err)
	}

	// Create a new, blank HTML file with panic option
	newFile, err := os.Create("first-post.html")
	err = t.Execute(newFile, string(fileContents))
	if err != nil {
		panic(err)
	}

}

// Writing a File:

// func main() {
// 	bytesToWrite := []byte("hello\ngo\n") // creates 2 new lines followed by an empty line
// 	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// }

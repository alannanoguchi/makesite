package main

// Translate text: https://cloud.google.com/translate/docs/basic/translating-text#translate_translate_text-go
import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"github.com/mind1949/googletrans" // https://github.com/mind1949/googletrans
	"golang.org/x/text/language"      // List of languages https://pkg.go.dev/golang.org/x/text/language
)

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

// Detect language: https://github.com/GoogleCloudPlatform/golang-samples/blob/master/translate/detect.go
func detectLang(fileName string) {
	detected, err := googletrans.Detect(fileName)
	if err != nil {
		panic(err)
	}

	format := "language: %q, confidence: %0.2f\n"
	fmt.Printf(format, detected.Lang, detected.Confidence)
}

// Translate from English to Spanish
func translation(fileName string) {
	params := googletrans.TranslateParams{
		Src:  "auto",
		Dest: language.Spanish.String(),
		Text: fileName,
	}
	translated, err := googletrans.Translate(params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("text: %q \npronunciation: %q", translated.Text, translated.Pronunciation)
}

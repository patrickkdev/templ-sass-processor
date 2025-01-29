package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var styleRegex = regexp.MustCompile(`(?s)<style>(.*?)</style>`)

func main() {
	projectPath := "."
	if len(os.Args) > 1 {
		projectPath = os.Args[1]
	}

	err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".templ") {
			processFile(path)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path: %v", err)
	}
}

func processFile(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read file %s: %v", filePath, err)
		return
	}

	originalText := string(content)
	modifiedText := styleRegex.ReplaceAllStringFunc(originalText, func(match string) string {
		cssContent := styleRegex.FindStringSubmatch(match)
		if len(cssContent) < 2 {
			return match
		}

		nestedCSS := strings.TrimSpace(cssContent[1])
		formattedCSS, err := convertSass(nestedCSS)
		if err != nil {
			log.Printf("Failed to convert SASS for %s: %v", filePath, err)
			return match
		}

		return fmt.Sprintf("<style>%s</style>", formattedCSS)
	})

	if originalText != modifiedText {
		err = ioutil.WriteFile(filePath, []byte(modifiedText), 0644)
		if err != nil {
			log.Printf("Failed to write file %s: %v", filePath, err)
		}
	}
}

func convertSass(nestedCSS string) (string, error) {
	cmd := exec.Command("sass", "--stdin", "--style", "expanded")
	cmd.Stdin = strings.NewReader(nestedCSS)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

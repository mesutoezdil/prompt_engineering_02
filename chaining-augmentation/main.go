package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

// characterTextSplitter splits the text into chunks of specified size with overlap.
func characterTextSplitter(text string, splitSize int, overlapSize int) []string {
	chunks := []string{}
	tokens := strings.Fields(text) // Use Fields to split on whitespace
	for i := 0; i < len(tokens); i += splitSize - overlapSize {
		end := i + splitSize
		if end > len(tokens) {
			end = len(tokens)
		}
		chunks = append(chunks, strings.Join(tokens[i:end], " "))
	}
	return chunks
}

// cleanMarkdown removes unnecessary parts from the markdown.
func cleanMarkdown(markdown string) string {
	// Remove HTML-like elements
	re := regexp.MustCompile(`\s*<[^>]+>\s*|\[[^\]]+\]\(.*?\)`)
	cleaned := re.ReplaceAllString(markdown, " ")
	return strings.TrimSpace(cleaned)
}

func main() {
	converter := md.NewConverter("", true, nil)

	// Download the Go contribution guide.
	res, err := http.Get("https://go.dev/doc/contribute")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // Ensure the response body is closed
	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	html := string(content)

	// Convert HTML to markdown.
	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}

	// Split to get the relevant section
	sections := strings.Split(markdown, "# Contribution Guide")
	if len(sections) < 2 {
		log.Fatal("Contribution Guide not found")
	}
	relevantSection := sections[1]

	// Clean up the markdown content
	cleanedContent := cleanMarkdown(relevantSection)

	// Split the cleaned content into chunks
	chunks := characterTextSplitter(cleanedContent, 100, 10)

	// Print out the first chunk and the number of chunks
	if len(chunks) > 0 {
		fmt.Println(chunks[0]) // Print the first chunk
	}
	fmt.Println("")
	fmt.Println(len(chunks)) // Print total chunk count
}

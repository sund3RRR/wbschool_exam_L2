package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

type Wget struct {
	visited map[string]bool
	baseDir string
}

func NewWget(basedir string) *Wget {
	visited := make(map[string]bool)

	return &Wget{
		visited: visited,
		baseDir: basedir,
	}
}

func (wget *Wget) DownloadPage(urlStr string, depth int) error {
	if depth <= 0 {
		return nil
	}

	if wget.visited[urlStr] {
		return nil
	}
	wget.visited[urlStr] = true

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Create directory for saved file
	filePath := filepath.Join(wget.baseDir, parsedURL.Host, parsedURL.Path)
	if filepath.Ext(parsedURL.Path) == "" {
		filePath = filepath.Join(filePath, "index.html")
	}

	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, body, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded: %s\n", urlStr)

	links := wget.extractLinks(body, urlStr)
	for _, link := range links {
		err := wget.DownloadPage(link, depth-1)
		if err != nil {
			fmt.Printf("Error downloading page %s: %v\n", link, err)
		}
	}

	return nil
}

func (wget *Wget) extractLinks(body []byte, baseUrl string) []string {
	var links []string

	tokenizer := html.NewTokenizer(bytes.NewReader(body))
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}

		token := tokenizer.Token()
		if tokenType == html.StartTagToken && (token.Data == "a" || token.Data == "link" || token.Data == "script" || token.Data == "img") {
			for _, attr := range token.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					link := attr.Val

					// Handle relative links
					if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
						link = baseUrl + link
					}

					links = append(links, link)
				}
			}
		}
	}

	return links
}

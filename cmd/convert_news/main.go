package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type NewsArticle struct {
	Title       string
	CoverImage  string
	PublishDate string
	Summary     string
	Content     string
	Status      string
}

func main() {
	file, err := os.Open("/Users/luca/work/ai/luosi/init/news_data.go")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Output file
	outFile, err := os.Create("/Users/luca/work/ai/luosi/init/news_data.sql")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// Write header
	outFile.WriteString("-- News data SQL insert statements\n")
	outFile.WriteString("-- Generated from news_data.go\n\n")

	scanner := bufio.NewScanner(file)
	// Increase buffer size for long lines
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	var (
		inArticle    bool
		title        string
		summary      string
		content      strings.Builder
		coverImage   string
		status       string
		currentMonth string
		currentYear  string
		articles     []NewsArticle
	)

	// Regex patterns
	monthYearRegex := regexp.MustCompile(`^\s*//\s*(January|February|March|April|May|June|July|August|September|October|November|December)\s+(\d{4})`)
	titleRegex := regexp.MustCompile(`"title":\s*"(.+)",`)
	summaryRegex := regexp.MustCompile(`"summary":\s*"(.+)",`)
	contentStartRegex := regexp.MustCompile(`"content":\s*`)
	coverImageRegex := regexp.MustCompile(`"cover_image":\s*"(.+)",`)
	statusRegex := regexp.MustCompile(`"status":\s*(\d+),?`)
	contentEndRegex := regexp.MustCompile("`,$")

	// Month to number mapping
	monthMap := map[string]string{
		"January": "01", "February": "02", "March": "03", "April": "04",
		"May": "05", "June": "06", "July": "07", "August": "08",
		"September": "09", "October": "10", "November": "11", "December": "12",
	}

	for scanner.Scan() {
		line := scanner.Text()

		// Check for month/year comment
		if matches := monthYearRegex.FindStringSubmatch(line); matches != nil {
			currentMonth = monthMap[matches[1]]
			currentYear = matches[2]
			continue
		}

		// Check for article start (title line)
		if !inArticle {
			if matches := titleRegex.FindStringSubmatch(line); matches != nil {
				inArticle = true
				title = matches[1]
				content.Reset()
				continue
			}
		}

		if inArticle {
			// Check for summary
			if matches := summaryRegex.FindStringSubmatch(line); matches != nil {
				summary = matches[1]
				continue
			}

			// Check for content start
			if contentStartRegex.MatchString(line) {
				// Content starts on this line, extract after the backtick
				idx := strings.Index(line, "`")
				if idx != -1 {
					content.WriteString(line[idx+1:])
					content.WriteString("\n")
				}
				continue
			}

			// Check for content end (backtick followed by comma)
			if contentEndRegex.MatchString(line) {
				// Remove the trailing backtick and comma
				contentStr := content.String()
				contentStr = strings.TrimSuffix(contentStr, "\n")
				contentStr = strings.TrimSuffix(contentStr, "`,")

				// Now look for cover_image and status on following lines
				for scanner.Scan() {
					nextLine := scanner.Text()
					if matches := coverImageRegex.FindStringSubmatch(nextLine); matches != nil {
						coverImage = matches[1]
					} else if matches := statusRegex.FindStringSubmatch(nextLine); matches != nil {
						status = matches[1]
					} else if strings.TrimSpace(nextLine) == "}," || strings.TrimSpace(nextLine) == "}" {
						break
					}
				}

				// Generate publish date
				publishDate := fmt.Sprintf("%s-%s-15", currentYear, currentMonth) // Use 15th as default day

				// Escape single quotes
				title = strings.ReplaceAll(title, "'", "''")
				summary = strings.ReplaceAll(summary, "'", "''")
				contentStr = strings.ReplaceAll(contentStr, "'", "''")
				coverImage = strings.ReplaceAll(coverImage, "'", "''")

				// Add to articles slice
				articles = append(articles, NewsArticle{
					Title:       title,
					CoverImage:  coverImage,
					PublishDate: publishDate,
					Summary:     summary,
					Content:     contentStr,
					Status:      status,
				})

				// Reset for next article
				inArticle = false
				title = ""
				summary = ""
				coverImage = ""
				status = ""
				continue
			}

			// Still in content
			content.WriteString(line)
			content.WriteString("\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Write batch INSERT statement
	outFile.WriteString("INSERT INTO news (title, cover_image, publish_date, summary, content, status, created_at, updated_at) VALUES\n")
	for i, article := range articles {
		value := fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', %s, NOW(), NOW())",
			article.Title, article.CoverImage, article.PublishDate, article.Summary, article.Content, article.Status)
		if i < len(articles)-1 {
			outFile.WriteString(value + ",\n")
		} else {
			outFile.WriteString(value + ";\n")
		}
	}

	fmt.Printf("SQL file generated successfully with %d articles in batch insert format.\n", len(articles))
}

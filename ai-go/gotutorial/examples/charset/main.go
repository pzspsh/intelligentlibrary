package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

// detectAndConvertCharset detects the charset of the content and converts it to UTF-8
func detectAndConvertCharset(body []byte) ([]byte, error) {
	// Detect the charset
	// 假设我们有一段以 ISO-8859-1 编码的文本
	// isoText := []byte{/* ... ISO-8859-1 编码的字节 ... */}
	// 检测字符集
	// 假设我们有一段以 ISO-8859-1 编码的文本
	// 创建一个 bytes.Reader 来读取这段文本
	reader := bytes.NewReader(body)
	// 使用 charset.NewReader 创建一个新的读取器，指定字符集为 ISO-8859-1
	charsetReader, err := charset.NewReader(reader, "ISO-8859-1")
	if err != nil {
		log.Fatalf("Failed to create charset reader: %v", err)
	}

	// 读取转换后的内容
	convertedText, err := io.ReadAll(charsetReader)
	if err != nil {
		log.Fatalf("Failed to read converted text: %v", err)
	}
	return convertedText, nil
}

func main() {
	// Fetch the webpage
	resp, err := http.Get("http://example.com")
	if err != nil {
		log.Fatalf("Failed to fetch webpage: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read webpage body: %v", err)
	}

	// Detect and convert the charset to UTF-8
	utf8Body, err := detectAndConvertCharset(body)
	if err != nil {
		log.Fatalf("Failed to convert charset to UTF-8: %v", err)
	}

	// Parse the UTF-8 encoded HTML using goquery
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(utf8Body))
	if err != nil {
		log.Fatalf("Failed to parse UTF-8 encoded HTML: %v", err)
	}

	// Example: Print the title of the page
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		fmt.Println("Title:", s.Text())
	})

	// Example: Print all paragraph texts
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		fmt.Println("Paragraph:", s.Text())
	})
}

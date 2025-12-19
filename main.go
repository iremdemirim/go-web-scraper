package main

import (
	"context"
	"log"
	"os"
	"strings"
	"github.com/chromedp/chromedp"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Hatalı URL Örnek: go run main.go https://google.com")
	}
	url := os.Args[1]
	screenshotFile := "ekrangor.png"
	htmlFile := "site.html"
	linksFile := "links.txt"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var imageBuf []byte
	var htmlContent string
	var links []string

	task := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.FullScreenshot(&imageBuf, 90),
		chromedp.OuterHTML("html", &htmlContent),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(a => a.href)`, &links),
	}

	if err := chromedp.Run(ctx, task); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(screenshotFile, imageBuf, 0644); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(htmlFile, []byte(htmlContent), 0644); err != nil {
		log.Fatal(err)
	}

	linksData := strings.Join(links, "\n")
	if err := os.WriteFile(linksFile, []byte(linksData), 0644); err != nil {
		log.Fatal(err)
	}

	log.Printf("%s adresi için:\n1. ekran görüntüsü: %s\n2. html: %s\n3. uzantılar: %s dosyalarına kaydedildi.", url, screenshotFile, htmlFile, linksFile)
}
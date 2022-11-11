package imgCrawler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/gocolly/colly"
)

func main() {
	var weburl string
	fmt.Printf("Enter the URL: ")
	fmt.Scanf("%s", &weburl)
	if weburl == "" {
		weburl = "https://www.ibon.com.tw/"
	}
	urlParse, _ := url.Parse(weburl)

	c := colly.NewCollector()

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		fmt.Println("Img src: ", e.Attr("src"))

		imgPath := "./img/" + urlParse.Host + "/"
		imgUrl := weburl + e.Attr("src")
		name := path.Base(imgUrl)

		resp, _ := http.Get(imgUrl)
		os.MkdirAll(imgPath, os.ModePerm)
		out, _ := os.Create(imgPath + name)
		io.Copy(out, resp.Body)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})

	c.Visit(weburl)
}

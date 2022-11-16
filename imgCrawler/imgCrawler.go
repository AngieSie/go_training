package imgCrawler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func ImgCrawler() {
	var weburl string
	fmt.Printf("Enter the URL: ")
	fmt.Scanf("%s", &weburl)

	count, imgSrc := CollyWebImg(weburl)
	showImg(count, imgSrc)
}

func CollyWebImg(weburl string) (count int, imgSrc []string) {
	if weburl == "" {
		weburl = "https://www.ibon.com.tw/"
	}

	c := colly.NewCollector()

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		// fmt.Println("Img src: ", e.Attr("src"))

		// urlParse, _ := url.Parse(weburl)
		// imgPath := "./imgCrawler/img/" + urlParse.Host + "/"
		// imgPath := "./img/"
		imgUrl := weburl + e.Attr("src")
		// downloadImg(imgPath, imgUrl)

		imgSrc = append(imgSrc, imgUrl)
		count++
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
	return
}

func downloadImg(imgPath string, imgUrl string) {

	// imgPath := "./imgCrawler/img/" + urlParse.Host + "/"
	// imgUrl := weburl + e.Attr("src")
	name := path.Base(imgUrl)

	resp, _ := http.Get(imgUrl)
	os.MkdirAll(imgPath, os.ModePerm)
	out, _ := os.Create(imgPath + name)
	io.Copy(out, resp.Body)
}

func showImg(count int, imgSrc []string) {
	router := gin.Default()
	router.LoadHTMLGlob("template/html/*")
	router.GET("/getImage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "showImg.html", gin.H{
			"count":  count,
			"imgSrc": imgSrc,
		})
	})
	router.Run()
}

func ShowImg(c *gin.Context, count int, imgSrc []string) {
	// router := gin.Default()
	// router.LoadHTMLGlob("template/*")
	// router.GET("/index", func(c *gin.Context) {
	c.HTML(http.StatusOK, "showImg.html", gin.H{
		"count":  count,
		"imgSrc": imgSrc,
	})
	// })
	// router.Run()
}

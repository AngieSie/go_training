package main

import (
	"errors"
	"net/http"

	"github.com/angiesie/go_training/imgCrawler"
	"github.com/angiesie/go_training/login"
	"github.com/gin-gonic/gin"
)

func main() {
	// imgCrawler.ImgCrawler()
	// login.Login()

	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")
	//設定靜態資源的讀取
	server.Static("/assets", "./template/assets")
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth)
	server.Run()
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginAuth(c *gin.Context) {
	var (
		username string
		password string
	)
	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入使用者名稱"),
		})
		return
	}
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入密碼名稱"),
		})
		return
	}
	if err := login.Auth(username, password); err == nil {
		// c.HTML(http.StatusOK, "login.html", gin.H{
		// 	"success": "登入成功",
		// })
		count, imgSrc := imgCrawler.CollyWebImg("")
		imgCrawler.ShowImg(c, count, imgSrc)
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}
}

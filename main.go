package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// 将gin提供的默认路由设置为路由
	router = gin.Default()

	// 在开始时加载所有的模版
	router.LoadHTMLGlob("templates/*")

	// 首页
	router.GET("/", showIndexPage)

	// 文章
	router.GET("/article/view/:article_id", getArticle)

	// 开启服务
	router.Run()
}
package main

import (
	_ "dingyuBBC/config"
	_ "dingyuBBC/dao"
	"dingyuBBC/handler/block"
	"dingyuBBC/handler/index"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./ui/assets")
	r.LoadHTMLGlob("./ui/admin-*.tmpl")

	registerRoute(r)
	//r.GET("/index", func(c *gin.Context) {
	//	c.HTML(200, "admin-form.tmpl", "")
	//
	//	url.ParseQuery()
	//})

	r.Run(":8080")
}

func registerRoute(r *gin.Engine) {
	r.GET("/index", index.Index)
	r.GET("/block/add/view", block.BlockAddView)
}

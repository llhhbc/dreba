package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


const BaseUrl = "/drebago/v1"

func LoadHandler(r *gin.RouterGroup)  {

	r.Use(cors.Default())

	r.GET("/health", CheckHealth)

	LoadImageHandles(r)
	LoadBlogHandles(r)
}

func CheckHealth(c *gin.Context)  {

	c.Writer.WriteString("ok")
}

package handler

import (
	"github.com/gin-gonic/gin"
	"dreba/models"
	"fmt"
	"github.com/satori/go.uuid"
)

func LoadBlogHandles(r *gin.RouterGroup)  {

	r.POST("/blog", UploadBlog)

	r.GET("/blogs", GetBlog)
	r.GET("/blogs/:uuid", GetBlog)
}

type BlogResult struct {
	Blogs []models.Blog `json:"blogs"`
	Total int `json:"total"`
}


func GetBlog(c *gin.Context)  {

	res := BlogResult{}
	res.Blogs = make([]models.Blog, 0)

	uid := c.Param("uuid")

	db := models.Gdb
	if uid != "" {
		db = db.Where("uuid = ? ", uid)
	}

	err := db.Find(&res.Blogs).Error
	if err != nil {
		c.AbortWithError(500, fmt.Errorf("get blog fail %v. ", err))
		return
	}

	res.Total = len(res.Blogs)

	c.JSON(200, &res)
}

func UploadBlog(c *gin.Context)  {

	blog := models.Blog{}

	err := c.ShouldBind(&blog)
	if err != nil {
		c.AbortWithError(401, fmt.Errorf("invalid post %s. ", err))
		return
	}

	if blog.Uuid == "" {
		uid := uuid.NewV4()
		blog.Uuid = uid.String()

		err = models.Gdb.Create(&blog).Error
		if err != nil {
			c.AbortWithError(500, fmt.Errorf("save blog fail %s. ", err))
			return
		}
		c.JSON(200, gin.H{
			"code" : 200,
		})
		return
	}

	// update
	err = models.Gdb.Model(&blog).Where("uuid = ? ", blog.Uuid).Updates(&blog).Error
	if err != nil {
		c.AbortWithError(500, fmt.Errorf("update blog fail %s. ", err))
		return
	}

	c.JSON(200, gin.H{
		"code" : 200,
	})
}

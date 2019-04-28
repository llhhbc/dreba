package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
	"dreba/models"
	"github.com/satori/go.uuid"
)

func LoadImageHandles(r *gin.RouterGroup)  {

	r.POST("/images", UploadImage)

	r.GET("/images/:uuid", GetImage)
}

func GetImage(c *gin.Context)  {
	imageInfo := models.ImageInfo{}

	uid := c.Param("uuid")

	err := models.Gdb.Find(&imageInfo, " uuid = ? ", uid).Error
	if err != nil {
		c.AbortWithError(401, fmt.Errorf("get image info fail %s. ", err))
		return
	}

	c.Writer.Write(imageInfo.ImageData)
}

func UploadImage(c *gin.Context) {

	f, err := c.MultipartForm()
	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	uploadFile := f.File["upload"][0]


	src, err := uploadFile.Open()
	if err != nil {
		c.AbortWithError(401, err)
		return
	}
	defer src.Close()

	msg, err := ioutil.ReadAll(src)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	uid := uuid.NewV4()

	imageInfo := models.ImageInfo{}

	imageInfo.Uuid = uid.String()
	imageInfo.ImageData = msg
	imageInfo.FileName = uploadFile.Filename

	err = models.Gdb.Create(&imageInfo).Error
	if err != nil {
		c.JSON(202, gin.H{
			"uploaded" : 0,
			"error":  map[string]string {
				"message" : fmt.Sprintf("save image info fail %s. ", err),
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"uploaded" : 1,
		"fileName" : imageInfo.FileName,
		"url" : fmt.Sprintf("%s/images/%s", BaseUrl, imageInfo.Uuid),
	})
}


package models

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"time"
	"strings"
)

var (
	once sync.Once
	Gdb  *gorm.DB
)

func InitDb(connStr string) {
	once.Do(func() {
		db, err := gorm.Open("mysql", connStr)
		if err != nil {
			glog.Fatal("open db fail ", err)
		}
		err = db.DB().Ping()
		if err != nil {
			glog.Fatal("ping db fail ", err)
		}
		db.LogMode(true)
		Gdb = db
	})
}

type DbTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ImageInfo struct {
	Uuid string
	ImageData []byte
	FileName string
	DbTime
}

type Blog struct {
	Uuid string `json:"uuid"`
	Title string `json:"title"`
	SrcTags string `json:"tags" gorm:"-"`
	Tags CustomerJson `json:"-"`
	ContextType string `json:"contextType"`
	Context string `json:"context"`
	DbTime `json:"-"`
}

func (t *Blog) BeforeSave()  {
	tags := strings.Split(t.SrcTags, ",")
	t.Tags.JsonObject = tags
}

func (t *Blog) AfterFind()  {
	t.SrcTags = strings.Replace(string(t.Tags.JsonBytes[1:len(t.Tags.JsonBytes)-1]), "\"", "",-1)
}

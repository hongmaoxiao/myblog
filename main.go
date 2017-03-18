package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var mux = http.NewServeMux()

type Article struct {
	gorm.Model
	Title   string `gorm:"not null;unique"`
	Content string
}

func GetArticlesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ok := db.HasTable(&Article{})
		if !ok {
			fmt.Println("prepare to create table")
			db.CreateTable(&Article{})
		}
		db.AutoMigrate(&Article{})
		var articles []Article
		db.Find(&articles)

		c.JSON(200, gin.H{"articles": articles})
	}
}
func GetSingleArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		id_int, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}
		var article Article
		db.First(&article, id_int)

		c.JSON(200, gin.H{"article": article})
	}
}
func EditArticlesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.PostForm("title")
		content := c.PostForm("content")
		if title == "" || content == "" {
			c.JSON(401, gin.H{"msg": "文章的标题和内容不能为空"})
			return
		}
		newarticle := Article{Title: title, Content: content}
		db.Create(&newarticle)
		c.JSON(200, gin.H{"msg": "add article success"})
	}
}
func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./dist/static", true)))
	db, err := gorm.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	r.POST("/admin", EditArticlesHandler(db))
	r.GET("/articles", GetArticlesHandler(db))
	r.GET("/article/:id", GetSingleArticleHandler(db))
	r.Run(":3000")
}

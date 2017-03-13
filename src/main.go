package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

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
func main() {
	r := gin.New()
	db, err := gorm.Open("sqlite3", "../data.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	r.GET("/api/articles", GetArticlesHandler(db))
	r.GET("/api/article/:id", GetSingleArticleHandler(db))
	r.Run(":3000")
}

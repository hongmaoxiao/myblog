package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
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

type User struct {
	Username string
	Password string
}

type QueryArticle struct {
	ID        uint
	CreatedAt string
	Title     string
	Content   string
	UpdatedAt  string
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
		res := db.Order("created_at desc").Find(&articles)

		if res.Error != nil {
			c.JSON(401, gin.H{"msg": res.Error})
			return
		}
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
		query_article := db.First(&article, id_int)
		if query_article.Error != nil {
			c.JSON(401, gin.H{"msg": query_article.Error})
			return
		}
		var res QueryArticle
		res.ID = article.ID
		res.Title = article.Title
		res.Content = article.Content
		res.CreatedAt = article.CreatedAt.Format("2006-01-02")
		res.UpdatedAt = article.UpdatedAt.Format("2006-01-02 15:04:05")

		var prev Article
		db.Where("id < ?", id_int).Last(&prev)

		var next Article
		db.Where("id > ?", id_int).First(&next)

		c.JSON(200, gin.H{"article": res, "prev": next.ID, "next": prev.ID})
	}
}
func EditArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.PostForm("title")
		content := c.PostForm("content")
		id := c.PostForm("id")
		if id == "" {
			if title == "" || content == "" {
				c.JSON(401, gin.H{"msg": "文章的标题和内容不能为空"})
				return
			}
			newarticle := Article{Title: title, Content: content}
			res := db.Create(&newarticle)
			if res.Error != nil {
				c.JSON(401, gin.H{"msg": res.Error})
				return
			}
			c.JSON(200, gin.H{"msg": "add article success"})
		} else {
			res := db.Model(&Article{}).Where("id = ?", id).Updates(map[string]interface{}{"title": title, "content": content})
			if res.Error != nil {
				c.JSON(401, gin.H{"msg": res.Error})
				return
			}
			c.JSON(200, gin.H{"msg": "update article success"})
		}
	}
}

func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		var user User
		db.Where("username = ? and password = ?", username, password).Find(&user)
		if user.Username == "" || user.Password == "" {
			c.JSON(401, gin.H{"msg": "login fail"})
			return
		}
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"msg": "login success", "count": count})
	}
}
func DeleteArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		id_int, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}
		res := db.Where("id = ?", id_int).Delete(&Article{})
		if res.Error != nil {
			c.JSON(401, gin.H{"msg": res.Error})
			return
		}
		c.JSON(200, gin.H{"msg": "delete article success"})
	}
}
func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./dist/static", true)))
	// store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6397", "", []byte("secret"))
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	db, err := gorm.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	r.GET("/articles", GetArticlesHandler(db))
	r.GET("/article/:id", GetSingleArticleHandler(db))
	r.POST("/login", LoginHandler(db))
	r.POST("/manages", GetArticlesHandler(db))
	r.POST("/edit", EditArticleHandler(db))
	r.POST("/delete/:id", DeleteArticleHandler(db))
	r.Run(":3000")
}

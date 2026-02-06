package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/duckdb"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Subscriber struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type PageData struct {
	Title           string
	ContentTemplate string
	Year            int
	Message         string
	Success         bool
}

func main() {
	db, err := setupDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/", renderPage("welcome", "Welcome"))
	router.GET("/news", renderPage("news", "News"))
	router.GET("/blog", renderPage("blog", "Blog"))
	router.GET("/about", renderPage("about", "About"))
	router.GET("/contact", renderPage("contact", "Contact"))
	router.GET("/signin", renderPage("signin", "Sign In"))
	router.GET("/signup", renderPage("signup", "Sign Up"))

	router.POST("/subscribe", handleSubscribe(db))

	router.NoRoute(func(c *gin.Context) {
		if isHTMX(c) {
			c.HTML(http.StatusNotFound, "notfound", PageData{Title: "Not Found"})
			return
		}
		c.HTML(http.StatusNotFound, "layout", PageData{
			Title:           "Not Found",
			ContentTemplate: "notfound",
			Year:            time.Now().Year(),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}

func setupDB() (*gorm.DB, error) {
	if err := os.MkdirAll("data", 0o755); err != nil {
		return nil, err
	}

	db, err := gorm.Open(duckdb.Open("data/thestreetpost.duckdb"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Subscriber{}); err != nil {
		return nil, err
	}

	return db, nil
}

func renderPage(page, title string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := PageData{
			Title:           title,
			ContentTemplate: page,
			Year:            time.Now().Year(),
		}

		if isHTMX(c) {
			c.HTML(http.StatusOK, page, data)
			return
		}

		c.HTML(http.StatusOK, "layout", data)
	}
}

func handleSubscribe(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := strings.TrimSpace(c.PostForm("email"))
		if email == "" {
			renderSubscribeResult(c, false, "Please provide a valid email address.")
			return
		}

		subscriber := Subscriber{Email: email}
		if err := db.Create(&subscriber).Error; err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "unique") {
				renderSubscribeResult(c, true, "You're already subscribed. Thanks for staying with us!")
				return
			}
			renderSubscribeResult(c, false, "Something went wrong. Please try again soon.")
			return
		}

		renderSubscribeResult(c, true, "Thanks for subscribing! We'll keep you posted.")
	}
}

func renderSubscribeResult(c *gin.Context, success bool, message string) {
	if isHTMX(c) {
		c.HTML(http.StatusOK, "subscribe-result", PageData{
			Message: message,
			Success: success,
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func isHTMX(c *gin.Context) bool {
	return c.GetHeader("HX-Request") == "true"
}

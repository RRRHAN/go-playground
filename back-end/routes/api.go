package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func AddAPIRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	api := rg.Group("/api")

	api.Use(cors.Default())

	api.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server running properly",
		})
	})

	api.GET("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"foo": os.Getenv("foo"),
		})
	})

	api.GET("/wd", func(c *gin.Context) {
		wd, err := os.Getwd()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintln("failed to get wd - %w", err.Error()))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"foo": wd,
		})
	})

	api.Static("/static", "./data/static")

	api.GET("/images", func(c *gin.Context) {

		images := make([]Image, 0)
		err := db.Select(&images, "SELECT * FROM image")
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get image - %s", err.Error()))
			return
		}

		c.JSON(http.StatusOK, images)
	})

	api.POST("/images", func(c *gin.Context) {
		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}
		files := form.File["images"]

		wd, err := os.Getwd()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintln("failed to get wd - %w", err.Error()))
			return
		}

		staticFolder := filepath.Join(wd, "data", "static")
		imageNames := make([]string, 0)

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			filename = uuid.NewString() + filepath.Ext(filename)
			imageNames = append(imageNames, filename)
			fileLoc := filepath.Join(staticFolder, filename)
			if err := c.SaveUploadedFile(file, fileLoc); err != nil {
				c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
				return
			}
		}

		baseQuery := sq.Insert("image").Columns("name")

		for _, image := range imageNames {
			baseQuery = baseQuery.Values(image)
		}

		sql, args, err := baseQuery.ToSql()
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to generate SQL: %s", err.Error())
			return
		}
		_, err = db.Exec(sql, args...)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed Insert To DB: %s", err.Error())
			return
		}

		c.String(http.StatusOK, "sukses bro")

	})
}

type Image struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

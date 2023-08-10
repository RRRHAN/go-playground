package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/RRRHAN/go-playground/back-end/database"
	"github.com/RRRHAN/go-playground/back-end/routes"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//go:embed ui
var UI embed.FS

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	api := r.Group("/")
	routes.AddAPIRoutes(api, db)

	uiFs := EmbedFolder(UI, "ui")

	staticServer := static.Serve("/", uiFs)

	r.Use(staticServer)
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet &&
			!strings.ContainsRune(c.Request.URL.Path, '.') &&
			!strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Request.URL.Path = "/"
			staticServer(c)
		}
	})

	err = r.Run(os.Getenv("host") + ":8080")
	if err != nil {
		log.Fatalf("failed to start Server %v", err)
	}
}

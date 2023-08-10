package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/RRRHAN/go-playground/back-end/database"
	"github.com/RRRHAN/go-playground/back-end/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInsertData(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	api := r.Group("/")
	routes.AddAPIRoutes(api, db)

	_, err = db.Exec("INSERT INTO image (name) VALUES (?)", "raihan.png")
	if err != nil {
		t.Errorf("err insert to db - %s", err.Error())
	}

	request := httptest.NewRequest(http.MethodGet, "/api/images", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	var actualBody []routes.Image
	err = json.Unmarshal(body, &actualBody)
	if err != nil {
		t.Error(err)
	}

	expectedBody := []routes.Image{
		routes.Image{
			Id: 1, Name: "raihan.png",
		},
	}

	assert.Equal(t, expectedBody, actualBody)
}

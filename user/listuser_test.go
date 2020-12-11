package user

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/getalluser", ListAllUser)
	r.GET("/detailuser/:id", DetailUser)
	r.POST("/postuser", InsertUser)
	return r
}

func TestListUser(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getalluser", nil)
	router.ServeHTTP(w, req)

	expected := `{"message":null,"result":{"UserID":1,"UserName":"asdam","UserFullname":"Amir Saddam","UserEmail":"asdam@gmail.com","UserPassword":"123"}}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
func TestDetailUser(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	id := "1"
	req, _ := http.NewRequest("GET", "/detailuser/"+id, nil)

	router.ServeHTTP(w, req)

	expected := `{"message":null,"result":{"UserID":1,"UserName":"asdam","UserFullname":"Amir Saddam","UserEmail":"asdam@gmail.com","UserPassword":"123"}}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestDetailUserNotFound(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	id := "700"
	req, _ := http.NewRequest("GET", "/detailuser/"+id, nil)

	router.ServeHTTP(w, req)

	// expected := `{"message":null,"result":{"UserID":1,"UserName":"asdam","UserFullname":"Amir Saddam","UserEmail":"asdam@gmail.com","UserPassword":"123"}}`
	assert.Equal(t, 404, w.Code)
	// assert.Equal(t, expected, w.Body.String())
}
func TestPostUser(t *testing.T) {
	router := setupRouter()

	var jsonInput = []byte(`{"UserName":"asdam3","UserFullname":"Amir Saddam3","UserEmail":"asdam@gmail.com3","UserPassword":"1233"}`)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/postuser", bytes.NewBuffer(jsonInput))

	router.ServeHTTP(w, req)

	// expected := `{"message": "Success Input User","result": null}`
	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, expected, w.Body.String())
}

package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	Conf "github.com/asdamwongmantap/apigomycms/config"
	Cconnect "github.com/asdamwongmantap/apigomycms/connect"
	Cuser "github.com/asdamwongmantap/apigomycms/user"

	"github.com/rs/cors"

	_ "github.com/asdamwongmantap/apigomycms/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5001
// @BasePath /api/v1/apigomycms
func main() {

	addr, err := Conf.MyPort()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	v1 := router.Group("/api/v1/apigomycms")
	{
		v1.GET("/tesconnect", Cconnect.Tesconnect)
		v1.GET("/getalluser", Cuser.ListAllUser)
		v1.POST("/postuser", Cuser.InsertUser)
		v1.PUT("/putuser/:id", Cuser.UpdateUser)
		v1.DELETE("/deleteuser/:id", Cuser.DeleteUser)
		v1.GET("/detailuser/:id", Cuser.DetailUser)

	}

	url := ginSwagger.URL("http://localhost" + addr + "/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	c := cors.AllowAll()

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(addr, handler))

}

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	Conf "apigomycms/config"
	Cconnect "apigomycms/connect"
	Cuser "apigomycms/user"

	"github.com/rs/cors"
)

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
	c := cors.AllowAll()

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(addr, handler))

}

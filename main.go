package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	Conf "apigomycms/config"
	Cconnect "apigomycms/connect"

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
	}
	c := cors.AllowAll()

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(addr, handler))

}

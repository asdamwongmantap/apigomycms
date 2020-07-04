package connect

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	Conf "apigomycms/config"
)

func Tesconnect(c *gin.Context) {
	var db, errdb = Conf.Connectdb()
	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()
	c.JSON(http.StatusOK, gin.H{"result": "Success Connection"})

}

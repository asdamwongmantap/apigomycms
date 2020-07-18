package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	Conf "apigomycms/config"
	ModelUser "apigomycms/struct"
)

func ListAllUser(c *gin.Context) {
	var db, errdb = Conf.Connectdb()
	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()
	// added 20200718 by asdam
	var ResultUser ModelUser.ShowUser
	// query show user
	var stmtuser = db.QueryRow("SELECT * FROM tbluser").
		Scan(&ResultUser.UserID, &ResultUser.UserName, &ResultUser.UserFullname, &ResultUser.UserEmail, &ResultUser.UserPassword)
	if stmtuser != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": &ResultUser, "message": stmtuser})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": &ResultUser, "message": nil})
	}
	// end execute query
	// end added
}
func InsertUser(c *gin.Context) {
	var db, errdb = Conf.Connectdb()
	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()
	// param input
	var txtuser ModelUser.ReqUser
	c.BindJSON(&txtuser)
	var UserName = txtuser.UserName
	var UserFullname = txtuser.UserFullname
	var UserEmail = txtuser.UserEmail
	var UserPassword = txtuser.UserPassword
	// query show user
	_, err := db.Query("INSERT INTO tbluser (user_name,user_fullname,user_email,user_password) VALUES (?,?,?,?)", UserName, UserFullname, UserEmail, UserPassword)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": err, "message": "Failed Input User"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": nil, "message": "Success Input User"})
	}
	// end execute query
	// end added
}
func UpdateUser(c *gin.Context) {
	var db, errdb = Conf.Connectdb()
	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()
	// param input
	var txtuser ModelUser.ReqUser
	c.BindJSON(&txtuser)
	var UserID = c.Param("id")
	var UserName = txtuser.UserName
	var UserFullname = txtuser.UserFullname
	var UserEmail = txtuser.UserEmail
	var UserPassword = txtuser.UserPassword
	// query show user
	_, err := db.Query("UPDATE tbluser SET user_name = ?,user_fullname = ?,user_email = ?,user_password = ? where user_id = ?", UserName, UserFullname, UserEmail, UserPassword, UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": err, "message": "Failed update User"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": nil, "message": "Success Update User"})
	}
	// end execute query
	// end added
}
func DeleteUser(c *gin.Context) {
	var db, errdb = Conf.Connectdb()
	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()
	// param input
	var UserID = c.Param("id")
	// query show user
	_, err := db.Query("DELETE FROM tbluser WHERE user_id = ?", UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": err, "message": "Failed Delete User"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": nil, "message": "Success Delete User"})
	}
	// end execute query
	// end added
}
func DetailUser(c *gin.Context) {
	var db, errdb = Conf.Connectdb()
	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()
	// added 20200718 by asdam
	var ResultUser ModelUser.ShowUser
	var UserID = c.Param("id")
	// query show user
	var stmtuser = db.QueryRow("SELECT * FROM tbluser WHERE user_id = ?", UserID).
		Scan(&ResultUser.UserID, &ResultUser.UserName, &ResultUser.UserFullname, &ResultUser.UserEmail, &ResultUser.UserPassword)
	if stmtuser != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": &ResultUser, "message": stmtuser})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": &ResultUser, "message": nil})
	}
	// end execute query
	// end added
}

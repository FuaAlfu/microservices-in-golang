package controllers

import(
	"github.com/gin-goinic/gin"
	"net/http"
)

func Ping(c *gin.Context){
	c.String(http.StatusOK(),"pong")
}
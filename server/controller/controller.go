package controller

import (
	"fmt"
	"net/http"
	"os"
	"unitclock/ml-logger/config"

	"github.com/gin-gonic/gin"
)

func Save(ctx *gin.Context) {

	c := config.Config

	path_ignore := c.GetStringSlice("path_ignore")
	log_path := c.GetString("log_path")

	files, err := os.ReadDir(log_path)

	if err != nil {
		fmt.Println("ReadDir err : %s ", err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Pong")
}

func Log(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

package main

import (
	"C"
)
import (
	"fmt"
	"os"
	"unitclock/ml-logger/server/controller"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func main() {

}

var config *viper.Viper

//export Server
func Server() {

	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.GET("/save", controller.Save)
	r.POST("/log", controller.Log)
	r.Run(":30001")
}

//export Init
func Init() {
	go Server()
}

//export Log
func Log() {}

//export Save
func Save(path string) bool {

	zipFile, err := os.Create(fmt.Sprintf("%s.zip", uuid.New()))
	if err != nil {
		return false
	}
	defer zipFile.Close()

	return false
}

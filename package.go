package main

import (
	"C"
)
import (
	_ "unitclock/ml-logger/config"
	"unitclock/ml-logger/server/controller"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	Server()
}

var config *viper.Viper

//export Server
func Server() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.GET("/save", controller.Save)
	r.POST("/log", controller.Log)
	r.POST("/os", controller.Os)
	r.Run(":30001")
}

//export Init
func Init() {
	go Server()
}

// export Os
func Os() {

}

//export Log
func Log() {

}

//export Save
func Save() {

}

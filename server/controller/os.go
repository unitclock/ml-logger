package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OsBinding struct {
	Platform        string `json:"platform"`
	PythonVersion   string `json:"python_version"`
	System          string `json:"system"`
	Architecture    string `json:"architecture"`
	Processor       string `json:"processor"`
	CpuCount        string `json:"cpu_count"`
	CpuLogicalCount string `json:"cpu_logical_count"`
	Memory          string `json:"memory"`
}

func Os(ctx *gin.Context) {

	os := OsBinding{}

	ctx.ShouldBindJSON(&os)

	fmt.Println(os)

	ctx.JSON(http.StatusOK, os)
}

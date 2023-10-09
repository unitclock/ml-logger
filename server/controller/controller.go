package controller

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"unitclock/ml-logger/config"

	"github.com/gin-gonic/gin"
)

func Save(ctx *gin.Context) {

	c := config.Config
	// 目录绝对路径
	root := c.GetString("project_local_path")

	// 运行id
	running_id := ctx.Query("id")

	path_ignore := c.GetStringSlice("path_ignore")

	// 日志存放的相对路径
	log_path := c.GetString("log_path")

	path_ignore = append(path_ignore, log_path)

	log_store_location := log_path + "/" + c.GetString("access_token") + "/" + c.GetString("project_id") + "/" + c.GetString("experiment_id") + "/" + running_id

	e := os.MkdirAll(log_store_location, os.ModePerm)
	if e != nil {
		fmt.Printf("Error 0  : %s \n", e.Error())
	}

	zipFile, e := os.Create(log_path + "/" + running_id + ".zip")
	if e != nil {
		fmt.Printf("Error 1  : %s \n", e.Error())
	}
	defer zipFile.Close()

	achive := zip.NewWriter(zipFile)
	defer achive.Close()

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		header, err := zip.FileInfoHeader(info)

		if err != nil {
			fmt.Printf("Error 2 : %s \n", err.Error())
			return err
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			fmt.Printf("Error 3 : %s \n", err.Error())
			return err
		}
		header.Name = relPath

		if isExcludeFolder(path, path_ignore) {
			return filepath.SkipDir
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := achive.CreateHeader(header)
		if err != nil {
			fmt.Printf("Error 4 : %s \n", err.Error())
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("Error 5 : %s \n", err.Error())
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				fmt.Printf("Error 6 : %s \n", err.Error())
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error 7 : %s ", err.Error())
	}

	ctx.String(http.StatusOK, fmt.Sprintln(http.StatusOK))
}

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Pong")
}

func Log(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

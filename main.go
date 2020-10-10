package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"themis/pkg/utils"
	"time"
)

func main()  {
	r := gin.Default()
	r.GET("/api/time", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"time": time.Now().Format(utils.TimeFormatter)})
	})
	r.Run()
}
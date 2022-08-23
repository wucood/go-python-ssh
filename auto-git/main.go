package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	r.GET("/api/v1/digit-app-h5",digitAppH5)
	r.POST("/api/v1/digit-app-h5",digitAppH5)
	r.Run("0.0.0.0:9000")
}

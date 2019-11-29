package main

import (
	"Experience/handler"
<<<<<<< HEAD
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler.Ping)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("err")
	}
=======

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Any("/ping", handler.Ping)
	r.POST("/api/video", handler.GetVideoList)
	r.Run()
>>>>>>> work_2

}

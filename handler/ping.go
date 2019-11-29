package handler

import (
<<<<<<< HEAD
	"fmt"
	"net/http"
)

//Ping 连通性测试
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pang")
=======
	"github.com/gin-gonic/gin"
)

//Ping 连通性测试
func Ping(c *gin.Context) {
	c.JSON(200, "pang")
>>>>>>> work_2
}

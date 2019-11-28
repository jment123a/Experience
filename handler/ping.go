package handler

import (
	"fmt"
	"net/http"
)

//Ping 连通性测试
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println(123)
}

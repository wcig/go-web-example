package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Ok(c *gin.Context, data interface{}) {
	res := &HttpResponse{
		Code: 0,
		Data: data,
		Msg:  "success",
	}
	c.JSON(http.StatusOK, res)
}

func Error(c *gin.Context, code int, msg string) {
	res := &HttpResponse{
		Code: code,
		Data: nil,
		Msg:  msg,
	}
	c.JSON(http.StatusOK, res)
}

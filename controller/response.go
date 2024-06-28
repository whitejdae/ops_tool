package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code":10001, 	//程序中的错误码
	"msg":xx,		//提示信息
	"data":{},		//数据
}
*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"` // omitempty 没有值就不展示
}

func ResponseError(c *gin.Context, status int, code ResCode) {
	c.JSON(status, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

func ResponseWithMsg(c *gin.Context, status int, code ResCode, msg interface{}) {
	c.JSON(status, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

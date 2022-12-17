package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
用于处理resp
格式固定,便于前端编写
{
	"code": 0, // 程序中的错误码
	"msg": xx,     // 提示信息
	"data": {},    // 数据
}
*/

// 或者原生的gin.H{}
type Resp struct {
	Code respCode    `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"` // 忽略空的字段
}

func RespInternalErr(c *gin.Context, code respCode) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": code,
		"msg":  code.Msg(),
		"data": nil,
	})
	//c.JSON(http.StatusInternalServerError, &Resp{
	//	Code: code,
	//	Msg:  code.Msg(),
	//	Data: nil,
	//})
}

func RespErr(c *gin.Context, code respCode) {
	c.JSON(http.StatusOK, &Resp{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// 手动指定msg
func RespMsg(c *gin.Context, code respCode, msg interface{}) {
	c.JSON(http.StatusOK, &Resp{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func RespOK(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  msg,
		"data": data,
	})
	//c.JSON(http.StatusOK, &Resp{
	//	Code: CodeSuccess,
	//	Msg:  CodeSuccess.Msg(),
	//	Data: data,
	//})
}

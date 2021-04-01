/**
 * @Author: gaoshiyao
 * @Description: response
 * @File:  response
 * @Date: 2021/03/19 17:31
 */

package response

import (
	"github.com/bg-bigdata/bgpkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//发送json响应信息
func SendJsonResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
	c.Abort()
}

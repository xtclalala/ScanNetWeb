package net

import (
	"github.com/gin-gonic/gin"
	"github.com/xtclalala/ScanNetWeb/internal/proError"
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

func Result(status int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Status:  status,
		Result:  data,
		Message: msg,
	})
}

func Ok(c *gin.Context) {
	Result(proError.SUCCESS, map[string]any{}, proError.GetErrorMessage(proError.SUCCESS), c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(proError.SUCCESS, map[string]any{}, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(proError.SUCCESS, data, proError.GetErrorMessage(proError.SUCCESS), c)
}

func Fail(c *gin.Context) {
	Result(proError.ERROR, map[string]any{}, proError.GetErrorMessage(proError.ERROR), c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(proError.ERROR, map[string]any{}, msg, c)
}

func FailWhitStatusAndMessage(status int, msg string, c *gin.Context) {
	Result(status, map[string]any{}, msg, c)
}

func FailWhitStatus(status int, c *gin.Context) {
	Result(status, map[string]any{}, proError.GetErrorMessage(status), c)
}

type PageVO struct {
	Items any   `json:"items"`
	Total int64 `json:"total"`
}

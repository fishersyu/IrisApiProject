package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ResponseOkMessage    = "OK"
	ResponseErrorMessage = "FAIL"
)

type Response struct {
	Code int         `json:"status"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"message"`
}

func Result(code int, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{code, data, msg})
}

func Ok(ctx *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, ResponseOkMessage, ctx)
}

func OkWithMessage(message string, ctx *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, message, ctx)
}

func OkWithData(data interface{}, ctx *gin.Context) {
	Result(http.StatusOK, data, ResponseOkMessage, ctx)
}

func OkWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(http.StatusOK, data, message, ctx)
}

func Fail(ctx *gin.Context) {
	Result(http.StatusBadRequest, map[string]interface{}{}, ResponseErrorMessage, ctx)
}

func UnauthorizedFailWithMessage(message string, ctx *gin.Context) {
	Result(http.StatusUnauthorized, map[string]interface{}{}, message, ctx)
}

func UnauthorizedFailWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(http.StatusUnauthorized, data, message, ctx)
}

func ForbiddenFailWithMessage(message string, ctx *gin.Context) {
	Result(http.StatusForbidden, map[string]interface{}{}, message, ctx)
}

func FailWithMessage(message string, ctx *gin.Context) {
	Result(http.StatusBadRequest, map[string]interface{}{}, message, ctx)
}

func FailWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(http.StatusBadRequest, data, message, ctx)
}

type PageResult struct {
	List     interface{} `json:"list,omitempty"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type BaseResponse struct {
	Id        uint       `json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

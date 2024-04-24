package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ResponseError struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
	Path    string `json:"path"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}

func OkResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "ok",
		"success": "true",
		"error":   nil,
		"data":    data,
	})
}

func OkResponseWithErrors(ctx *gin.Context, data interface{}, errs []map[string]interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":          http.StatusOK,
		"status":        "ok",
		"success":       "true",
		"errorsCounter": len(errs),
		"errors":        errs,
		"data":          data,
	})
}

func OkEmptyResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "ok",
		"success": "true",
		"error":   nil,
		"data":    map[string]interface{}{},
	})
}

func BadRequestResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"status":  "bad request",
		"success": "false",
		"error":   err.Error(),
		"data":    nil,
	})
}

func ValidationErrorResponse(ctx *gin.Context, err error, validationErrs validator.ValidationErrors) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"status":  "bad request",
		"success": "false",
		"error":   err.Error(),
		"errors":  validationErrs,
	})
}

func InternalServerErrorResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"status":  "internal server error",
		"success": "false",
		"error":   err.Error(),
		"data":    nil,
	})
}

func SimpleErrorResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": false,
	})
}

func SimpleOkResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

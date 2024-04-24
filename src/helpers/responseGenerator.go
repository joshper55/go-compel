package helpers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	ginI18n "github.com/gin-contrib/i18n"
)

func ResponseValidateField(context *gin.Context, err error) (int, ResponseError) {

	status := http.StatusBadRequest

	arrayError, errorsConcat := HandleFieldValidations(err)

	response := NewResponseError(
		context, ginI18n.MustGetMessage("ERROR_VALID_PARAMS_MESSAGE")+" ("+errorsConcat+" )",
		int64(status), arrayError,
	)

	return status, response
}

func ResponseBadRequest(context *gin.Context, message string) (int, ResponseError) {

	status := http.StatusBadRequest
	response := NewResponseError(
		context, ginI18n.MustGetMessage(message),
		int64(status), nil,
	)

	return status, response
}

func ResponseGeneral(context *gin.Context, message string, status int) (int, ResponseError) {

	response := NewResponseError(
		context, ginI18n.MustGetMessage(message),
		int64(status), nil,
	)

	return status, response
}

func ResponseGeneralWithStatus(context *gin.Context, message string, status int, value any) (int, ResponseError) {

	response := NewResponseError(
		context, ginI18n.MustGetMessage(message),
		int64(status), value,
	)

	return status, response
}

func ResponseNotFoundValue(context *gin.Context, message string, value any, status int) (int, ResponseError) {
	response := NewResponseError(
		context, ginI18n.MustGetMessage(&i18n.LocalizeConfig{
			MessageID: message,
			TemplateData: map[string]any{
				"value": value,
			},
		}),
		int64(status), nil,
	)

	return status, response
}

func ResponseCreateWithData(context *gin.Context, message string, value any) (int, ResponseError) {
	status := http.StatusCreated
	response := NewResponseError(
		context, ginI18n.MustGetMessage(message),
		int64(status), value,
	)

	return status, response
}

func ResponseWithOutMessage(context *gin.Context, value any, status int) (int, ResponseError) {
	response := NewResponseError(
		context, "",
		int64(status), value,
	)

	return status, response
}

func HandleFieldValidations(err error) ([]ErrorMsg, string) {

	var ve validator.ValidationErrors
	var errorConcat string
	if errors.As(err, &ve) {

		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			errorConcat += fe.Field() + ": " + getErrorMsg(fe) + " \n"
		}

		return out, errorConcat
	}

	return nil, ""
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return ginI18n.MustGetMessage("FIELD_REQUIRED_MESSAGE")
	case "email":
		return ginI18n.MustGetMessage("ERROR_FORMAT_EMAIL_MESSAGE")
	case "lte":
		return ginI18n.MustGetMessage("FIELD_LESS_THAN_MESSAGE") + fe.Param()
	case "gte":
		return ginI18n.MustGetMessage("FIELD_GRE_THAN_MESSAGE") + fe.Param()

	case "oneof":
		return ginI18n.MustGetMessage("FIELD_RESTRICTED_MESSAGE") + fe.Param()

	case "min":
		return ginI18n.MustGetMessage(&i18n.LocalizeConfig{
			MessageID: "MIN_FIELD_ERROR_MESSAGE",
			TemplateData: map[string]string{
				"length": fe.Param(),
			},
		})

	case "max":
		return ginI18n.MustGetMessage(&i18n.LocalizeConfig{
			MessageID: "MAX_FIELD_ERROR_MESSAGE",
			TemplateData: map[string]string{
				"length": fe.Param(),
			},
		})

	case "regexpValidate":
		return ginI18n.MustGetMessage(&i18n.LocalizeConfig{
			MessageID: "ERROR_FORMAT_MESSAGE",
			TemplateData: map[string]string{
				"format": fe.Param(),
			},
		})

	}

	fmt.Println(fe.Tag())
	return "Unknown error"
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewResponseError(contex *gin.Context, message string, status int64, data any) ResponseError {
	response := ResponseError{
		Message: message,
		Code:    status,
		Path:    contex.Request.URL.Path,
		Data:    data,
	}

	return response
}

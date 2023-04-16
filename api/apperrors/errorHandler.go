package apperrors

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(ctx echo.Context, err error) error {
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	return ErrorResponse(ctx, statusCode, appErr.Message, appErr.Err)
}

func ErrorResponse(ctx echo.Context, code int, message string, err error) error {
	errorResponse := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Error   error  `json:"-"`
	}{
		Code:    code,
		Message: message,
		Error:   err,
	}

	return ctx.JSON(code, errorResponse)
}

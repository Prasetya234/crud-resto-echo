package utils

import "github.com/labstack/echo/v4"

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponse(c echo.Context, statusCode int, messages string, data interface{}) error {
	return c.JSON(statusCode, JsonResponse{
		Code:    statusCode,
		Message: messages,
		Data:    data,
	})
}

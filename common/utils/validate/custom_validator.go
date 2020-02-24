package validate

import (
	"fmt"
	"github.com/labstack/echo/v4"
	. "github.com/trungnguyengtbt/core/common/models"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"strings"
)

type CustomValidator struct {
	validator *validator.Validate
}

type ErrorResponse struct {
	Error  bool              `json:"error"` // define whether validate function return any errors or not
	Errors map[string]string `json:"errors"`
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func SetupValidator(e *echo.Echo) {
	config := &validator.Config{
		TagName:      "validate",
		FieldNameTag: "json",
	}
	e.Validator = &CustomValidator{validator: validator.New(config)}
}

func validationErrorToText(e *validator.FieldError) string {
	switch e.Tag {
	case "required":
		return fmt.Sprintf("%s is is required", e.Field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", e.Field, e.Param)
	case "max":
		return fmt.Sprintf("%s must not be longer than %s characters", e.Field, e.Param)
	case "alphanum":
		return fmt.Sprintf("%s must be alphanumeric", e.Field)
	}
	return fmt.Sprintf("%s is not valid", e.Field)
}

func BadRequest(c echo.Context, errors validator.ValidationErrors) {
	errorTexts := make([]string, 0)
	for _, err := range errors {
		errorTexts = append(errorTexts, validationErrorToText(err))
	}

	c.JSON(http.StatusOK, BadRequestResponse(strings.Join(errorTexts, ",")))
}
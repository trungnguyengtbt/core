package ctx

import (
	"github.com/labstack/echo/v4"
	"github.com/trungnguyengtbt/shopping/app/domain/services"
)

type CustomContext struct {
	echo.Context
	Services *services.Services
}

package ctx

import (
	"github.com/labstack/echo"
	"github.com/trungnguyengtbt/shopping/app/domain/services"
)

type CustomContext struct {
	echo.Context
	Services *services.Services
}

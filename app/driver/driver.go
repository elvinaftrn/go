package driver

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
)

type DriverUsecase interface {
	GetAllDrivers(c *gin.Context) ([]app.Driver, *tools.Pagination, int, error)
	GetDetailDriver(c *gin.Context) (*app.Driver, int, error)
	CreateDriver(c *gin.Context) (int, error)
	UpdateDriver(c *gin.Context) (int, error)
	DeleteDriver(c *gin.Context) (int, error)
}

type DriverRepository interface {
	GetAllDriver(ctx context.Context, pagination *tools.Pagination) ([]app.Driver, *tools.Pagination, error)
	GetDetailDriver(ctx context.Context, ID uint) (*app.Driver, error)
	CreateDriver(ctx context.Context, form *app.Driver) error
	UpdateDriver(ctx context.Context, form *app.Driver) error
	DeleteDriver(ctx context.Context, ID uint) error
}

package driverincentive

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
)

type DriverIncentiveUsecase interface {
	GetAllDriverIncentives(c *gin.Context) ([]app.DriverIncentive, *tools.Pagination, int, error)
	GetDetailDriverIncentive(c *gin.Context) (*app.DriverIncentive, int, error)
	CreateDriverIncentive(c *gin.Context) (int, error)
	DeleteDriverIncentive(c *gin.Context) (int, error)
}

type DriverIncentiveRepository interface {
	GetAllDriverIncentive(ctx context.Context, pagination *tools.Pagination) ([]app.DriverIncentive, *tools.Pagination, error)
	GetDetailDriverIncentive(ctx context.Context, ID uint) (*app.DriverIncentive, error)
	CreateDriverIncentive(ctx context.Context, form *app.DriverIncentive) error
	DeleteDriverIncentive(ctx context.Context, ID uint) error
}

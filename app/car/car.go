package car

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
)

type CarUsecase interface {
	GetAllCars(c *gin.Context) ([]app.Car, *tools.Pagination, int, error)
	GetDetailCar(c *gin.Context) (*app.Car, int, error)
	CreateCar(c *gin.Context) (int, error)
	UpdateCar(c *gin.Context) (int, error)
	DeleteCar(c *gin.Context) (int, error)
}

type CarRepository interface {
	GetAllCar(ctx context.Context, pagination *tools.Pagination) ([]app.Car, *tools.Pagination, error)
	GetDetailCar(ctx context.Context, ID uint) (*app.Car, error)
	CreateCar(ctx context.Context, form *app.Car) error
	UpdateCar(ctx context.Context, form *app.Car) error
	DeleteCar(ctx context.Context, ID uint) error
}

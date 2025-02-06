package bookingtype

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
)

type BookingTypeUsecase interface {
	GetAllBookingTypes(c *gin.Context) ([]app.BookingType, *tools.Pagination, int, error)
	GetDetailBookingType(c *gin.Context) (*app.BookingType, int, error)
	CreateBookingType(c *gin.Context) (int, error)
	UpdateBookingType(c *gin.Context) (int, error)
	DeleteBookingType(c *gin.Context) (int, error)
}

type BookingTypeRepository interface {
	GetAllBookingType(ctx context.Context, pagination *tools.Pagination) ([]app.BookingType, *tools.Pagination, error)
	GetDetailBookingType(ctx context.Context, ID uint) (*app.BookingType, error)
	CreateBookingType(ctx context.Context, form *app.BookingType) error
	UpdateBookingType(ctx context.Context, form *app.BookingType) error
	DeleteBookingType(ctx context.Context, ID uint) error
}

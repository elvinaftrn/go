package booking

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
)

type BookingUsecase interface {
	GetAllBookings(c *gin.Context) ([]app.Booking, *tools.Pagination, int, error)
	GetDetailBooking(c *gin.Context) (*app.Booking, int, error)
	CreateBooking(c *gin.Context) (int, error)
	UpdateBooking(c *gin.Context) (int, error)
	DeleteBooking(c *gin.Context) (int, error)
}

type BookingRepository interface {
	GetAllBooking(ctx context.Context, pagination *tools.Pagination) ([]app.Booking, *tools.Pagination, error)
	GetDetailBooking(ctx context.Context, ID uint) (*app.Booking, error)
	CreateBooking(ctx context.Context, form *app.Booking) error
	UpdateBooking(ctx context.Context, form *app.Booking) error
	DeleteBooking(ctx context.Context, ID uint) error
}

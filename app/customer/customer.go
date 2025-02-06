package customer

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
)

type CustomerUsecase interface {
	GetAllCustomers(c *gin.Context) ([]app.Customer, *tools.Pagination, int, error)
	GetDetailCustomer(c *gin.Context) (*app.Customer, int, error)
	CreateCustomer(c *gin.Context) (int, error)
	UpdateCustomer(c *gin.Context) (int, error)
	DeleteCustomer(c *gin.Context) (int, error)
}

type CustomerRepository interface {
	GetAllCustomer(ctx context.Context, pagination *tools.Pagination) ([]app.Customer, *tools.Pagination, error)
	GetDetailCustomer(ctx context.Context, ID uint) (*app.Customer, error)
	CreateCustomer(ctx context.Context, form *app.Customer) error
	UpdateCustomer(ctx context.Context, form *app.Customer) error
	DeleteCustomer(ctx context.Context, ID uint) error
}

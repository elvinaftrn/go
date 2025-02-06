package membership

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"github.com/gin-gonic/gin"
)

type MembershipUsecase interface {
	GetAllMemberships(c *gin.Context) ([]app.Membership, *tools.Pagination, int, error)
	GetDetailMembership(c *gin.Context) (*app.Membership, int, error)
	CreateMembership(c *gin.Context) (int, error)
	UpdateMembership(c *gin.Context) (int, error)
	DeleteMembership(c *gin.Context) (int, error)
}

type MembershipRepository interface {
	GetAllMembership(ctx context.Context, pagination *tools.Pagination) ([]app.Membership, *tools.Pagination, error)
	GetDetailMembership(ctx context.Context, ID uint) (*app.Membership, error)
	CreateMembership(ctx context.Context, form *app.Membership) error
	UpdateMembership(ctx context.Context, form *app.Membership) error
	DeleteMembership(ctx context.Context, ID uint) error
}

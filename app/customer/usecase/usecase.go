package usecase

import (
	"context"
	"net/http"
	"rental/app"
	"rental/app/customer"
	"rental/app/tools"

	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerUsecase struct {
	repo customer.CustomerRepository
	ctx  context.Context
}

func NewCustomerUsecase(repo customer.CustomerRepository, ctx context.Context) *CustomerUsecase {
	return &CustomerUsecase{
		repo: repo,
		ctx:  ctx,
	}
}

func (uc *CustomerUsecase) GetAllCustomers(c *gin.Context) ([]app.Customer, *tools.Pagination, int, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, http.StatusBadRequest, err
	}

	result, pagination, err := uc.repo.GetAllCustomer(uc.ctx, pagination)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	return result, pagination, http.StatusOK, nil
}

func (uc *CustomerUsecase) GetDetailCustomer(c *gin.Context) (*app.Customer, int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	result, err := uc.repo.GetDetailCustomer(uc.ctx, uint(id))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

func (uc *CustomerUsecase) CreateCustomer(c *gin.Context) (int, error) {
	var createCustomer *app.Customer
	err := c.ShouldBindJSON(&createCustomer)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.CreateCustomer(uc.ctx, createCustomer)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (uc *CustomerUsecase) UpdateCustomer(c *gin.Context) (int, error) {
	var updateCustomer *app.Customer
	err := c.ShouldBindJSON(&updateCustomer)
	if err != nil {
		return http.StatusBadRequest, err
	}

	ID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	updateCustomer.ID = uint(ID)

	err = uc.repo.UpdateCustomer(uc.ctx, updateCustomer)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}

func (uc *CustomerUsecase) DeleteCustomer(c *gin.Context) (int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.DeleteCustomer(uc.ctx, uint(id))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

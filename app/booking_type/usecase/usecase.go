package usecase

import (
	"context"
	"net/http"
	"rental/app"
	bookingtype "rental/app/booking_type"
	"rental/app/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingTypeUsecase struct {
	repo bookingtype.BookingTypeRepository
	ctx  context.Context
}

func NewBookingTypeUsecase(repo bookingtype.BookingTypeRepository, ctx context.Context) *BookingTypeUsecase {
	return &BookingTypeUsecase{
		repo: repo,
		ctx:  ctx,
	}
}

func (uc *BookingTypeUsecase) GetAllBookingTypes(c *gin.Context) ([]app.BookingType, *tools.Pagination, int, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, http.StatusBadRequest, err
	}

	result, pagination, err := uc.repo.GetAllBookingType(uc.ctx, pagination)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	return result, pagination, http.StatusOK, nil
}

func (uc *BookingTypeUsecase) GetDetailBookingType(c *gin.Context) (*app.BookingType, int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	result, err := uc.repo.GetDetailBookingType(uc.ctx, uint(id))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

func (uc *BookingTypeUsecase) CreateBookingType(c *gin.Context) (int, error) {
	var createBookingType *app.BookingType
	err := c.ShouldBindJSON(&createBookingType)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.CreateBookingType(uc.ctx, createBookingType)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (uc *BookingTypeUsecase) UpdateBookingType(c *gin.Context) (int, error) {
	var updateBookingType *app.BookingType
	err := c.ShouldBindJSON(&updateBookingType)
	if err != nil {
		return http.StatusBadRequest, err
	}

	ID, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	updateBookingType.ID = uint(ID)

	err = uc.repo.UpdateBookingType(uc.ctx, updateBookingType)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}

func (uc *BookingTypeUsecase) DeleteBookingType(c *gin.Context) (int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.DeleteBookingType(uc.ctx, uint(id))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

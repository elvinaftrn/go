package usecase

import (
	"context"
	"net/http"
	"rental/app"
	"rental/app/booking"
	"rental/app/car"
	driverincentive "rental/app/driver_incentive"
	"rental/app/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DriverIncentiveUsecase struct {
	repo        driverincentive.DriverIncentiveRepository
	repoBooking booking.BookingRepository
	repoCar     car.CarRepository
	ctx         context.Context
}

func NewDriverIncentiveUsecase(repo driverincentive.DriverIncentiveRepository, repoBooking booking.BookingRepository, repoCar car.CarRepository, ctx context.Context) *DriverIncentiveUsecase {
	return &DriverIncentiveUsecase{
		repo:        repo,
		repoBooking: repoBooking,
		repoCar:     repoCar,
		ctx:         ctx,
	}
}

func (uc *DriverIncentiveUsecase) GetAllDriverIncentives(c *gin.Context) ([]app.DriverIncentive, *tools.Pagination, int, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, http.StatusBadRequest, err
	}

	result, pagination, err := uc.repo.GetAllDriverIncentive(uc.ctx, pagination)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	return result, pagination, http.StatusOK, nil
}

func (uc *DriverIncentiveUsecase) GetDetailDriverIncentive(c *gin.Context) (*app.DriverIncentive, int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	result, err := uc.repo.GetDetailDriverIncentive(uc.ctx, uint(id))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

func (uc *DriverIncentiveUsecase) CreateDriverIncentive(c *gin.Context) (int, error) {
	var form *app.DriverIncentive
	err := c.ShouldBindJSON(&form)
	if err != nil {
		return http.StatusBadRequest, err
	}

	booking, err := uc.repoBooking.GetDetailBooking(c, form.BookingID)
	if err != nil {
		return http.StatusBadRequest, err
	}

	car, err := uc.repoCar.GetDetailCar(c, booking.CarID)
	if err != nil {
		return http.StatusBadRequest, err
	}

	createDriverIncentive := app.DriverIncentive{}
	createDriverIncentive.Incentive = countIncentive(booking, car)
	createDriverIncentive.BookingID = form.BookingID

	err = uc.repo.CreateDriverIncentive(uc.ctx, &createDriverIncentive)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (uc *DriverIncentiveUsecase) DeleteDriverIncentive(c *gin.Context) (int, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.DeleteDriverIncentive(uc.ctx, uint(id))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

func countIncentive(booking *app.Booking, car *app.Car) float64 {
	duration := booking.EndRent.Sub(booking.StartRent)
	days := duration.Hours() / 24
	return (days * car.DailyRent) * 5 / 100
}

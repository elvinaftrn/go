package repository

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (r *BookingRepository) GetAllBooking(ctx context.Context, pagination *tools.Pagination) ([]app.Booking, *tools.Pagination, error) {
	var bookings []app.Booking

	if err := r.db.Offset(pagination.Offset).Limit(pagination.Limit).Preload("Customer").Preload("Car").Preload("BookingType").Preload("Driver").Find(&bookings).Error; err != nil {
		return nil, nil, err
	}

	var count int64
	if err := r.db.Model(&app.Booking{}).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	pagination.Count = int(count)

	pagination = tools.Paging(pagination)
	return bookings, pagination, nil
}

func (r *BookingRepository) GetDetailBooking(ctx context.Context, ID uint) (*app.Booking, error) {
	var booking app.Booking
	if err := r.db.Preload("Customer").Preload("Car").Preload("BookingType").Preload("Driver").First(&booking, ID).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *BookingRepository) CreateBooking(ctx context.Context, form *app.Booking) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(form).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *BookingRepository) UpdateBooking(ctx context.Context, form *app.Booking) error {
	var cus *app.Booking

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Model(&cus).Where("id = ?", form.ID).First(&cus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&cus).Where("id = ?", form.ID).Updates(form).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *BookingRepository) DeleteBooking(ctx context.Context, ID uint) error {
	var cus *app.Booking

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&cus).Where("id = ?", ID).First(&cus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&cus).Where("id = ?", ID).Delete(&app.Booking{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

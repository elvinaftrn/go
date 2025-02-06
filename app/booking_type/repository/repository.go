package repository

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"gorm.io/gorm"
)

type BookingTypeRepository struct {
	db *gorm.DB
}

func NewBookingTypeRepository(db *gorm.DB) *BookingTypeRepository {
	return &BookingTypeRepository{
		db: db,
	}
}

func (r *BookingTypeRepository) GetAllBookingType(ctx context.Context, pagination *tools.Pagination) ([]app.BookingType, *tools.Pagination, error) {
	var bookingTypes []app.BookingType

	if err := r.db.Offset(pagination.Offset).Limit(pagination.Limit).Find(&bookingTypes).Error; err != nil {
		return nil, nil, err
	}

	var count int64
	if err := r.db.Model(&app.BookingType{}).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	pagination.Count = int(count)

	pagination = tools.Paging(pagination)
	return bookingTypes, pagination, nil
}

func (r *BookingTypeRepository) GetDetailBookingType(ctx context.Context, ID uint) (*app.BookingType, error) {
	var bookingType app.BookingType
	if err := r.db.First(&bookingType, ID).Error; err != nil {
		return nil, err
	}
	return &bookingType, nil
}

func (r *BookingTypeRepository) CreateBookingType(ctx context.Context, form *app.BookingType) error {
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

func (r *BookingTypeRepository) UpdateBookingType(ctx context.Context, form *app.BookingType) error {
	var cus *app.BookingType

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

func (r *BookingTypeRepository) DeleteBookingType(ctx context.Context, ID uint) error {
	var cus *app.BookingType

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&cus).Where("id = ?", ID).First(&cus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&cus).Where("id = ?", ID).Delete(&app.BookingType{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

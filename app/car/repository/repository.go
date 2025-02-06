package repository

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"gorm.io/gorm"
)

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

func (r *CarRepository) GetAllCar(ctx context.Context, pagination *tools.Pagination) ([]app.Car, *tools.Pagination, error) {
	var cars []app.Car

	if err := r.db.Offset(pagination.Offset).Limit(pagination.Limit).Find(&cars).Error; err != nil {
		return nil, nil, err
	}

	var count int64
	if err := r.db.Model(&app.Car{}).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	pagination.Count = int(count)

	pagination = tools.Paging(pagination)
	return cars, pagination, nil
}

func (r *CarRepository) GetDetailCar(ctx context.Context, ID uint) (*app.Car, error) {
	var car app.Car
	if err := r.db.First(&car, ID).Error; err != nil {
		return nil, err
	}
	return &car, nil
}

func (r *CarRepository) CreateCar(ctx context.Context, form *app.Car) error {
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

func (r *CarRepository) UpdateCar(ctx context.Context, form *app.Car) error {
	var cus *app.Car

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

func (r *CarRepository) DeleteCar(ctx context.Context, ID uint) error {
	var cus *app.Car

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&cus).Where("id = ?", ID).First(&cus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&cus).Where("id = ?", ID).Delete(&app.Car{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

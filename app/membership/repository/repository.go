package repository

import (
	"context"
	"rental/app"
	"rental/app/tools"

	"gorm.io/gorm"
)

type MembershipRepository struct {
	db *gorm.DB
}

func NewMembershipRepository(db *gorm.DB) *MembershipRepository {
	return &MembershipRepository{
		db: db,
	}
}

func (r *MembershipRepository) GetAllMembership(ctx context.Context, pagination *tools.Pagination) ([]app.Membership, *tools.Pagination, error) {
	var memberships []app.Membership

	if err := r.db.Offset(pagination.Offset).Limit(pagination.Limit).Find(&memberships).Error; err != nil {
		return nil, nil, err
	}

	var count int64
	if err := r.db.Model(&app.Membership{}).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	pagination.Count = int(count)

	pagination = tools.Paging(pagination)
	return memberships, pagination, nil
}

func (r *MembershipRepository) GetDetailMembership(ctx context.Context, ID uint) (*app.Membership, error) {
	var membership app.Membership
	if err := r.db.First(&membership, ID).Error; err != nil {
		return nil, err
	}
	return &membership, nil
}

func (r *MembershipRepository) CreateMembership(ctx context.Context, form *app.Membership) error {
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

func (r *MembershipRepository) UpdateMembership(ctx context.Context, form *app.Membership) error {
	var cus *app.Membership

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

func (r *MembershipRepository) DeleteMembership(ctx context.Context, ID uint) error {
	var cus *app.Membership

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&cus).Where("id = ?", ID).First(&cus).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&cus).Where("id = ?", ID).Delete(&app.Membership{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

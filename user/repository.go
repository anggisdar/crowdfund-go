package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	dbgo *gorm.DB
}

func NewRepository(dbgo *gorm.DB) *repository {
	return &repository{dbgo}
}

func (r *repository) Save(user User) (User, error) {
	err := r.dbgo.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
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

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.dbgo.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

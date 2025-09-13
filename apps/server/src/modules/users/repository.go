package users

import (
	"sodnix/apps/server/src/common/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	repository.GenericRepository[User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{
		GenericRepository: repository.NewGenericRepository[User](db),
		db:                db,
	}
}

func (r *userRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.db.First(&user, "email = ?", email).Error
	return &user, err
}

var _ UserRepositoryPort = (*userRepository)(nil)

package userService

import "gorm.io/gorm"

type UserRepository interface {
	GetUsers() ([]Users, error)
	PostUser(user Users) (Users, error)
	PatchUserByID(id uint, users Users) (Users, error)
	DeleteUserByID(id uint) error
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *usersRepository {
	return &usersRepository{db: db}
}
func (ur *usersRepository) GetUsers() ([]Users, error) {
	var users []Users
	err := ur.db.Find(&users).Error
	return users, err

}
func (ur *usersRepository) PostUser(user Users) (Users, error) {
	err := ur.db.Create(&user).Error
	return user, err
}
func (ur *usersRepository) PatchUserByID(id uint, users Users) (Users, error) {
	err := ur.db.Model(&users).Where("id = ?", id).Updates(users).Error
	return users, err
}
func (ur *usersRepository) DeleteUserByID(id uint) error {
	return ur.db.Delete(&Users{}, &id).Error
}

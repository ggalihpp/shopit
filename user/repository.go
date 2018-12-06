package user

import "github.com/jinzhu/gorm"

// Repository is
type Repository struct {
	db *gorm.DB
}

// Connection - Maintain connection
func Connection(database *gorm.DB) Repository {
	return Repository{database}
}

// GetUser - Will return you a user by its username
func (r *Repository) GetUser(username string) (u User, err error) {

	err = r.db.Where("username = ?", username).First(&u).Error

	return
}

// InsertUser - Will create a new user
func (r *Repository) InsertUser(data *User) (err error) {

	hp, err := hashString(data.Password)
	if err != nil {
		return
	}

	data.HashPassword = hp

	err = r.db.Create(data).Error

	return
}

// DeleteUser - Will delete user from database (Soft delete, only update the deleted_at to time.now)
func (r *Repository) DeleteUser(data *User) (err error) {

	err = r.db.Delete(data).Error

	return
}

// IsUserExist -
func (r *Repository) IsUserExist(data *User) (isExist bool) {
	isExist = !r.db.Where(&User{Username: data.Username}).First(&User{}).RecordNotFound()

	return
}

// CheckUserPassword -
func (r *Repository) CheckUserPassword(u *User) (user *User, correct bool, err error) {
	err = r.db.Where(&User{Username: u.Username}).First(&u).Error
	if err != nil {
		return
	}

	user = u
	correct = checkHash(u.Password, u.HashPassword)

	return
}

// InsertAddress - Will create a user address
func (r *Repository) InsertAddress(data *Address) (err error) {

	err = r.db.Create(data).Error

	return
}

// DeleteAddress - Will delete user address from database (Soft delete, only update the deleted_at to time.now)
func (r *Repository) DeleteAddress(data *Address) (err error) {

	err = r.db.Delete(data).Error

	return
}

// InsertBankAccount - Will create a user bank account
func (r *Repository) InsertBankAccount(data *BankAccount) (err error) {

	err = r.db.Create(data).Error

	return
}

// DeleteBankAccount - Will delete user bank account from database (Soft delete, only update the deleted_at to time.now)
func (r *Repository) DeleteBankAccount(data *BankAccount) (err error) {

	err = r.db.Delete(data).Error

	return
}

// InsertOrder - Will create a user order
func (r *Repository) InsertOrder(data *Order) (err error) {

	err = r.db.Create(data).Error

	return
}

// DeleteOrder - Will delete user order from database (Soft delete, only update the deleted_at to time.now)
func (r *Repository) DeleteOrder(data *Order) (err error) {

	err = r.db.Delete(data).Error

	return
}

// InsertOrderItem - Will create a user order item
func (r *Repository) InsertOrderItem(data *OrderItem) (err error) {

	err = r.db.Create(data).Error

	return
}

// DeleteOrderItem - Will delete user order item from database (Soft delete, only update the deleted_at to time.now)
func (r *Repository) DeleteOrderItem(data *OrderItem) (err error) {

	err = r.db.Delete(data).Error

	return
}

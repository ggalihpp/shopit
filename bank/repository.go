package bank

import "github.com/jinzhu/gorm"

// Repository is
type Repository struct {
	db *gorm.DB
}

// Connection - Maintain connection
func Connection(database *gorm.DB) Repository {
	return Repository{database}
}

// Insert - Will insert new data to bank
func (r *Repository) Insert(data *Bank) (err error) {

	err = r.db.Create(data).Error

	return
}

// Delete - Will delete bank data from database (Soft delete, only update the deleted_at to time.now)
func (r *Repository) Delete(data *Bank) (err error) {

	err = r.db.Delete(data).Error

	return
}

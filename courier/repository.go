package courier

import "github.com/jinzhu/gorm"

// Repository is
type Repository struct {
	db *gorm.DB
}

// Connection - Maintain connection
func Connection(database *gorm.DB) Repository {
	return Repository{database}
}

// Insert - Will insert new data to courier
func (r *Repository) Insert(data *Courier) (err error) {

	err = r.db.Create(data).Error

	return
}

// Delete - Will delete courier from database (Soft delete, only update the deleted_at to time.now)
func (r *Repository) Delete(data *Courier) (err error) {

	err = r.db.Delete(data).Error

	return
}

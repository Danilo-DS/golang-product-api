package repository

import (
	"database/sql"
	. "product-api/src/models"

	"gorm.io/gorm"
)

// CategoryRepository is an implementation IRepository
type CategoryRepository struct {
	connection    *sql.DB
	connectionOrm *gorm.DB
}

// Without ORM
// func (c *CategoryRepository) ListAll() ([]Category, error) {
// 	connection := c.connection
// 	defer connection.Close()

// 	result, err := connection.Query("SELECT ID, NAME, ACTIVE FROM CATEGORY")

// 	if err != nil {
// 		return nil, err
// 	}

// 	var categories []Category

// 	for result.Next() {
// 		category := Category{}
// 		if err = result.Scan(&category.ID, &category.Name, &category.Active); err != nil {
// 			return nil, err
// 		}

// 		categories = append(categories, category)
// 	}

// 	return categories, nil
// }

// With ORM
func (c *CategoryRepository) ListAll() ([]Category, error) {

	connection := c.connectionOrm

	var categories []Category

	err := connection.Model(&categories).Scan(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// Without ORM
// func (c *CategoryRepository) GetById(id uint64) (Category, error) {

// 	connection := c.connection
// 	defer connection.Close()

// 	var category Category

// 	result, err := connection.Query("SELECT ID, NAME, ACTIVE FROM CATEGORY WHERE ID = ?", id)

// 	if err != nil {
// 		return category, err
// 	}

// 	if result.Next() {
// 		err = result.Scan(&category.ID, &category.Name, &category.Active)
// 	}

// 	return category, err
// }

// With ORM
func (c *CategoryRepository) GetById(id uint64) (Category, error) {

	connection := c.connectionOrm

	var category Category

	err := connection.Model(&category).Where(&Category{ID: id}).Scan(&category).Error

	if err != nil {
		return category, err
	}

	return category, err
}

// Without ORM
// func (c *CategoryRepository) Save(entity Category) (Category, error) {

// 	connection := c.connection
// 	defer connection.Close()

// 	statement, err := connection.Prepare("INSERT INTO CATEGORY (NAME, ACTIVE, DESCRIPTION) VALUES (?, ?, ?)")

// 	if err != nil {
// 		return Category{}, err
// 	}

// 	result, err := statement.Exec(entity.Name, entity.Active, entity.Description)

// 	if err != nil {
// 		return Category{}, err
// 	}

// 	categoryIdRegisted, err := result.LastInsertId()

// 	if err != nil {
// 		return Category{}, err
// 	}

// 	entity.ID = uint64(categoryIdRegisted)

// 	return entity, nil
// }

// With ORM
func (c *CategoryRepository) Save(entity Category) (Category, error) {

	connection := c.connectionOrm

	err := connection.Create(&entity).Error

	if err != nil {
		return Category{}, err
	}

	return c.GetById(entity.ID)
}

// Without ORM
// func (c *CategoryRepository) Update(entity Category) error {

// 	connection := c.connection
// 	defer connection.Close()

// 	statement, err := connection.Prepare("UPDATE CATEGORY SET NAME = ?, DESCRIPTION = ? WHERE ID = ?")

// 	if err != nil {
// 		return err
// 	}

// 	if _, err = statement.Exec(entity.Name, entity.Description, entity.ID); err != nil {
// 		return err
// 	}

// 	return nil
// }

// With ORM
func (c *CategoryRepository) Update(entity Category) error {

	connection := c.connectionOrm

	err := connection.Save(&entity).Error

	if err != nil {
		return err
	}

	return nil
}

// Without ORM
// func (c *CategoryRepository) DeleteById(id uint64) error {

// 	connection := c.connection
// 	defer connection.Close()

// 	statement, err := connection.Prepare("DELETE FROM CATEGORY WHERE ID = ?")

// 	if err != nil {
// 		return err
// 	}

// 	if _, err = statement.Exec(id); err != nil {
// 		return err
// 	}

// 	return nil
// }

// With ORM
func (c *CategoryRepository) DeleteById(id uint64) error {

	connection := c.connectionOrm

	err := connection.Delete(&Category{ID: id}).Error

	if err != nil {
		return err
	}

	return nil
}

// Without ORM
// func (c *CategoryRepository) EnableCategoryById(id uint64) error {

// 	connection := c.connection
// 	defer connection.Close()

// 	statement, err := connection.Prepare("UPDATE CATEGORY SET ACTIVE = true WHERE ID = ? ")

// 	if err != nil {
// 		return err
// 	}

// 	if _, err = statement.Exec(id); err != nil {
// 		return err
// 	}

// 	return nil
// }

// With ORM
func (c *CategoryRepository) EnableCategoryById(id uint64) error {

	connection := c.connectionOrm

	category, err := c.GetById(id)

	if err != nil {
		return err
	}

	category.Active = true

	err = connection.Save(&category).Error

	if err != nil {
		return err
	}

	return nil
}

// Without ORM
// func (c *CategoryRepository) DisableCategoryById(id uint64) error {

// 	connection := c.connection
// 	defer connection.Close()

// 	statement, err := connection.Prepare("UPDATE CATEGORY SET ACTIVE = false WHERE ID = ? ")

// 	if err != nil {
// 		return err
// 	}

// 	if _, err = statement.Exec(id); err != nil {
// 		return err
// 	}

// 	return nil
// }

// With ORM

func (c *CategoryRepository) DisableCategoryById(id uint64) error {

	connection := c.connectionOrm
	category, err := c.GetById(id)

	if err != nil {
		return err
	}

	category.Active = false

	err = connection.Save(&category).Error

	if err != nil {
		return err
	}

	return nil
}

// Without ORM
// func (c *CategoryRepository) ExistsById(id uint64) (bool, error) {

// 	connection := c.connection
// 	defer connection.Close()

// 	result, err := connection.Query("SELECT EXISTS (SELECT ID FROM CATEGORY WHERE ID = ?)", id)

// 	if err != nil {
// 		return false, err
// 	}

// 	var exists bool

// 	if result.Next() {
// 		result.Scan(&exists)
// 	}

// 	return exists, nil
// }

// With ORM
func (c *CategoryRepository) ExistsById(id uint64) (bool, error) {

	connection := c.connectionOrm

	var exists bool

	err := connection.Model(&Category{}).Select("COUNT(*) > 0").Where(&Category{ID: id}).Scan(&exists).Error

	if err != nil {
		return false, err
	}

	return exists, nil
}

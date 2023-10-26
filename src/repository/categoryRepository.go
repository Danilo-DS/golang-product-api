package repository

import (
	"database/sql"
	. "product-api/src/models"
)

type CategoryRepository struct {
	connection *sql.DB
}

func (c *CategoryRepository) ListAll() ([]Category, error) {
	connection := c.connection
	defer connection.Close()

	result, err := connection.Query("SELECT ID, NAME, ACTIVE FROM CATEGORY")

	if err != nil {
		return nil, err
	}

	var categories []Category

	for result.Next() {
		category := Category{}
		if err = result.Scan(&category.ID, &category.Name, &category.Active); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (c *CategoryRepository) GetById(id uint64) (Category, error) {

	connection := c.connection
	defer connection.Close()

	var category Category

	result, err := connection.Query("SELECT ID, NAME, ACTIVE FROM CATEGORY WHERE ID = ?", id)

	if err != nil {
		return category, err
	}

	if result.Next() {
		err = result.Scan(&category.ID, &category.Name, &category.Active)
	}

	return category, err
}

func (c *CategoryRepository) Save(entity Category) (Category, error) {

	connection := c.connection
	defer connection.Close()

	statement, err := connection.Prepare("INSERT INTO CATEGORY (NAME, ACTIVE, DESCRIPTION) VALUES (?, ?, ?)")

	if err != nil {
		return Category{}, err
	}

	result, err := statement.Exec(entity.Name, entity.Active, entity.Description)

	if err != nil {
		return Category{}, err
	}

	categoryIdRegisted, err := result.LastInsertId()

	if err != nil {
		return Category{}, err
	}

	entity.ID = uint64(categoryIdRegisted)

	return entity, nil
}

func (c *CategoryRepository) Update(entity Category) error {

	connection := c.connection
	defer connection.Close()

	statement, err := connection.Prepare("UPDATE CATEGORY SET NAME = ?, DESCRIPTION = ? WHERE ID = ?")

	if err != nil {
		return err
	}

	if _, err = statement.Exec(entity.Name, entity.Description, entity.ID); err != nil {
		return err
	}

	return nil
}

func (c *CategoryRepository) DeleteById(id uint64) error {

	connection := c.connection
	defer connection.Close()

	statement, err := connection.Prepare("DELETE FROM CATEGORY WHERE ID = ?")

	if err != nil {
		return err
	}

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (c *CategoryRepository) EnableCategoryById(id uint64) error {

	connection := c.connection
	defer connection.Close()

	statement, err := connection.Prepare("UPDATE CATEGORY SET ACTIVE = true WHERE ID = ? ")

	if err != nil {
		return err
	}

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (c *CategoryRepository) DisableCategoryById(id uint64) error {

	connection := c.connection
	defer connection.Close()

	statement, err := connection.Prepare("UPDATE CATEGORY SET ACTIVE = false WHERE ID = ? ")

	if err != nil {
		return err
	}

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (c *CategoryRepository) ExistsById(id uint64) (bool, error) {

	connection := c.connection
	defer connection.Close()

	result, err := connection.Query("SELECT EXISTS (SELECT ID FROM CATEGORY WHERE ID = ?)", id)

	if err != nil {
		return false, err
	}

	var exists bool

	if result.Next() {
		result.Scan(&exists)
	}

	return exists, nil
}

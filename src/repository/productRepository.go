package repository

import (
	"database/sql"
	"fmt"
	. "product-api/src/models"

	"gorm.io/gorm"
)

// ProductRepository is an implementation IRepository
type ProductRepository struct {
	connection    *sql.DB
	connectionOrm *gorm.DB
}

// Without ORM
// func (p *ProductRepository) ListAll() ([]Product, error) {
// 	connection := p.connection
// 	defer connection.Close()

// 	result, err := connection.Query("SELECT P.ID, P.NAME, P.PRICE, P.BARCODE, C.ID, C.NAME, C.ACTIVE FROM PRODUCT P INNER JOIN CATEGORY C ON P.ID_CATEGORY = C.ID")

// 	if err != nil {
// 		return nil, err
// 	}

// 	var products []Product

// 	for result.Next() {
// 		product := Product{}
// 		if err = result.Scan(&product.ID, &product.Name, &product.Price, &product.Barcode, &product.CategoryId, &product.Category.ID, &product.Category.Name, &product.Category.Active, &product.Category.Description); err != nil {
// 			return nil, err
// 		}

// 		products = append(products, product)
// 	}

// 	return products, nil
// }

// With ORM
func (p *ProductRepository) ListAll() ([]Product, error) {

	connection := p.connectionOrm

	var products []Product

	err := connection.Model(&products).InnerJoins("Category").Scan(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

// Without ORM
// func (p *ProductRepository) GetById(id uint64) (Product, error) {

// 	connection := p.connection
// 	defer connection.Close()

// 	var product Product

// 	result, err := connection.Query("SELECT P.ID, P.NAME, P.PRICE, P.BARCODE, C.ID, C.NAME FROM PRODUCT P INNER JOIN CATEGORY C ON P.ID_CATEGORY = C.ID WHERE P.ID = ?", id)

// 	if err != nil {
// 		return product, err
// 	}

// 	if result.Next() {
// 		err = result.Scan(&product.ID, &product.Name, &product.Price, &product.Barcode, &product.Category.ID, &product.Category.Name)
// 	}

// 	return product, err
// }

// With ORM
func (p *ProductRepository) GetById(id uint64) (Product, error) {

	connection := p.connectionOrm

	var product Product

	err := connection.Model(&product).InnerJoins("Category").Where(Product{ID: id}).Scan(&product).Error

	if err != nil {
		return product, err
	}

	return product, err
}

// Without ORM
// func (p *ProductRepository) Save(entity Product) (Product, error) {

// 	connetion := p.connection
// 	defer connetion.Close()

// 	statement, err := connetion.Prepare("INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES (?, ?, ?, ?)")

// 	if err != nil {
// 		return Product{}, err
// 	}

// 	result, err := statement.Exec(entity.Name, entity.Price, entity.Barcode, entity.Category.ID)

// 	if err != nil {
// 		return Product{}, err
// 	}

// 	productIdRegistred, err := result.LastInsertId()

// 	if err != nil {
// 		return Product{}, err
// 	}

// 	entity.ID = uint64(productIdRegistred)

// 	return entity, nil
// }

// With ORM
func (p *ProductRepository) Save(entity Product) (Product, error) {

	connetion := p.connectionOrm

	err := connetion.Create(&entity).Error

	if err != nil {
		return Product{}, err
	}

	return p.GetById(entity.ID)
}

// Without ORM
// func (p *ProductRepository) Update(entity Product) error {

// 	connection := p.connection
// 	defer connection.Close()

// 	statement, err := connection.Prepare("UPDATE PRODUCT SET NAME = ?, PRICE = ?, BARCODE = ?, ID_CATEGORY = ? WHERE ID = ?")

// 	if err != nil {
// 		return err
// 	}

// 	if _, err = statement.Exec(entity.Name, entity.Price, entity.Barcode, entity.Category.ID, entity.ID); err != nil {
// 		return err
// 	}

// 	return nil
// }

// With ORM
func (p *ProductRepository) Update(entity Product) error {

	connection := p.connectionOrm

	err := connection.Save(&entity).Error

	if err != nil {
		return err
	}

	return nil
}

// Without ORM
// func (p *ProductRepository) DeleteById(id uint64) error {

// 	connection := p.connection
// 	defer connection.Close()

// 	statement, err := connection.Prepare("DELETE FROM PRODUCT P WHERE P.ID = ?")

// 	if err != nil {
// 		return err
// 	}

// 	if _, err = statement.Exec(id); err != nil {
// 		return err
// 	}

// 	return nil
// }

// With ORM
func (p *ProductRepository) DeleteById(id uint64) error {

	connection := p.connectionOrm

	err := connection.Delete(&Product{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

// Withour ORM
// func (p *ProductRepository) GetProductsByName(name string) ([]Product, error) {

// 	connection := p.connection
// 	defer connection.Close()

// 	likeName := fmt.Sprintf("%%%s%%", name)

// 	result, err := connection.Query(" SELECT  P.ID, P.NAME, P.PRICE, P.BARCODE, C.ID, C.NAME, C.ACTIVE FROM PRODUCT P INNER JOIN CATEGORY C ON P.ID_CATEGORY = C.ID WHERE P.NAME LIKE ? AND C.ACTIVE = TRUE", likeName)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var produtcs []Product

// 	for result.Next() {
// 		product := Product{}

// 		if result.Scan(&product.ID, &product.Name, &product.Price, &product.Barcode, &product.Category.ID, &product.Category.Name, &product.Category.Active); err != nil {
// 			return nil, err
// 		}

// 		produtcs = append(produtcs, product)
// 	}

// 	return produtcs, nil
// }

// With ORM
func (p *ProductRepository) GetProductsByName(name string) ([]Product, error) {

	connection := p.connectionOrm

	var produtcs []Product

	likeName := fmt.Sprintf("%%%s%%", name)

	err := connection.Model(&produtcs).InnerJoins("Category").Where("PRODUCT.NAME like ?", likeName).Find(&produtcs).Error

	if err != nil {
		return nil, err
	}

	return produtcs, nil
}

// Without orm
// func (p *ProductRepository) ExistsById(id uint64) (bool, error) {

// 	connection := p.connection
// 	defer connection.Close()

// 	result, err := connection.Query("SELECT EXISTS (SELECT ID FROM PRODUCT WHERE ID = ?)", id)

// 	if err != nil {
// 		return false, err
// 	}

// 	var exists bool

// 	if result.Next() {
// 		result.Scan(&exists)
// 	}

// 	fmt.Println("ExistsById", exists)
// 	return exists, nil
// }

// With orm
func (p *ProductRepository) ExistsById(id uint64) (bool, error) {

	connection := p.connectionOrm

	var exists bool
	err := connection.Model(&Product{}).Select("COUNT(*) > 0").Where(&Product{ID: id}).Scan(&exists).Error

	if err != nil {
		return false, err
	}

	fmt.Println("ExistsById", exists)
	return exists, nil
}

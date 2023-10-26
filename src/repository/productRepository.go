package repository

import (
	"database/sql"
	"fmt"
	. "product-api/src/models"
)

type ProductRepository struct {
	connection *sql.DB
}

func (p *ProductRepository) ListAll() ([]Product, error) {
	connection := p.connection
	defer connection.Close()

	result, err := connection.Query("SELECT P.ID, P.NAME, P.PRICE, P.BARCODE, C.ID, C.NAME, C.ACTIVE FROM PRODUCT P INNER JOIN CATEGORY C ON P.ID_CATEGORY = C.ID")

	if err != nil {
		return nil, err
	}

	var products []Product

	for result.Next() {
		product := Product{}
		if err = result.Scan(&product.ID, &product.Name, &product.Price, &product.Barcode, &product.Category.ID, &product.Category.Name, &product.Active); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (p *ProductRepository) GetById(id uint64) (Product, error) {

	connection := p.connection
	defer connection.Close()

	var product Product

	result, err := connection.Query("SELECT P.ID, P.NAME, P.PRICE, P.BARCODE, C.ID, C.NAME FROM PRODUCT P INNER JOIN CATEGORY C ON P.ID_CATEGORY = C.ID WHERE P.ID = ?", id)

	if err != nil {
		return product, err
	}

	if result.Next() {
		err = result.Scan(&product.ID, &product.Name, &product.Price, &product.Barcode, &product.Category.ID, &product.Category.Name)
	}

	return product, err
}

func (p *ProductRepository) Save(entity Product) (Product, error) {

	connetion := p.connection
	defer connetion.Close()

	statement, err := connetion.Prepare("INSERT INTO PRODUCT (NAME, PRICE, BARCODE, ID_CATEGORY) VALUES (?, ?, ?, ?)")

	if err != nil {
		return Product{}, err
	}

	result, err := statement.Exec(entity.Name, entity.Price, entity.Barcode, entity.Category.ID)

	if err != nil {
		return Product{}, err
	}

	productIdRegistred, err := result.LastInsertId()

	if err != nil {
		return Product{}, err
	}

	entity.ID = uint64(productIdRegistred)

	return entity, nil
}

func (p *ProductRepository) Update(entity Product) error {

	connection := p.connection
	defer connection.Close()

	statement, err := connection.Prepare("UPDATE PRODUCT SET NAME = ?, PRICE = ?, BARCODE = ?, ID_CATEGORY = ? WHERE ID = ?")

	if err != nil {
		return err
	}

	if _, err = statement.Exec(entity.Name, entity.Price, entity.Barcode, entity.Category.ID, entity.ID); err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) DeleteById(id uint64) error {

	connection := p.connection
	defer connection.Close()

	statement, err := connection.Prepare("DELETE FROM PRODUCT P WHERE P.ID = ?")

	if err != nil {
		return err
	}

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) GetProductsByName(name string) ([]Product, error) {

	connection := p.connection
	defer connection.Close()

	likeName := fmt.Sprintf("%%%s%%", name)

	result, err := connection.Query(" SELECT  P.ID, P.NAME, P.PRICE, P.BARCODE, C.ID, C.NAME, C.ACTIVE FROM PRODUCT P INNER JOIN CATEGORY C ON P.ID_CATEGORY = C.ID WHERE P.NAME LIKE ? AND C.ACTIVE = TRUE", likeName)

	if err != nil {
		return nil, err
	}

	var produtcs []Product

	for result.Next() {
		product := Product{}

		if result.Scan(&product.ID, &product.Name, &product.Price, &product.Barcode, &product.Category.ID, &product.Category.Name, &product.Active); err != nil {
			return nil, err
		}

		produtcs = append(produtcs, product)
	}

	return produtcs, nil
}

func (p *ProductRepository) ExistsById(id uint64) (bool, error) {

	connection := p.connection
	defer connection.Close()

	result, err := connection.Query("SELECT EXISTS (SELECT ID FROM PRODOCT WHERE ID = ?)", id)

	if err != nil {
		return false, err
	}

	var exists bool

	if result.Next() {
		result.Scan(&exists)
	}

	return exists, nil
}

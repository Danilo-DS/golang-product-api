package services

import (
	"errors"
	. "product-api/src/dto"
	. "product-api/src/mapper"
	. "product-api/src/models"
	. "product-api/src/repository"
)

var productRepository *FactoryMethodImpl[ProductRepository] = InitFactory[ProductRepository]()

func GetAllProducts() ([]ProductResponseDTO, error) {

	repository, err := productRepository.GetRepository()

	if err != nil {
		return nil, err
	}

	products, err := repository.ListAll()

	if err != nil {
		return nil, err
	}

	return ParseResponseSlice[Product, ProductResponseDTO](products)
}

func GetProductById(id uint64) (ProductDetailResponseDTO, error) {

	repository, err := productRepository.GetRepository()

	if err != nil {
		return ProductDetailResponseDTO{}, err
	}

	product, err := repository.GetById(id)

	if err != nil {
		return ProductDetailResponseDTO{}, err
	}

	if (product == Product{}) {
		return ProductDetailResponseDTO{}, errors.New("Product Not Found")
	}

	return ParseResponseStruct[Product, ProductDetailResponseDTO](product)
}

func SearchProductByName(name string) ([]ProductResponseDTO, error) {

	repository, err := productRepository.GetRepository()

	if err != nil {
		return nil, err
	}

	products, err := repository.GetProductsByName(name)

	if err != nil {
		return nil, err
	}

	if len(products) > 0 {
		return ParseResponseSlice[Product, ProductResponseDTO](products)
	}

	return nil, errors.New("No products found")
}

func SaveProduct(request ProductRequestDTO) (ProductResponseDTO, error) {

	product, err := ParseRequestStruct[ProductRequestDTO, Product](request)

	var productRegistred ProductResponseDTO

	if err != nil {
		return productRegistred, err
	}

	repository, err := productRepository.GetRepository()

	if err != nil {
		return ProductResponseDTO{}, err
	}

	product.GenerateBarcode()
	product, err = repository.Save(product)

	if err != nil {
		return productRegistred, err
	}

	return ParseResponseStruct[Product, ProductResponseDTO](product)
}

func UpdateProduct(id uint64, request ProductRequestDTO) error {

	productUpdate, err := ParseRequestStruct[ProductRequestDTO, Product](request)

	if err != nil {
		return err
	}

	repository, err := productRepository.GetRepository()

	if err != nil {
		return err
	}

	product, err := repository.GetById(id)

	if err != nil {
		return err
	}

	if (product == Product{}) {
		return errors.New("Update fial, product not found.")
	}

	productUpdate.ID = id

	if productUpdate.Name != product.Name {
		productUpdate.GenerateBarcode()
	}

	if err = repository.Update(productUpdate); err != nil {
		return err
	}

	return nil
}

func DeleteProductById(id uint64) error {

	repository, err := productRepository.GetRepository()

	if err != nil {
		return err
	}

	existsProduct, err := repository.ExistsById(id)
	if existsProduct {
		if err = repository.DeleteById(id); err != nil {
			return err
		}
	}

	return nil
}

package services

import (
	"errors"
	. "product-api/src/dto"
	. "product-api/src/mapper"
	. "product-api/src/models"
	. "product-api/src/repository"
)

var categoryRepository *FactoryMethodImpl[CategoryRepository] = InitFactory[CategoryRepository]()

func GetAllCategories() ([]CategoryResponseDTO, error) {

	repository, err := categoryRepository.GetRepository()

	if err != nil {
		return nil, err
	}

	categories, err := repository.ListAll()

	if err != nil {
		return nil, err
	}

	return ParseSliceEntityToDto[Category, CategoryResponseDTO](categories)
}

func GetCategoryById(id uint64) (CategoryDetailResponseDTO, error) {

	repository, err := categoryRepository.GetRepository()

	if err != nil {
		return CategoryDetailResponseDTO{}, err
	}

	category, err := repository.GetById(id)

	if err != nil {
		return CategoryDetailResponseDTO{}, err
	}

	if (category == Category{}) {
		return CategoryDetailResponseDTO{}, errors.New("Category Not Found")
	}

	return ParseEntityToDto[Category, CategoryDetailResponseDTO](category)
}

func EnableCategory(id uint64) error {

	repository, err := categoryRepository.GetRepository()

	if err != nil {
		return err
	}

	existsCategory, err := repository.ExistsById(id)

	if err != nil {
		return err
	}

	if existsCategory {
		if err := repository.EnableCategoryById(id); err != nil {
			return err
		}

		return nil
	}

	return errors.New("Enable fail, category not found.")
}

func DisableCategory(id uint64) error {

	repository, err := categoryRepository.GetRepository()

	if err != nil {
		return err
	}

	existsCategory, err := repository.ExistsById(id)

	if err != nil {
		return err
	}

	if existsCategory {
		if err := repository.DisableCategoryById(id); err != nil {
			return err
		}

		return nil
	}

	return errors.New("Disable fail, category not found.")
}

func SaveCategory(request CategoryRequestDTO) (CategoryDetailResponseDTO, error) {

	category, err := ParseDtoToEntity[CategoryRequestDTO, Category](request)

	var categoryRegistred CategoryDetailResponseDTO

	if err != nil {
		return categoryRegistred, err
	}

	repository, err := categoryRepository.GetRepository()

	if err != nil {
		return categoryRegistred, nil
	}

	category, err = repository.Save(category)

	if err != nil {
		return categoryRegistred, err
	}

	return ParseEntityToDto[Category, CategoryDetailResponseDTO](category)
}

func UpdateCategory(id uint64, categoryRequest CategoryRequestDTO) error {

	categoryUpdate, err := ParseDtoToEntity[CategoryRequestDTO, Category](categoryRequest)

	if err != nil {
		return err
	}

	repository, err := categoryRepository.GetRepository()

	if err != nil {
		return nil
	}

	existsCategory, err := repository.ExistsById(id)

	if err != nil {
		return err
	}

	if existsCategory {
		categoryUpdate.ID = id
		if err = repository.Update(categoryUpdate); err != nil {
			return err
		}
		return nil
	}

	return errors.New("Update fail, category Not Found")
}

func DeleteCategoryById(id uint64) error {

	repository, err := categoryRepository.GetRepository()

	if err != nil {
		return nil
	}

	existsCategory, err := repository.ExistsById(id)

	if existsCategory {
		if err := repository.DeleteById(id); err != nil {
			return err
		}
	}

	return nil
}

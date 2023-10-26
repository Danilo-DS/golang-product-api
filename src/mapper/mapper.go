package mapper

import (
	. "product-api/src/dto"
	. "product-api/src/repository"

	dtoMapper "github.com/devfeel/mapper"
)

type mapperType dtoType

type dtoType interface {
	ProductResponseDTO | CategoryResponseDTO | ProductRequestDTO | CategoryRequestDTO | CategoryDetailResponseDTO | ProductDetailResponseDTO
}

func ParseResponseStruct[F EntityType, T mapperType](model F) (T, error) {

	var genericStruct T

	if err := dtoMapper.Mapper(&model, &genericStruct); err != nil {
		return genericStruct, err
	}

	return genericStruct, nil
}

func ParseRequestStruct[F mapperType, T EntityType](model F) (T, error) {

	var genericStruct T

	if err := dtoMapper.Mapper(&model, &genericStruct); err != nil {
		return genericStruct, err
	}

	return genericStruct, nil
}

func ParseResponseSlice[F EntityType, T mapperType](models []F) ([]T, error) {

	genericSlice := []T{}

	for _, value := range models {
		genericStruct, err := ParseResponseStruct[F, T](value)

		if err != nil {
			return nil, err
		}

		genericSlice = append(genericSlice, genericStruct)
	}

	return genericSlice, nil
}

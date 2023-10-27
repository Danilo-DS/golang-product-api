package mapper

import (
	. "product-api/src/dto"
	. "product-api/src/repository"

	dtoMapper "github.com/devfeel/mapper"
)

// Custom type
type mapperType dtoType

// Associated types
type dtoType interface {
	ProductResponseDTO | CategoryResponseDTO | ProductRequestDTO | CategoryRequestDTO | CategoryDetailResponseDTO | ProductDetailResponseDTO
}

// ParseResponseStruct convert struct model to DTO type.
func ParseEntityToDto[F EntityType, T mapperType](model F) (T, error) {

	var genericStruct T

	if err := dtoMapper.Mapper(&model, &genericStruct); err != nil {
		return genericStruct, err
	}

	return genericStruct, nil
}

// ParseResponseStruct convert struct DTO to entity type.
func ParseDtoToEntity[F mapperType, T EntityType](model F) (T, error) {

	var genericStruct T

	if err := dtoMapper.Mapper(&model, &genericStruct); err != nil {
		return genericStruct, err
	}

	return genericStruct, nil
}

// ParseSliceEntityToDto convert slice of the type entity to dto type
func ParseSliceEntityToDto[F EntityType, T mapperType](models []F) ([]T, error) {

	genericSlice := []T{}

	for _, value := range models {
		genericStruct, err := ParseEntityToDto[F, T](value)

		if err != nil {
			return nil, err
		}

		genericSlice = append(genericSlice, genericStruct)
	}

	return genericSlice, nil
}

package dto

type CategoryRequestDTO struct {
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	Description string `json:"description"`
}

type CategoryProductRequestDTO struct {
	Id uint64 `json:"id" mapper:"id"`
}

type CategoryResponseDTO struct {
	Id     uint64 `json:"id" mapper:"id"`
	Name   string `json:"name" mapper:"name"`
	Active bool   `json:"isActive,omitempty" mapper:"active"`
}

type CategoryDetailResponseDTO struct {
	Id          uint64 `json:"id" mapper:"id"`
	Name        string `json:"name" mapper:"name"`
	Description string `json:"description,omitempty" mapper:"description"`
	Active      bool   `json:"isActive" mapper:"active"`
}

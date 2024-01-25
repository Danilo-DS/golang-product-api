package dto

type ProductRequestDTO struct {
	Name       string  `json:"name" mapper:"name"`
	Price      float64 `json:"price" mapper:"price"`
	CategoryId uint64  `json:"categoryId" mapper:"categoryId"`
	//CategoryProductRequestDTO `json:"category" mapper:"category"`
}

type ProductResponseDTO struct {
	Id                  uint64  `json:"id" mapper:"id"`
	Name                string  `json:"name" mapper:"name"`
	Price               float64 `json:"price" mapper:"price"`
	CategoryResponseDTO `json:"category" mapper:"category"`
}

type ProductDetailResponseDTO struct {
	Id                  uint64  `json:"id" mapper:"id"`
	Name                string  `json:"name" mapper:"name"`
	Price               float64 `json:"price" mapper:"price"`
	Barcode             string  `json:"barcode" mapper:"barcode"`
	CategoryResponseDTO `json:"category" mapper:"category"`
}

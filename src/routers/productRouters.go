package routers

import (
	"net/http"
	"product-api/src/controller"
)

var ProductRouters []RouterModel = []RouterModel{
	{
		URI:          "/product",
		HasParameter: true,
		Parameters:   []string{"name", "{name}"},
		Method:       http.MethodGet,
		HandleFunc:   controller.SearchProductByName,
	},
	{
		URI:          "/product",
		HasParameter: false,
		Method:       http.MethodGet,
		HandleFunc:   controller.GetAllProducts,
	},
	{
		URI:        "/product/{id}",
		Method:     http.MethodGet,
		HandleFunc: controller.GetProductById,
	},
	{
		URI:        "/product",
		Method:     http.MethodPost,
		HandleFunc: controller.SaveProduct,
	},
	{
		URI:        "/product/{id}",
		Method:     http.MethodPut,
		HandleFunc: controller.UpdateProduct,
	},
	{
		URI:        "/product/{id}",
		Method:     http.MethodDelete,
		HandleFunc: controller.DeleteProductById,
	},
}

package routers

import (
	"net/http"
	"product-api/src/controller"
)

var CategoryRouters []RouterModel = []RouterModel{
	{
		URI:        "/category",
		Method:     http.MethodGet,
		HandleFunc: controller.GetAllCategories,
	},
	{
		URI:        "/category/{id}",
		Method:     http.MethodGet,
		HandleFunc: controller.GetCategoryById,
	},
	{
		URI:        "/category/enable/{id}",
		Method:     http.MethodPatch,
		HandleFunc: controller.EnableCategory,
	},
	{
		URI:        "/category/disable/{id}",
		Method:     http.MethodPatch,
		HandleFunc: controller.DisableCategory,
	},
	{
		URI:        "/category",
		Method:     http.MethodPost,
		HandleFunc: controller.SaveCategory,
	},
	{
		URI:        "/category/{id}",
		Method:     http.MethodPut,
		HandleFunc: controller.UpdateCategory,
	},
	{
		URI:        "/category/{id}",
		Method:     http.MethodDelete,
		HandleFunc: controller.DeleteCategoryById,
	},
}

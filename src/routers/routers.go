package routers

import (
	"product-api/src/middlewares"

	muxRouter "github.com/gorilla/mux"
)

// Routers return all routers of the application
func Routers() *muxRouter.Router {

	routers := muxRouter.NewRouter()

	for _, router := range joinAllRouters(ProductRouters, CategoryRouters) {
		if router.HasParameter {
			routers.HandleFunc(router.URI, middlewares.Logger(router.HandleFunc)).Methods(router.Method).Queries(router.Parameters...)
		} else {
			routers.HandleFunc(router.URI, middlewares.Logger(router.HandleFunc)).Methods(router.Method)
		}
	}

	return routers
}

func joinAllRouters(anyRouters ...[]RouterModel) []RouterModel {

	var routers []RouterModel

	for _, router := range anyRouters {
		for _, valueRouter := range router {
			routers = append(routers, valueRouter)
		}
	}

	return routers
}

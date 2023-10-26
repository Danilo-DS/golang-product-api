package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	. "product-api/src/dto"
	"product-api/src/services"
	"product-api/src/utils"
)

func GetAllProducts(responseWriter http.ResponseWriter, request *http.Request) {
	products, err := services.GetAllProducts()

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseAsBody(responseWriter, http.StatusOK, products)
}

func GetProductById(responseWriter http.ResponseWriter, request *http.Request) {

	pathVariables := mux.Vars(request)

	idRequest, err := strconv.ParseUint(pathVariables["id"], 10, 64)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
	}

	product, err := services.GetProductById(idRequest)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseAsBody(responseWriter, http.StatusOK, product)
}

func SearchProductByName(responseWriter http.ResponseWriter, request *http.Request) {

	productName := request.URL.Query().Get("name")

	if len(strings.TrimSpace(productName)) == 0 {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, errors.New("Parameter \"name\" is empty"))
	}

	products, err := services.SearchProductByName(productName)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusNotFound, err)
		return
	}

	utils.ResponseAsBody(responseWriter, http.StatusOK, products)
}

func SaveProduct(responseWriter http.ResponseWriter, request *http.Request) {

	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusInternalServerError, err)
		return
	}

	var producRequest ProductRequestDTO

	if err = json.Unmarshal(requestBody, &producRequest); err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	productRegistred, err := services.SaveProduct(producRequest)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseAsBody(responseWriter, http.StatusCreated, productRegistred)
}

func UpdateProduct(responseWriter http.ResponseWriter, request *http.Request) {

	pathVariables := mux.Vars(request)

	idRequest, err := strconv.ParseUint(pathVariables["id"], 10, 64)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusInternalServerError, err)
		return
	}

	var productRequest ProductRequestDTO

	if err = json.Unmarshal(requestBody, &productRequest); err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	if err = services.UpdateProduct(idRequest, productRequest); err != nil {
		utils.ErrorResponse(responseWriter, http.StatusInternalServerError, err)
		return
	}

	utils.Response(responseWriter, http.StatusNoContent)
}

func DeleteProductById(responseWriter http.ResponseWriter, request *http.Request) {

	pathVariables := mux.Vars(request)

	idRequest, err := strconv.ParseUint(pathVariables["id"], 10, 64)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	if err = services.DeleteProductById(idRequest); err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	utils.Response(responseWriter, http.StatusNoContent)
}

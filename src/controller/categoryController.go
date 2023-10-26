package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	. "product-api/src/dto"
	"product-api/src/services"
	"product-api/src/utils"
)

func GetAllCategories(responseWriter http.ResponseWriter, request *http.Request) {
	category, err := services.GetAllCategories()

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseAsBody(responseWriter, http.StatusOK, category)
}

func GetCategoryById(responseWriter http.ResponseWriter, request *http.Request) {

	pathVariables := mux.Vars(request)

	idRequest, err := strconv.ParseUint(pathVariables["id"], 10, 64)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
	}

	category, err := services.GetCategoryById(idRequest)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseAsBody(responseWriter, http.StatusOK, category)
}

func EnableCategory(responseWriter http.ResponseWriter, request *http.Request) {

	pathVariables := mux.Vars(request)

	idRequest, err := strconv.ParseUint(pathVariables["id"], 10, 64)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	if err = services.EnableCategory(idRequest); err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	utils.Response(responseWriter, http.StatusNoContent)
}

func DisableCategory(responseWriter http.ResponseWriter, request *http.Request) {

	pathVariables := mux.Vars(request)

	idRequest, err := strconv.ParseUint(pathVariables["id"], 10, 64)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	if err = services.DisableCategory(idRequest); err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	utils.Response(responseWriter, http.StatusNoContent)
}

func SaveCategory(responseWriter http.ResponseWriter, request *http.Request) {

	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusInternalServerError, err)
		return
	}

	var categoryRequest CategoryRequestDTO

	if err = json.Unmarshal(requestBody, &categoryRequest); err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	categoryRegistred, err := services.SaveCategory(categoryRequest)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseAsBody(responseWriter, http.StatusCreated, categoryRegistred)
}

func UpdateCategory(responseWriter http.ResponseWriter, request *http.Request) {

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

	var categoryRequest CategoryRequestDTO

	if err = json.Unmarshal(requestBody, &categoryRequest); err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	if err = services.UpdateCategory(idRequest, categoryRequest); err != nil {
		utils.ErrorResponse(responseWriter, http.StatusInternalServerError, err)
		return
	}

	utils.Response(responseWriter, http.StatusNoContent)
}

func DeleteCategoryById(responseWriter http.ResponseWriter, request *http.Request) {

	pathVariables := mux.Vars(request)

	idRequest, err := strconv.ParseUint(pathVariables["id"], 10, 64)

	if err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	if err = services.DeleteCategoryById(idRequest); err != nil {
		utils.ErrorResponse(responseWriter, http.StatusBadRequest, err)
		return
	}

	utils.Response(responseWriter, http.StatusNoContent)
}

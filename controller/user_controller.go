package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sinarmas/models"
	"sinarmas/service"
)

type userController struct {
	service service.IUserService
}

func NewUserController(userService service.IUserService, router *mux.Router) *userController {

	controller := userController{service: userService}
	router.HandleFunc("/otp/validate", controller.ValidateOtp).Methods("POST")
	router.HandleFunc("/otp/generate", controller.GenerateOtp).Methods("POST")

	return &controller
}

func (c userController) ValidateOtp(w http.ResponseWriter, r *http.Request) {
	var request models.OtpValidationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	response, err := c.service.ValidateOtp(&request)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, ok := response.(*models.OtpValidationErrorResponse)
	if ok {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(response)
}

func (c userController) GenerateOtp(w http.ResponseWriter, r *http.Request) {
	var request models.OtpGenerationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	response, err := c.service.GenerateOtp(&request)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

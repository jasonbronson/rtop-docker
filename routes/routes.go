package routes

import (
	"rtop/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.GetInfoMachine).Methods("GET", "OPTIONS")

	return router

}

package routes

import (
	"crudapirest/usecases/buyer"
	"crudapirest/usecases/seller"
	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {
	// Seller
	subSellerRoute := router.PathPrefix("/seller/api").Subrouter()
	subSellerRoute.HandleFunc("/all", seller.GetAll).Methods("GET")
	subSellerRoute.HandleFunc("/find/{id}", seller.Get).Methods("GET")
	subSellerRoute.HandleFunc("/find/code/{code}", seller.GetByCode).Methods("GET")
	subSellerRoute.HandleFunc("/save", seller.Save).Methods("POST")
	subSellerRoute.HandleFunc("/delete/{id}", seller.Delete).Methods("POST")

	// Buyer
	subBuyerRoute := router.PathPrefix("/buyer/api").Subrouter()
	subBuyerRoute.HandleFunc("/all", buyer.GetAll).Methods("GET")
	subBuyerRoute.HandleFunc("/find/{id}", buyer.Get).Methods("GET")
	subSellerRoute.HandleFunc("/find/code/{code}", seller.GetByCode).Methods("GET")
	subBuyerRoute.HandleFunc("/save", buyer.Save).Methods("POST")
	subBuyerRoute.HandleFunc("/delete/{id}", buyer.Delete).Methods("POST")

}

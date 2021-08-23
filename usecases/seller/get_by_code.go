package seller

import (
	"crudapirest/commons"
	"crudapirest/models"
	json2 "encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetByCode(writer http.ResponseWriter, request *http.Request) {
	seller := models.Seller{}
	code := mux.Vars(request)["code"]
	db := commons.GetConnection()
	defer db.Close()

	db.Where("code = ?", code).Find(&seller)
	//db.Find(&seller, code)
	if seller.Code != 0 {
		json, _ := json2.Marshal(seller)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

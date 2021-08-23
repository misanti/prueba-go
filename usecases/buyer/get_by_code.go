package buyer

import (
	"crudapirest/commons"
	"crudapirest/models"
	json2 "encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetByCode(writer http.ResponseWriter, request *http.Request) {
	buyer := models.Buyer{}
	code := mux.Vars(request)["code"]
	db := commons.GetConnection()
	defer db.Close()

	db.Where("code = ?", code).Find(&buyer)
	//db.Find(&buyer, code)
	if buyer.Code != 0 {
		json, _ := json2.Marshal(buyer)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

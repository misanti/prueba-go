package seller

import (
	"crudapirest/commons"
	"crudapirest/models"
	json2 "encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Get(writer http.ResponseWriter, request *http.Request) {
	seller := models.Seller{}
	id := mux.Vars(request)["id"]
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&seller, id)
	if seller.ID > 0 {
		json, _ := json2.Marshal(seller)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

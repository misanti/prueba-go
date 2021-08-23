package buyer

import (
	"crudapirest/commons"
	"crudapirest/models"
	json2 "encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Get(writer http.ResponseWriter, request *http.Request) {
	buyer := models.Buyer{}
	id := mux.Vars(request)["id"]
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&buyer, id)
	if buyer.ID > 0 {
		json, _ := json2.Marshal(buyer)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

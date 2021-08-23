package buyer

import (
	"crudapirest/commons"
	"crudapirest/models"
	"github.com/gorilla/mux"
	"net/http"
)

func Delete(writer http.ResponseWriter, request *http.Request) {
	buyer := models.Buyer{}
	db := commons.GetConnection()
	defer db.Close()
	id := mux.Vars(request)["id"]

	db.Find(&buyer, id)
	if buyer.ID > 0 {
		db.Delete(buyer)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

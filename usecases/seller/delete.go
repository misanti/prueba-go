package seller

import (
	"crudapirest/commons"
	"crudapirest/models"
	"github.com/gorilla/mux"
	"net/http"
)

func Delete(writer http.ResponseWriter, request *http.Request) {
	seller := models.Seller{}
	db := commons.GetConnection()
	defer db.Close()
	id := mux.Vars(request)["id"]

	db.Find(&seller, id)
	if seller.ID > 0 {
		db.Delete(seller)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

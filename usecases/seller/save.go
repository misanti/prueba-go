package seller

import (
	"crudapirest/commons"
	"crudapirest/models"
	"encoding/json"
	"log"
	"net/http"
)

func Save(writer http.ResponseWriter, request *http.Request) {
	seller := models.Seller{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&seller)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&seller).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(seller)

	commons.SendReponse(writer, http.StatusCreated, json)
}

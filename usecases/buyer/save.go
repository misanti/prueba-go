package buyer

import (
	"crudapirest/commons"
	"crudapirest/models"
	"encoding/json"
	"log"
	"net/http"
)

func Save(writer http.ResponseWriter, request *http.Request) {
	buyer := models.Buyer{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&buyer)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&buyer).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(buyer)

	commons.SendReponse(writer, http.StatusCreated, json)
}

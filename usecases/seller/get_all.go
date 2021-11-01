package seller

import (
	"crudapirest/commons"
	"crudapirest/models"
	json2 "encoding/json"
	"net/http"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	sellers := []models.Seller{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&sellers)
	json, _ := json2.Marshal(sellers)
	commons.SendReponse(writer, http.StatusOK, json)
}


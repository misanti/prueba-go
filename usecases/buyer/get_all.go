package buyer

import (
	"crudapirest/commons"
	"crudapirest/models"
	json2 "encoding/json"
	"net/http"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	buyers := []models.Buyer{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&buyers)
	json, _ := json2.Marshal(buyers)
	commons.SendReponse(writer, http.StatusOK, json)
}

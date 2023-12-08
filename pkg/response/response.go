package response

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/bonaventurabs/be-tokopedia-auction/model"
)

type ErrorMsg struct {
	Error string `json:"error"`
}

func SetResponse(ctx context.Context, w http.ResponseWriter, httpStatus int, data interface{}) {

	// Setup header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PATCH,OPTIONS,POST")

	// Prepare Response
	response := model.Response{
		Code:   httpStatus,
		Status: http.StatusText(httpStatus),
		Data:   data,
	}

	// Encode Response
	encoded, err := json.Marshal(response)
	if err != nil {
		log.Printf("[ERROR][Util][WriteSuccess] marshaling response. %+v\n", err)
	}

	// Write Response
	w.WriteHeader(httpStatus)
	w.Write(encoded)
	return
}

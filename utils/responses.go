package utils

import (
	"encoding/json"
	"net"
	"net/http"

	// "kura_coin/models"
)

// ReturnData ...
type ReturnData struct {
	Status int
	Data []interface{}
	Error []interface{}
}

// RespondWithError ... This function sends error responses
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

// RespondWithOk ... This function sends error responses
func RespondWithOk(w http.ResponseWriter, msg string) {
	RespondWithJSON(w, http.StatusOK, map[string]string{"message": msg})
}


// RespondWithJSON ... This 
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return
}


func RespondTCP(data interface{}, conn net.Conn){

	responseByte, err := json.Marshal(data)
	if err != nil {
		println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}
	conn.Write(responseByte)
	
}

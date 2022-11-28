package health_controller

import (
	"encoding/json"
	"net/http"

	health_service "apis/src/services/health"

	"github.com/julienschmidt/httprouter"
)

func GetHealth(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var status = health_service.GetHealth()
	json.NewEncoder(writer).Encode(status)
}

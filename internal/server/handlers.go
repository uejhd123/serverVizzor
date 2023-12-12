package server

import (
	"encoding/json"
	"net/http"
	serverapi "serverVizzor/internal/serverApi"
)

func getStats(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(serverapi.GetSystemData())
}

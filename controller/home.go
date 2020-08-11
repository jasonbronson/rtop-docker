package controller

import (
	"encoding/json"
	"net/http"
	"rtop/rtop"
)

func GetInfoMachine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stats := rtop.Stats{}
	rtop.GetAllStats(&stats)
	displayStats := &rtop.StatsDisplay{}
	rtop.ParseStatsDisplay(stats, displayStats)
	Send(w, http.StatusOK, displayStats)
}

//Send respond with JSON
func Send(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)

}

package controller

import (
	"net/http"
	"rtop/rtop"
)

func GetInfoMachine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := rtop.SSHConnect(config.Cfg.HostUser, config.Cfg.HostName, config.Cfg.HostPassword, "")
	stats := rtop.Stats{}
	rtop.GetAllStats(client, &stats)
	Send(w, http.StatusOK, stats)
}

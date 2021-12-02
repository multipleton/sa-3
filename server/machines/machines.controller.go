package machines

import (
	"encoding/json"
	"net/http"
)

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(ms *MachinesService) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListMachines(rw, ms)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListMachines(rw http.ResponseWriter, ms MachinesService) {
	rw.Header().Set("content-type", "application/json")
	res, err := ms.ReadAll()
	if err != nil {
		rw.WriteHeader(404)
		_ := json.NewEncoder(rw).Encode("cannot read all machines")
		return
	}
	rw.WriteHeader(200)
	err := json.NewEncoder(rw).Encode("123")
	if err != nil {
	}
}

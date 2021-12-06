package machines

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/multipleton/sa-3/utils"
)

type MachinesController struct {
	ms *MachinesService
}

func (mc *MachinesController) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/machines", mc.ServeAllMachines).Methods("GET")
}

func (mc *MachinesController) ServeAllMachines(rw http.ResponseWriter, r *http.Request) {
	res, err := mc.ms.ReadAll()
	if err != nil {
		utils.Respond(rw, 500, err)
		return
	}
	if err := json.NewEncoder(rw).Encode(res); err != nil {
		utils.Respond(rw, 500, err)
		return
	}
}

func NewMachinesController(ms *MachinesService) *MachinesController {
	return &MachinesController{ms: ms}
}

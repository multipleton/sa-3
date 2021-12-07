package machines

import (
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
	result, err := mc.ms.ReadAll()
	if err != nil {
		utils.Respond(rw, http.StatusInternalServerError, err)
		return
	}
	utils.Respond(rw, http.StatusOK, result)
}

func NewMachinesController(ms *MachinesService) *MachinesController {
	return &MachinesController{ms: ms}
}

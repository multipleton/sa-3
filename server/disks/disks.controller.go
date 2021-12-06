package disks

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/multipleton/sa-3/utils"
)

type DisksController struct {
	disksService *DisksService
}

func (dc *DisksController) HandleRoutes(router *mux.Router) {
	router.HandleFunc("/disks/{diskId:[0-9]+}", dc.ConnectToMachine).Methods("PATCH")
}

func (dc *DisksController) ConnectToMachine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	diskId, err := strconv.Atoi(vars["diskId"])
	if err != nil {
		message := "invalid request param: diskId"
		utils.Respond(w, http.StatusBadRequest, errors.New(message))
	}
	var dto DiskConnectToMachineDto
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		message := "invalid request body"
		utils.Respond(w, http.StatusBadRequest, errors.New(message))
	}
	result, err := dc.disksService.ConnectToMachine(uint32(diskId), &dto)
	if err != nil {
		/*
		   TODO: implement cheching for error type to return 404
		   cannot be done until custom EntityNotFoundError won't be implemented
		*/
		utils.Respond(w, http.StatusInternalServerError, err)
	}
	utils.Respond(w, http.StatusOK, result)
}

func NewDisksController(disksService *DisksService) *DisksController {
	return &DisksController{disksService: disksService}
}

package machines

import "github.com/google/wire"

var Module = wire.NewSet(NewMachinesRepository, NewMachinesService, NewMachinesController)

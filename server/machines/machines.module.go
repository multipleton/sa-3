package machines

import "github.com/google/wire"

var Provider = wire.NewSet(NewMachinesRepository, NewMachinesService, NewMachinesController)

package disks

import "github.com/google/wire"

var Module = wire.NewSet(NewDisksRepository, NewDisksService, NewDisksController)

package disks

import "github.com/google/wire"

var Provider = wire.NewSet(NewDisksRepository, NewDisksService, NewDisksController)

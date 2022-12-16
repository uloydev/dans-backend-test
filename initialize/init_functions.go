package initialize

import (
	"dans-backend-test/app/controller"
)

var InitFunctions = []InitFunc{
	controller.InitializeUserController,
	controller.InitializeJobController,
}

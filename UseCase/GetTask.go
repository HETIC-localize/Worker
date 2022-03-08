package UseCase

import (
	"HETIC-localize/Worker/Model"
	"HETIC-localize/Worker/Service"
)

func GetTask() Model.Task {
	return Service.GCPPubSubGetTask()
}
package UseCase

import (
	"HETIC-localize/Worker/Model"
	"HETIC-localize/Worker/Service"
)

func SendTranslationToBackend(translation Model.Translation) {
	Service.StoragePersistTranslation(translation)
}

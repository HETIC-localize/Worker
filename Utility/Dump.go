package Utility

import (
	"HETIC-localize/Worker/Model"
	"fmt"
)

func DumpTranslation(translation Model.Translation)  {

	fmt.Println("ID :", translation.ID)
	for lang, text := range translation.Items {
		fmt.Println(lang, ":" , text)
	}
}
package Utility

import (
	"HETIC-localize/Worker/Model"
	"fmt"
)

func DumpTranslation(translation Model.Translation)  {

	fmt.Println("Code :", translation.Code)
	for lang, text := range translation.Items {
		fmt.Println(lang, ":" , text)
	}
}
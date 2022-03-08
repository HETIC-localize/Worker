package main

import (
	"HETIC-localize/Worker/Config"
	"HETIC-localize/Worker/Model"
	"HETIC-localize/Worker/UseCase"
	"HETIC-localize/Worker/Utility"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	godotenv.Load()

	fmt.Println("Welcome to" , os.Getenv("APP_NAME") , "!")

	for {

		fmt.Println("* Waiting task")

		task := UseCase.GetTask()

		fmt.Println("* Handle task")

		translation := Model.Translation{Items: make(map[string]string)}

		translation.ID = task.ID
		translation.Items[task.Src] = task.Val

		for _, lang := range Config.Lang() {

			if lang == task.Src {
				continue
			}

			translation.Items[lang] = UseCase.TranslateByTextAndLang(task.Val, lang)
		}

		Utility.DumpTranslation(translation)

		UseCase.StoreTranslation(translation)

		fmt.Println("* Terminate task")
	}
}
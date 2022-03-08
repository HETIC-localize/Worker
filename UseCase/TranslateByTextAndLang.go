package UseCase

import "HETIC-localize/Worker/Service"

func TranslateByTextAndLang(text string, lang string) string {
	return Service.DLTranslate(text, lang)
}

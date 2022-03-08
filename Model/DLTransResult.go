package Model

type DLTransResult struct {

	Translations []struct {
		DetectedSourceLanguage string
		Text                   string
	}
}

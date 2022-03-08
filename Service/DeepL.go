package Service

// @see https://www.deepl.com/docs-api

import (
	"HETIC-localize/Worker/Model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func DLTranslate(text string, lang string) string {

	params := url.Values{}
	params.Add("auth_key", os.Getenv("DEEPL_AUTH_KEY"))
	params.Add("text", text)
	params.Add("target_lang", lang)

	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://api-free.deepl.com/v2/translate", body)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result Model.DLTransResult

	if err := json.Unmarshal([]byte(string(data)), &result); err != nil {
		log.Fatal(err)
	}

	return result.Translations[0].Text
}

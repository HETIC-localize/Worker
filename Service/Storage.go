package Service

import (
	"HETIC-localize/Worker/Model"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)


func storageConnection() *sql.DB {

	db, err := sql.Open(
		"mysql",
		os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME"),
	)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func StoragePersistTranslation(translation Model.Translation) {

	db := storageConnection()

	for lang, text := range translation.Items {

		_, err := db.Query("UPDATE `translation` SET `value` = '" + text + "' WHERE country = '" + lang + "' AND code = '" + translation.Code + "'")

		if err != nil {
			panic(err.Error())
		}
	}

	defer db.Close()
}
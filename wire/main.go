package wire

import (
	"fmt"

	"github.com/arham09/go-di-example/service/configs/database"
)

func Main() {
	dbHost := "localhost"
	dbUser := "medea"
	dbPassword := "developer"
	dbName := "battlefield"

	dsn := fmt.Sprintf(`postgres://%s:%s@%s/%s?sslmode=disable`, dbUser, dbPassword, dbHost, dbName)

	db, err := database.NewDB(dsn)

	if err != nil {
		panic(err)
	}

	s := StartingService(db)

	s.Run()
}

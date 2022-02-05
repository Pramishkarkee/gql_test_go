package db

import (
    "log"

    "github.com/parmishKarkee/test/graph/pkg/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init() *gorm.DB {
    databasename := "postgres"
	// database := "postgres"
	databasepassword := "root1234"
    dbURL := "postgres://postgres:" + databasepassword + "@localhost/" + databasename + "?sslmode=disable"


    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Book{})

    return db
}
// func GetDatabase() *gorm.DB {
// 	databasename := "postgres"
// 	database := "postgres"
// 	databasepassword := "root1234"
// 	databaseurl := "postgres://postgres:" + databasepassword + "@localhost/" + databasename + "?sslmode=disable"

// 	connection, err := gorm.Open(database, databaseurl)
// 	if err != nil {
// 		log.Fatalln("Invalid database url")
// 	}
// 	sqldb := connection.DB()

// 	err = sqldb.Ping()
// 	if err != nil {
// 		log.Fatal("Database connected")
// 	}
// 	fmt.Println("Database connection successful.")
// 	return connection
// }

// func CloseDatabase(connection *gorm.DB) {
// 	sqldb := connection.DB()
// 	sqldb.Close()
// }
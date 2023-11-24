package config

//============= cara native
// import (
// 	"database/sql"
// 	"fmt"

// 	_"github.com/lib/pq"
// )

// const(
// 	host		="localhost"
// 	port		= 5432
// 	user		="postgres"
// 	password	="1234"
// 	dbname		="db-go-sql"
// )
// var(
// 	db *sql.DB
// 	err error
// )

// func Connect()(*sql.DB, error){
// 	psqlInfo := fmt. Sprintf ("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	db, err = sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic (err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		panic (err)
// 	}

// 	fmt.Println ("Successfully connected to database")

// 	return db, nil
// }
//===================================================

//====== GORM

import (
	"fmt"
	"harikedua/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"github.com/joho/godotenv"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "1234"
// 	dbname   = "db-go-sql"
// )


var (
	db  *gorm.DB
	err error
)

func Connect() {

	err := godotenv.Load()
	if err != nil{
		panic(err)
	}

	var(
		host 		= os.Getenv("PGHOST")
		port 		= os.Getenv("PGPORT")
		user 		= os.Getenv("PGUSER")
		password	= os.Getenv("PGPASSWORD")
		dbname		= os.Getenv("PGDATABASE")	
	)
	psqlInfo := fmt.Sprintf(`
	host=%s 
	port=%s 
	user=%s `+`
	password=%s 
	dbname=%s 
	sslmode=disable`,
		host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	// db.AutoMigrate(&model.Employee{}) biar debug nya gak muncul
	db.Debug().AutoMigrate(&model.Employee{})
	db.Debug().AutoMigrate(&model.Item{})
}

func GetDB() *gorm.DB {
	return db
}


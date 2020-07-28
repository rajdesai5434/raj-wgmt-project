package models

import (
	"database/sql"

	_ "github.com/lib/pq" // here
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

//MyDB global variable to have a cached connection to DB
var MyDB *sql.DB

//ConnectToDB To connect to PG
func ConnectToDB() {
	//psqlInfo := fmt.Sprintf("postgres://postgres:%s@%s:%d/%s?sslmode=disable", user, host, port, dbname)
	//psqlInfo := fmt.Sprintf("port=%d host=%s user=%s "+
	//"password=%s dbname=%s sslmode=disable",
	//port, host, user, password, dbname)

	var err error
	MyDB, err = sql.Open("postgres", "postgres://rajdesai:@localhost:5433/wing_mate?sslmode=disable")
	if err != nil {
		panic(err)
	}

}

//DBClose will close connection
func DBClose() {
	_ = MyDB.Close()
}

/*func Connect() {
	opts := &pg.Options{
		User:     user,
		Password: "",
		Addr:     "localhost:5433",
		Database: dbname,
	}

	MyDB = pg.Connect(opts)
	if MyDB == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
}

func DBClose() {
	_ = MyDB.Close()
}*/

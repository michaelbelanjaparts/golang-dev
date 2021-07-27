package utils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type server struct {
	url      string
	username string
	password string
	hostname string
	dbname   string
	driver   string
}

type Server interface {
	InitDB() *sql.DB
}

func Setup() Server {
	godotenv.Load()
	app := server{
		url:      os.Getenv("DB_URI"),
		username: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		hostname: os.Getenv("DB_HOST"),
		dbname:   os.Getenv("DB_NAME"),
		driver:   "mysql",
	}
	return &app
}

func (s *server) InitDB() *sql.DB {
	db, err := sql.Open(s.driver, fmt.Sprintf("%v:%v@tcp(%v:3306)/%v", s.username, s.password, s.hostname, s.dbname))
	if err != nil {
		panic(err)
	}
	return db
}

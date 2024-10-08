package config

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"sync"
	"xm-companies/repository"

	"github.com/go-sql-driver/mysql"
)

type Application struct {
	Models *repository.Models
	JwtKey []byte
	SqlCfg *mysql.Config
}

var instance *Application
var once sync.Once
var db *sql.DB

func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

func GetInstance() *Application {

	once.Do(func() {
		instance = &Application{
			Models: repository.New(db),
			JwtKey: []byte("default_secret_key"),
		}
	})
	return instance
}

func (a *Application) LoadConfig(configFile string) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	var dat map[string]interface{}
	json.Unmarshal(data, &dat)

	jwtKey, ok := dat["jwt_key"].(string)
	if !ok {
		log.Fatalf("jwt_key is not a string")
	}
	a.JwtKey = []byte(jwtKey)

	// db_user, ok := dat["db_user"].(string)
	// if !ok {
	// 	log.Fatalf("db_user is not a string")
	// }

	// db_password, ok := dat["db_password"].(string)
	// if !ok {
	// 	log.Fatalf("db_password is not a string")
	// }

	// db_host, ok := dat["db_host"].(string)
	// if !ok {
	// 	log.Fatalf("db_host is not a string")
	// }

	// db_name, ok := dat["db_name"].(string)
	// if !ok {
	// 	log.Fatalf("db_name is not a string")
	// }

	// a.SqlCfg = &mysql.Config{
	// 	User:   db_user,
	// 	Passwd: db_password,
	// 	Net:    "tcp",
	// 	Addr:   db_host,
	// 	DBName: db_name,
	// }
	// Bad practise :)
	// log.Println("loaded jwt key:", a.JwtKey)
}

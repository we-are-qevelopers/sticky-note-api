package infra

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Endpoint string
	Port     int
	User     string
	Password string
	DBName   string
}

// todo: 正しい値を入力
var dbConfigs = map[string]DBConfig{
	"dev": {
		Endpoint: "postgres",
		Port:     5432,
		User:     "api_user",
		Password: "password",
		DBName:   "app_db",
	},
}

func (d DBConfig) GetDSN() string {
	dsn := fmt.Sprintf("host=%s user=%s port=%d dbname=%s password=%s sslmode=disable",
		d.Endpoint,
		d.User,
		d.Port,
		d.DBName,
		d.Password,
	)

	return dsn
}

func SetupDB() error {
	var err error
	dsn := dbConfigs["dev"].GetDSN()
	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DriverName: "postgres",
			DSN:        dsn,
		}))

	if err != nil {
		log.Fatalln("接続失敗", err)
	}

	return nil
}

// memo
// DSN=データソース名

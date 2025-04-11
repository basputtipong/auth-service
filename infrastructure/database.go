package infrastructure

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host       string              `mapstructure:"host"`
	Port       string              `mapstructure:"port"`
	Username   string              `mapstructure:"username"`
	Password   string              `mapstructure:"password"`
	Name       string              `mapstructure:"name"`
	Connection SqlConnectionConfig `mapstructure:"connection"`
}

type SqlConnectionConfig struct {
	MaxOpen     *int `mapstructure:"maxopen"`
	MaxIdle     *int `mapstructure:"maxidle"`
	MaxLifetime *int `mapstructure:"maxlifetime"`
}

var (
	DB *gorm.DB
)

func InitDB() {
	var dbConfig DBConfig
	if err := viper.UnmarshalKey("db", &dbConfig); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDb, _ := db.DB()
	if dbConfig.Connection.MaxOpen != nil {
		sqlDb.SetMaxOpenConns(*dbConfig.Connection.MaxOpen)
	} else {
		sqlDb.SetMaxOpenConns(10)
	}

	if dbConfig.Connection.MaxIdle != nil {
		sqlDb.SetMaxIdleConns(*dbConfig.Connection.MaxIdle)
	} else {
		sqlDb.SetMaxIdleConns(2)
	}

	if dbConfig.Connection.MaxLifetime != nil {
		sqlDb.SetConnMaxLifetime(time.Minute * time.Duration(*dbConfig.Connection.MaxLifetime))
	} else {
		sqlDb.SetConnMaxLifetime(time.Minute * 5)
	}

	DB = db
}

func pingDb(dsn *gorm.DB) error {
	db, err := dsn.DB()
	if err != nil {
		return err
	} else {
		return db.Ping()
	}
}

func PingAllDb() error {
	dbs := []*gorm.DB{
		DB,
	}

	for _, db := range dbs {
		if err := pingDb(db); err != nil {
			return err
		}
	}
	return nil
}

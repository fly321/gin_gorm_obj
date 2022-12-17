package global

import (
	"fmt"
	"github.com/spf13/viper"
)

// Database yaml
/**
Database:
  DBType: mysql
  UserName: root
  Password: password
  Host: 127.0.0.1:3306
  DBName: dig
  Charset: utf8
  ParseTime: True
  MaxIdleConns: 10
  MaxOpenConns: 30
*/
type Database struct {
	DbType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
	Loc          string
}

// Server yaml
/**
Server:
  RunMode: debug
  HttpPort: 8000
  ReadTimeout: 60
  WriteTimeout: 60
*/
type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  int
	WriteTimeout int
}

// viper

var ServerSetting *Server
var DatabaseSetting *Database

func init() {
	println("init config.yaml")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	fmt.Println(viper.AllSettings())
	if err != nil {
		panic(err)
	}
	err = viper.UnmarshalKey("Server", &ServerSetting)
	//fmt.Printf("ServerSetting: %#v", ServerSetting)
	if err != nil {
		panic(err)
	}
	err = viper.UnmarshalKey("Database", &DatabaseSetting)
	//fmt.Printf("DatabaseSetting: %#v", DatabaseSetting)
	if err != nil {
		panic(err)
	}
}

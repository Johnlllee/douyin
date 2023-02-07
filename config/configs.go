package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"strings"
)

// USER INFO
const MAX_USERNAME_LENGTH = 100
const MAX_PASSWORD_LENGTH = 20
const MIN_PASSWORD_LENGTH = 8

type Mysql struct {
	Host      string
	Port      int
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime bool `toml:"parse_time"`
	Loc       string
}

type Server struct {
	IP   string
	Port int
}

type Redis struct {
	IP       string
	Port     int
	Database int
}

type Config struct {
	DB               Mysql `toml:"mysql"`
	Server           `toml:"server"`
	StaticSourcePath string `toml:"static_source_path"`
	RDB              Redis  `toml:"redis"`
}

var Info Config

func init() {
	if _, err := toml.DecodeFile("/Users/lianghaoran/golangWorkSpace/src/douyin/config/config.toml", &Info); err != nil {
		panic(err)
	}
	//去除左右的空格
	//strings.Trim(Info.Server.IP, " ")
	//strings.Trim(Info.RDB.IP, " ")
	strings.Trim(Info.DB.Host, " ")
}

func DBConnectString() string {
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		Info.DB.Username, Info.DB.Password, Info.DB.Host, Info.DB.Port, Info.DB.Database,
		Info.DB.Charset, Info.DB.ParseTime, Info.DB.Loc)
	log.Println(arg)
	return arg
}

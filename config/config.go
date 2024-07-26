package config

import (
	"github.com/spf13/viper"
)

func Configuration() string {
	viper.SetConfigFile("config/config.env")
	viper.ReadInConfig()
	user := viper.Get("user").(string)
	host := viper.Get("host").(string)
	password := viper.Get("password").(string)
	dbname := viper.Get("dbname").(string)
	port := viper.Get("port").(string)
	var dsn string
	dsn += "host=" + host
	dsn += " user=" + user
	dsn += " password=" + password
	dsn += " dbname=" + dbname
	dsn += " port=" + port
	return dsn
}

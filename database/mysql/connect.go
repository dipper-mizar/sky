package mysql

import (
	"fmt"
	"github.com/dipper-mizar/sky/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type InitParams struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type ParamSetter func(*InitParams)

func (params *InitParams) SetHost(host string) ParamSetter {
	return func(params *InitParams) {
		params.Host = host
	}
}

func (params *InitParams) SetPort(port string) ParamSetter {
	return func(params *InitParams) {
		params.Port = port
	}
}

func (params *InitParams) SetUser(user string) ParamSetter {
	return func(params *InitParams) {
		params.User = user
	}
}

func (params *InitParams) SetPassword(password string) ParamSetter {
	return func(params *InitParams) {
		params.Password = password
	}
}

func (params *InitParams) SetDatabase(database string) ParamSetter {
	return func(params *InitParams) {
		params.Database = database
	}
}

func (params *InitParams) Init(functions ...ParamSetter) *gorm.DB {
	p := &InitParams{
		Host:     conf.MySQLHost,
		Port:     conf.MySQLPort,
		User:     conf.MySQLUser,
		Password: conf.MySQLPassword,
		Database: conf.MySQLDatabase,
	}
	for _, function := range functions {
		function(p)
	}
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", params.User,
		params.Password, params.Host, params.Port, params.Database)
	database, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		// TODO: Write error into logger.
	}
	return database
}

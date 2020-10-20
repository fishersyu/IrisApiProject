package config

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"path/filepath"

	"github.com/jinzhu/configor"
	"github.com/snowlyg/IrisAdminApi/libs"
)

var Config = struct {
	Bindata  bool   `default:"false" env:"Bindata"`
	Debug    bool   `default:"false" env:"Debug"`
	LogLevel string `default:"info" env:"Loglevel"`
	HTTPS    bool   `default:"false" env:"HTTPS"`
	Certpath string `default:"" env:"Certpath"`
	Certkey  string `default:"" env:"Certkey"`
	Port     int    `default:"8085" env:"PORT"`
	Host     string `default:"127.0.0.1" env:"Host"`
	Admin    struct {
		UserName        string `env:"AdminUserName" default:"username"`
		Name            string `env:"AdminName" default:"name"`
		Pwd             string `env:"AdminPwd" default:"123456"`
		RoleName        string `env:"AdminRoleName" default:"admin"`
		RoleDisplayName string `env:"TenantRoleDisplayName" default:"超级管理员"`
	}
	DB struct {
		Prefix   string `env:"DBPrefix" default:"iris_"`
		Name     string `env:"DBName" default:"goirisadminapi"`
		Adapter  string `env:"DBAdapter" default:"sqlite3"`
		Host     string `env:"DBHost" default:"mysql"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser" default:"root"`
		Password string `env:"DBPassword" default:"123456"`
	}
}{}

func init() {
	configPath := filepath.Join(libs.CWD(), "application.yml")

	logger.Println(fmt.Sprintf("配置路径：%s", configPath))
	if err := configor.Load(&Config, configPath); err != nil {
		logger.Println(fmt.Sprintf("Config Path:%s ,Error:%s", configPath, err.Error()))
	} else {
		logger.Println(fmt.Sprintf("系统配置：%+v", Config))
	}
}

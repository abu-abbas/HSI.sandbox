package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port  string `json:"port"`
	Debug string `json:"debug"`
}

type DbConfig struct {
	Driver string `json:"driver"`
	DbFile string `json:"dbfile"`
}

type yamlConfig struct {
	ServerConfig ServerConfig `yaml:"serverDetails"`
	DbConfig     DbConfig     `yaml:"dbConfig"`
}

var Viper *viper.Viper

func init() {
	readConfig("config/config")
}

func readConfig(filename string) {
	Viper = viper.New()
	Viper.AddConfigPath(".")
	Viper.SetConfigName(filename)

	err := Viper.ReadInConfig()
	if err != nil {
		log.Error("Error saat membaca file config", err)
	}

	replacer := strings.NewReplacer(".", "_")
	Viper.SetEnvKeyReplacer(replacer)
	Viper.AutomaticEnv()
}

func GetYamlValue() *yamlConfig {
	server := &ServerConfig{
		Port:  Viper.GetString("service.port"),
		Debug: Viper.GetString("service.debug"),
	}

	db := &DbConfig{
		Driver: Viper.GetString("database.driver"),
		DbFile: Viper.GetString("database.dbfile"),
	}

	yaml := &yamlConfig{*server, *db}
	return yaml
}

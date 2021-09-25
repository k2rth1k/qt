package config

import (
	"github.com/spf13/viper"
	"os"
)

// SQLConfig stores connection data for the database
type SQLConfig struct {
	Host    string
	Port    string
	DBName  string
	User    string
	Pass    string
	SSLMode string
}

var conf ServiceConfig

type ServiceConfig struct {
	DBConfig SQLConfig
}

func setViperDefaults() {
	// viper is case-insensitive

	// service configuration
	viper.SetDefault("grpc-port", 50051)
	viper.SetDefault("rest-port", 50443)
	// empty to listen on all interfaces
	// postgress configuration
	viper.SetDefault("postgres_host", "localhost")
	viper.SetDefault("postgres_port", "5432")
	viper.SetDefault("postgres_db", "quick_trade")
	viper.SetDefault("postgres_user", "quick_trade")
	viper.SetDefault("postgres_password", "quick_trade")
	viper.SetDefault("postgres_ssl_mode", "disable")
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")
	// kubernetes namespace
}
func SetupConfig() {
	// viper is case-insensitive
	// viper will check in the following order:
	// override, flag, env, config file, key/value store, default
	setViperDefaults()
	cfg := &conf
	dbConfig := SQLConfig{
		Host:    viper.GetString("postgres_host"),
		Port:    viper.GetString("postgres_port"),
		DBName:  viper.GetString("postgres_db"),
		User:    viper.GetString("postgres_user"),
		Pass:    viper.GetString("postgres_password"),
		SSLMode: viper.GetString("postgres_ssl_mode"),
	}
	cfg.DBConfig = dbConfig
}

func GetServiceConfig() *ServiceConfig {
	// todo: this re-evaluates configuration whenever it is called.
	// if needed, move this into the main entry point of the application/package
	SetupConfig()
	return &conf
}

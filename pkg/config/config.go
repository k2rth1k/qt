package config

import (
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

func getConfigs(key, defaultValue string) string {
	dsn := os.Getenv(key)
	if len(dsn) == 0 {
		dsn = defaultValue
	}
	return dsn
}

func SetupConfig() {
	// viper is case-insensitive
	// viper will check in the following order:
	// override, flag, env, config file, key/value store, default
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")
	cfg := &conf
	dbConfig := SQLConfig{
		Host:    getConfigs("postgres_host", "localhost"),
		Port:    getConfigs("postgres_port", "5432"),
		DBName:  getConfigs("postgres_db", "qt"),
		User:    getConfigs("postgres_user", "qt"),
		Pass:    getConfigs("postgres_password", "qt"),
		SSLMode: getConfigs("postgres_ssl_mode", "disable"),
	}
	cfg.DBConfig = dbConfig
}

func GetServiceConfig() *ServiceConfig {
	// todo: this re-evaluates configuration whenever it is called.
	// if needed, move this into the main entry point of the application/package
	SetupConfig()
	return &conf
}

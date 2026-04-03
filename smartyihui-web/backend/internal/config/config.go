package config

import "os"

type Config struct {
	AppPort        string
	MySQLHost      string
	MySQLPort      string
	MySQLUser      string
	MySQLPassword  string
	MySQLDatabase  string
	MySQLParams    string
	FrontendOrigin string
}

func Load() Config {
	return Config{
		AppPort:        getenv("APP_PORT", "31098"),
		MySQLHost:      getenv("MYSQL_HOST", "127.0.0.1"),
		MySQLPort:      getenv("MYSQL_PORT", "3306"),
		MySQLUser:      getenv("MYSQL_USER", "root"),
		MySQLPassword:  os.Getenv("MYSQL_PASSWORD"),
		MySQLDatabase:  getenv("MYSQL_DATABASE", "smartyihui"),
		MySQLParams:    getenv("MYSQL_PARAMS", "charset=utf8mb4&parseTime=True&loc=Local"),
		FrontendOrigin: getenv("FRONTEND_ORIGIN", "http://127.0.0.1:31098"),
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

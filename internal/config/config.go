package config

import "os"

type Config struct {
	DBMode       string
	DBHost       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBPort       string
	RunMode      string
	GeoNamesUser string
	HTTPPort     string
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func NewConfig() *Config {
	config := &Config{
		DBMode:       getEnv("DB_MODE", "postgres"),
		DBHost:       getEnv("PGHOST", "localhost"),
		DBUser:       getEnv("VECIN_DB_USER", ""),
		DBPassword:   getEnv("VECIN_DB_PASSWORD", ""),
		DBName:       getEnv("VECIN_DB", ""),
		DBPort:       getEnv("PGPORT", "5432"),
		RunMode:      getEnv("RUN_MODE", "dev"),
		GeoNamesUser: getEnv("GEONAMES_USERNAME", ""),
		HTTPPort:     getEnv("PORT", "8180"),
	}

	return config
}

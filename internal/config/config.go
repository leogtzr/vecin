package config

import (
	"fmt"
	"os"
	"time"
)

type Mailing struct {
	ApiKey           string
	ConfirmationLink string
	EmailSubject     string
	MailSenderEmail  string
}

type Config struct {
	DBMode              string
	DBHost              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBPort              string
	RunMode             string
	GeoNamesUser        string
	HTTPPort            string
	UserTokenLen        int
	UserTokenExpiryDays time.Duration
	MailSenderApiKey    string

	Mailing Mailing
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func NewConfig() (*Config, error) {
	config := &Config{
		DBMode:              getEnv("DB_MODE", "postgres"),
		DBHost:              getEnv("PGHOST", "localhost"),
		DBUser:              getEnv("VECIN_DB_USER", ""),
		DBPassword:          getEnv("VECIN_DB_PASSWORD", ""),
		DBName:              getEnv("VECIN_DB", ""),
		DBPort:              getEnv("VECIN_DB_PORT", "5432"),
		RunMode:             getEnv("RUN_MODE", "dev"),
		GeoNamesUser:        getEnv("GEONAMES_USERNAME", ""),
		HTTPPort:            getEnv("PORT", "8180"),
		UserTokenLen:        16, // 32 characters length (hex)
		UserTokenExpiryDays: 30 * 24 * time.Hour,
		MailSenderApiKey:    getEnv("MAILSENDER_API_KEY", ""),
		Mailing: Mailing{
			ApiKey:           getEnv("MAILSENDER_API_KEY", ""),
			ConfirmationLink: getEnv("EMAIL_ACCOUNT_CONFIRMATION_LINK", ""),
			EmailSubject:     getEnv("VECIN_CONFIRMATION_EMAIL_SUBJECT", ""),
			MailSenderEmail:  getEnv("MAILSENDER_EMAIL_SENDER", ""),
		},
	}

	if config.MailSenderApiKey == "" {
		return nil, fmt.Errorf("mail sender api key not set")
	}

	if config.Mailing.ApiKey == "" {
		return nil, fmt.Errorf("mail sender api key not set")
	}

	return config, nil
}

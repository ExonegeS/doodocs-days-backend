package config

import (
	"fmt"
	"log/slog"
	"os"
)

var (
	ENV_AUTH_USER string
	ENV_AUTH_PASS string
	ENV_MAIL_USER string
	ENV_MAIL_PASS string
)

func UpdateENV() error {
	ENV_AUTH_USER = os.Getenv("DOODOCS_DAYS2_BACKEND_AUTH_USERNAME")
	ENV_AUTH_PASS = os.Getenv("DOODOCS_DAYS2_BACKEND_AUTH_PASSWORD")
	ENV_MAIL_USER = os.Getenv("DOODOCS_DAYS2_BACKEND_MAIL_USERNAME")
	ENV_MAIL_PASS = os.Getenv("DOODOCS_DAYS2_BACKEND_MAIL_PASSWORD")

	if ENV_AUTH_USER == "" {
		slog.Error("Error: DOODOCS_DAYS2_BACKEND_AUTH_USERNAME is not set")
		return fmt.Errorf("DOODOCS_DAYS2_BACKEND_AUTH_USERNAME is not set")

	}
	if ENV_AUTH_PASS == "" {
		slog.Error("Error: DOODOCS_DAYS2_BACKEND_AUTH_PASSWORD is not set")
		return fmt.Errorf("DOODOCS_DAYS2_BACKEND_AUTH_PASSWORD is not set")
	}
	if ENV_MAIL_USER == "" {
		slog.Error("Error: DOODOCS_DAYS2_BACKEND_MAIL_USERNAME is not set")
		return fmt.Errorf("DOODOCS_DAYS2_BACKEND_MAIL_USERNAME is not set")
	}
	if ENV_MAIL_PASS == "" {
		slog.Error("Error: DOODOCS_DAYS2_BACKEND_MAIL_PASSWORD is not set")
		return fmt.Errorf("DOODOCS_DAYS2_BACKEND_MAIL_PASSWORD is not set")
	}
	return nil
}

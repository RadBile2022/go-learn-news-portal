package infrastructure

import (
	"context"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func NewViper() {
	// Load the .env file using godotenv
	if err := godotenv.Load(); err != nil {
		slog.WarnContext(context.Background(), "No .env file found")
	}

	viper.SetConfigName(".env")   // Nama file tanpa ekstensi
	viper.SetConfigType("env")    // Format file
	viper.AddConfigPath("../../") // Path ke root proyek relatif dari cmd/api
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	viper.SetDefault("AES_KEY", "super-long-secret-key-12345678901")
	viper.SetDefault("JWT_KEY", "super-long-secret-key")
	viper.SetDefault("OTP_LENGTH", 6)

	viper.SetDefault("DISABLE_SINGLE_SESSION", false)

	viper.SetDefault("WEBHOOK_STRATEGY", "http")

	// Enable automatic environment variable lookup
	viper.AutomaticEnv()
}

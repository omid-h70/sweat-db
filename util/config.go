package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Env string `mapstructure:"ENVIRONMENT"`
	// DBDriver             string        `mapstructure:"DB_DRIVER"`
	// DBSource             string        `mapstructure:"DB_SOURCE"`
	// MigrationSrc         string        `mapstructure:"MIGRATION_FILE"`
	// RedisAddr            string        `mapstructure:"REDIS_ADDRESS"`
	// HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	// GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	// TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	// AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	// RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	// EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	// EmailSenderAddr      string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	// EmailSenderPass      string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml , whatever

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

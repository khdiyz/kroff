package config

import (
	"kroff/utils/logger"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DefaultLimit = "10"
	DefaultPage  = "1"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	HTTPHost string
	HTTPPort int

	Environment string
	Debug       bool

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	RedisHost     string
	RedisPort     int
	RedisPassword string
	RedisDB       int

	JWTSecret                string
	JWTAccessExpirationHours int
	JWTRefreshExpirationDays int

	HashKey string

	MinioEndpoint   string
	MinioAccessKey  string
	MinioSecretKey  string
	MinioUseSSL     bool
	MinioBucketName string
	MinioFileUrl    string

	AdminUsername string
	AdminPassword string
}

func GetConfig() *Config {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			logger.GetLogger().Error(err)
		}

		instance = &Config{
			HTTPHost:    cast.ToString(getOrReturnDefault("HTTP_HOST", "localhost")),
			HTTPPort:    cast.ToInt(getOrReturnDefault("HTTP_PORT", 4040)),
			Environment: cast.ToString(getOrReturnDefault("ENVIRONMENT", "development")),
			Debug:       cast.ToBool(getOrReturnDefault("DEBUG", true)),

			PostgresHost:     cast.ToString(getOrReturnDefault("POSTGRE_HOST", "localhost")),
			PostgresPort:     cast.ToInt(getOrReturnDefault("POSTGRE_PORT", 5432)),
			PostgresDatabase: cast.ToString(getOrReturnDefault("POSTGRE_DB", "kroff_db")),
			PostgresUser:     cast.ToString(getOrReturnDefault("POSTGRE_USER", "khdiyz")),
			PostgresPassword: cast.ToString(getOrReturnDefault("POSTGRE_PASSWORD", "password")),

			RedisHost:     cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost")),
			RedisPort:     cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379)),
			RedisPassword: cast.ToString(getOrReturnDefault("REDIS_PASSWORD", "")),
			RedisDB:       cast.ToInt(getOrReturnDefault("REDIS_DB", 0)),

			JWTSecret:                cast.ToString(getOrReturnDefault("JWT_SECRET", "kroff-forever-2025")),
			JWTAccessExpirationHours: cast.ToInt(getOrReturnDefault("JWT_ACCESS_EXPIRATION_HOURS", 12)),
			JWTRefreshExpirationDays: cast.ToInt(getOrReturnDefault("JWT_REFRESH_EXPIRATION_DAYS", 3)),

			HashKey: cast.ToString(getOrReturnDefault("HASH_KEY", "skd32r8wdahkkN2HSdqw")),

			MinioEndpoint:   cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", "")),
			MinioAccessKey:  cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY", "")),
			MinioSecretKey:  cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY", "")),
			MinioUseSSL:     cast.ToBool(getOrReturnDefault("MINIO_USE_SSL", "")),
			MinioBucketName: cast.ToString(getOrReturnDefault("MINIO_BUCKET_NAME", "")),
			MinioFileUrl:    cast.ToString(getOrReturnDefault("MINIO_FILE_URL", "")),

			AdminUsername: cast.ToString(getOrReturnDefault("ADMIN_USERNAME", "")),
			AdminPassword: cast.ToString(getOrReturnDefault("ADMIN_PASSWORD", "")),
		}
	})

	return instance
}

func getOrReturnDefault(key string, defaultValue any) any {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}

package config

import (
	"log"
	"os"
	"sync"
)

type Config struct {
	Port string

	MongoURI string
	MongoDB  string

	MinioEndpoint string
	MinioKey      string
	MinioSecret   string
	MinioUseSSL   bool
	MinioBaseURL  string

	RawBucket string
	HLSBucket string
}

var (
	cfg  *Config
	once sync.Once
)

func Load() *Config {
	once.Do(func() {
		cfg = &Config{
			Port:           os.Getenv("APP_PORT"),
			MongoURI:       os.Getenv("MONGO_URI"),
			MongoDB:        os.Getenv("MONGO_DB"),
			MinioEndpoint:  os.Getenv("MINIO_ENDPOINT"),
			MinioKey:       os.Getenv("MINIO_ACCESS_KEY"),
			MinioSecret:   os.Getenv("MINIO_SECRET_KEY"),
			MinioUseSSL:   os.Getenv("MINIO_USE_SSL") == "true",
			MinioBaseURL:  os.Getenv("MINIO_BASE_URL"),
			RawBucket:     os.Getenv("MINIO_BUCKET_RAW"),
			HLSBucket:     os.Getenv("MINIO_BUCKET_HLS"),
		}

		validate()
	})

	return cfg
}

func validate() {
	if cfg.Port == "" {
		log.Fatal("APP_PORT is required")
	}
	if cfg.RawBucket == "" || cfg.HLSBucket == "" {
		log.Fatal("MINIO_BUCKET_* is required")
	}
}

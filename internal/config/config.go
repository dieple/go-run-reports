package config

import (
	"os"
)

type Config struct {
	DynamoTable string
	Region      string
	Environment string
}

func Load() (*Config, error) {
	return &Config{
		DynamoTable: os.Getenv("DYNAMODB_TABLE_NAME"),
		Region:      os.Getenv("AWS_REGION"),
		Environment: os.Getenv("ENVIRONMENT"),
	}, nil
}
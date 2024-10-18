package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config holds the configuration values
type Config struct {
	MongoURI string
	Database string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		MongoURI: os.Getenv("MONGO_URI"),
		Database: os.Getenv("MONGO_DB"),
	}
}

// NewMongoClient initializes a new MongoDB client
func NewMongoClient(cfg *Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

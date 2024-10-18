package config

import (
	"context"

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
	// 환경 변수가 없을 경우 기본값 제공
	mongoURI := "mongodb://localhost:27017"
	database := "test_golang_mongo"

	return &Config{
		MongoURI: mongoURI,   // 환경 변수에서 URI를 가져옴
		Database: database,   // 환경 변수에서 DB 이름을 가져옴
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

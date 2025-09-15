package services

import (
	"GoATTHStart/internal/cache"
	"GoATTHStart/internal/database"
	"encoding/json"
)

type HealthService struct {
	db    database.DBService
	redis *cache.RedisClient
}

func NewHealthService(db database.DBService) *HealthService {
	return &HealthService{
		db: db,
		// redis: redis,
	}
}

type HealthStatus struct {
	Database string `json:"database"`
	// Redis    string `json:"redis"`
	// Websocket string            `json:"websocket"`
}

func (s *HealthService) CheckHealth() (HealthStatus, error) {
	dbHealth := s.db.Health()

	dbBytes, err := json.Marshal(dbHealth)
	if err != nil {
		return HealthStatus{}, err
	}

	// redisHealth := s.redis.Health()

	// redisBytes, err := json.Marshal(redisHealth)
	// if err != nil {
	// 	return HealthStatus{}, err
	// }

	return HealthStatus{
		Database: string(dbBytes),
		// Redis:    string(redisBytes),
	}, nil
}

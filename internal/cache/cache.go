package cache

import (
	"GoATTHStart/internal/config"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(cfg *config.CacheConfig) (*RedisClient, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Username: cfg.RedisUsername,
		Password: cfg.RedisPassword,
		DB:       0, // use default DB

		// Connection Pool Settings
		PoolSize:        50,        // Increased for concurrent timer operations
		MinIdleConns:    10,        // Keep more connections ready
		ConnMaxLifetime: time.Hour, // Refresh connections periodically

		// Timeouts
		DialTimeout:  5 * time.Second, // Reduced as it's local/internal service
		ReadTimeout:  2 * time.Second, // Reduced for faster failure detection
		WriteTimeout: 2 * time.Second, // Reduced for faster failure detection
		PoolTimeout:  4 * time.Second, // Should be greater than operation timeout

		// Retry Strategy
		MaxRetries:      3,
		MinRetryBackoff: 8 * time.Millisecond,
		MaxRetryBackoff: 500 * time.Millisecond,
		ConnMaxIdleTime: 5 * time.Minute,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &RedisClient{
		Client: client,
	}, nil
}

func (r *RedisClient) Health() map[string]string {
	stats := make(map[string]string)

	if r.Client == nil {
		stats["status"] = "down"
		stats["error"] = "redis client is nil"
		return stats
	}

	// Create context with timeout for health check
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Try to ping Redis
	if err := r.Client.Ping(ctx).Err(); err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("redis down: %v", err)
		return stats
	}

	// Get Redis info
	info, err := r.Client.Info(ctx).Result()
	if err != nil {
		stats["status"] = "degraded"
		stats["error"] = fmt.Sprintf("could not get redis info: %v", err)
		return stats
	}

	// Redis is up, get pool stats
	poolStats := r.Client.PoolStats()
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Add connection pool statistics
	stats["total_conns"] = strconv.FormatInt(int64(poolStats.TotalConns), 10)
	stats["idle_conns"] = strconv.FormatInt(int64(poolStats.IdleConns), 10)
	stats["stale_conns"] = strconv.FormatInt(int64(poolStats.StaleConns), 10)
	stats["hits"] = strconv.FormatInt(int64(poolStats.Hits), 10)
	stats["misses"] = strconv.FormatInt(int64(poolStats.Misses), 10)
	stats["timeouts"] = strconv.FormatInt(int64(poolStats.Timeouts), 10)

	// Parse Redis INFO command for additional metrics
	infoLines := strings.Split(info, "\r\n")
	for _, line := range infoLines {
		if strings.HasPrefix(line, "used_memory:") {
			stats["used_memory"] = strings.Split(line, ":")[1]
		}
		if strings.HasPrefix(line, "connected_clients:") {
			stats["connected_clients"] = strings.Split(line, ":")[1]
		}
		if strings.HasPrefix(line, "blocked_clients:") {
			stats["blocked_clients"] = strings.Split(line, ":")[1]
		}
		if strings.HasPrefix(line, "total_connections_received:") {
			stats["total_connections_received"] = strings.Split(line, ":")[1]
		}
	}

	// Evaluate stats to provide health messages
	if poolStats.TotalConns > 45 { // Near max pool size (50)
		stats["message"] = "Redis is experiencing high connection usage"
	}

	if poolStats.Misses > poolStats.Hits/2 {
		stats["message"] = "High number of pool misses, consider increasing pool size"
	}

	if poolStats.Timeouts > 100 {
		stats["message"] = "High number of timeouts, check Redis server load"
	}

	if poolStats.StaleConns > 0 {
		stats["message"] = "Detected stale connections, check network stability"
	}

	// Check memory usage if available
	if memStr, ok := stats["used_memory"]; ok {
		if mem, err := strconv.ParseInt(memStr, 10, 64); err == nil {
			// Alert if memory usage is above 1GB
			if mem > 1024*1024*1024 {
				stats["message"] = "High memory usage in Redis"
			}
		}
	}

	// Get key space info for timer metrics
	keySpace, err := r.Client.DBSize(ctx).Result()
	if err == nil {
		stats["total_keys"] = strconv.FormatInt(keySpace, 10)

		// Count timer keys specifically
		timerKeys, err := r.Client.Keys(ctx, "timer:*").Result()
		if err == nil {
			stats["timer_keys"] = strconv.Itoa(len(timerKeys))
		}
	}

	return stats
}

// Close closes the Redis connection
func (r *RedisClient) Close() error {
	if r.Client != nil {
		return r.Client.Close()
	}
	return nil
}

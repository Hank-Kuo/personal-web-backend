package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/allegro/bigcache"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/Hank-Kuo/personal-web-backend/config"
)

func ConnectDB(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	switch cfg.Adapter {
	case "sqlite3":
		return connectSqlite3(cfg)
	case "mysql":
		return connectMySQL(cfg)
	case "postgres":
		return connectPostgres(cfg)
	default:
		return nil, errors.New("Database: " + cfg.Adapter + " not exist")
	}
}

func connectSqlite3(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", cfg.Host)

}
func connectMySQL(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	setting := fmt.Sprintf("%s:%s@tcp(%s):%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Db)
	return sqlx.Open("mysql", setting)
}

func connectPostgres(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	setting := fmt.Sprintf("%s:%s@tcp(%s):%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Db)
	return sqlx.Open("mysql", setting)
}

func ConnectCacheDB(cfg *config.CacheConfig) (*bigcache.BigCache, error) {
	setting := bigcache.DefaultConfig(time.Duration(cfg.MaxLiftime) * time.Minute)
	return bigcache.NewBigCache(setting)
}

func ConnectRedis(ctx context.Context, cfg *config.RedisConfig) (*redis.Client, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
	})
	_, err := r.Ping(ctx).Result()
	return r, err
}

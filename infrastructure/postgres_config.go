package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host                  string
	Port                  string
	UserName              string
	Password              string
	DbName                string
	MaxConnections        string
	MaxConnectionIdleTime string
}

func GetConnectionPool(ctx context.Context, config Config) *pgxpool.Pool {
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable pool_max_conns=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Port,
		config.UserName,
		config.Password,
		config.DbName,
		config.MaxConnections,
		config.MaxConnectionIdleTime,
	)

	connConfig, parseConfigErr := pgxpool.ParseConfig(dbUrl)

	if parseConfigErr != nil {
		panic(parseConfigErr)
	}
	withConfig, connError := pgxpool.ConnectConfig(ctx, connConfig)

	if connError != nil {
		panic(connError)
	}
	return withConfig
}

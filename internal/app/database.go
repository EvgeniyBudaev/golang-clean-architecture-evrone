package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/config"
)

func newPostgresConnection(cfg *config.Config) (*sql.DB, error) {
	databaseURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSlMode,
	)

	return sql.Open("postgres", databaseURL)
}

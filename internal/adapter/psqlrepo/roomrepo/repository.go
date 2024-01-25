package roomrepo

import (
	"database/sql"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
)

type Repository struct {
	logger logger.Logger
	conn   *sql.DB
}

func NewRepository(logger logger.Logger, conn *sql.DB) *Repository {
	return &Repository{
		logger: logger,
		conn:   conn,
	}
}

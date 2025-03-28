package state

import (
	"github.com/sonnq2010/blog-aggregator/internal/config"
	"github.com/sonnq2010/blog-aggregator/internal/database"
)

type State struct {
	DB     *database.Queries
	Config *config.Config
}

package state

import (
	"github.com/mecebeci/blog-aggregator/internal/config"
	"github.com/mecebeci/blog-aggregator/internal/database"
)

type State struct {
	Config *config.Config
	DB     *database.Queries
}

package handler

import (
	"github.com/italopatrick/api-opportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializedHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}

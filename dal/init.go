package dal

import (
	"github.com/JackTJC/gmFS_backend/dal/cache"
	"github.com/JackTJC/gmFS_backend/dal/db"
)

func InitDB() {
	cache.InitCache()
	db.InitDB()
}

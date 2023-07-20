package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/Jenietor2/app-inventory/settings"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", s.DB.User, s.DB.Password, s.DB.Host, s.DB.Port, s.DB.Name)
	//connectionString := "root:rootroot@tcp(localhost:3333)/app_inventory?parseTime=true"
	return sqlx.ConnectContext(ctx, "mysql", connectionString)
}

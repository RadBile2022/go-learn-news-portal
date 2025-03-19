package infrastructure

import (
	"database/sql"
	"go-learn-news-portal/database/seeds"
	"log"
	"log/slog"

	_ "github.com/lib/pq"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Connection string
}

type Postgres interface {
	GetClient() *sql.DB
	GetClientGorm() *gorm.DB
	Close()
}

type postgresClient struct {
	client *sql.DB
	gorm   *gorm.DB
}

func (p postgresClient) GetClient() *sql.DB {
	return p.client
}

func (p postgresClient) GetClientGorm() *gorm.DB {
	return p.gorm
}

func (p postgresClient) Close() {
	p.client.Close()
}

func NewPostgres(connection string) Postgres {
	slog.Info("Connecting to PostgreSQL...")

	// Initialize sql.DB client
	conn, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	// Test the connection
	if err = conn.Ping(); err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	// Initialize GORM client
	gormDB, err := gorm.Open(gormPostgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to initialize GORM for PostgreSQL:", err)
	}

	log.Println("Connection to PostgreSQL established successfully...")

	seeds.SeedRoles(gormDB)

	//db = gormDB

	// Return the client
	return postgresClient{client: conn, gorm: gormDB}
}

//func NewDBTrx(ctx context.Context) *gorm.DB {
//	return db.WithContext(ctx).Begin()
//}

package main

import (
	"go-learn-news-portal/infrastructure"
	"go-learn-news-portal/internal/core/entity"
	"log"

	"github.com/spf13/viper"
)

func main() {
	infrastructure.NewViper()
	postgres := infrastructure.NewPostgres(
		viper.GetString("POSTGRES_CONNECTION"),
		//&gorm.Config{
		//	DisableForeignKeyConstraintWhenMigrating: false,
		//	IgnoreRelationshipsWhenMigrating:         false,
		//},
	)
	db := postgres.GetClientGorm()

	err := db.AutoMigrate(
		&entity.File{},
	)
	if err != nil {
		log.Fatalf("error %v", err)
		return
	}
	log.Printf("success")
}

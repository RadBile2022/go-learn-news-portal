package main

import (
	"github.com/RadBile2022/go-learn-news-portal/infrastructure"
	"github.com/RadBile2022/go-learn-news-portal/internal/core/entity"
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

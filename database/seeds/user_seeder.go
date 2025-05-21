package seeds

import (
	"github.com/rs/zerolog/log"
	"github.com/RadBile2022/go-learn-news-portal/internal/core/entity"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/convert"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	bytes, err := convert.HashPassword("12345678")
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating password hash")
	}

	admin := entity.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: string(bytes),
	}

	if err := db.FirstOrCreate(&admin, entity.User{Email: "admin@gmail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("Error seeding admin role")
	} else {
		log.Info().Msg("Admin role seeded successfully")
	}
}

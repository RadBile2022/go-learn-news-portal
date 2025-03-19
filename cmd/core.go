package cmd

import (
	radstore "github.com/RadBile2022/go-library-radstore"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"go-learn-news-portal/infrastructure"
	"go-learn-news-portal/internal/core/service"
	"go-learn-news-portal/internal/framework/secondary/repository"
	"time"
)

type CoreOptions struct {
	userCore     service.User
	authCore     service.Auth
	categoryCore service.Category
	contentCore  service.Content
}

func initCore() *CoreOptions {
	infrastructure.NewViper()

	infrastructure.NewSentry(sentry.ClientOptions{
		Dsn:              viper.GetString("SENTRY_DSN"),
		EnableTracing:    viper.GetBool("SENTRY_ENABLE_TRACING"),
		TracesSampleRate: viper.GetFloat64("SENTRY_TRACES_SAMPLE_RATE"),
	})
	defer sentry.Flush(2 * time.Second)

	db := infrastructure.NewPostgres(viper.GetString("POSTGRES_CONNECTION"))
	defer db.Close()

	// storage
	r2Storage := radstore.NewCloudflareR2Adapter(&radstore.CloudflareR2Options{
		Endpoint:      viper.GetString("CLOUDFLARE_ENDPOINT"),
		Region:        viper.GetString("CLOUDFLARE_REGION"),
		AccessKey:     viper.GetString("CLOUDFLARE_ACCESS_KEY"),
		SecretKey:     viper.GetString("CLOUDFLARE_SECRET_KEY"),
		Token:         viper.GetString("CLOUDFLARE_TOKEN"),
		BucketName:    viper.GetString("CLOUDFLARE_BUCKET_NAME"),
		UserTagsKey:   viper.GetString("CLOUDFLARE_USER_TAGS_KEY"),
		UserTagsValue: viper.GetString("CLOUDFLARE_USER_TAGS_VALUE"),
		AesKey:        viper.GetString("CLOUDFLARE_AES_KEY"),
		UseAesKey:     viper.GetBool("CLOUDFLARE_USE_AES_KEY"),
		PublicUrl:     viper.GetString("CLOUDFLARE_PUBLIC_URL"),
	})

	// repository
	userRepo := repository.NewUser(db.GetClientGorm())
	categoryRepo := repository.NewCategory(db.GetClientGorm())
	contentRepo := repository.NewContent(db.GetClientGorm())

	// core
	userCore := service.NewUser(userRepo)
	authCore := service.NewAuth(userRepo)
	categoryCore := service.NewCategory(categoryRepo)
	contentCore := service.NewContent(contentRepo, r2Storage)

	return &CoreOptions{
		userCore:     userCore,
		authCore:     authCore,
		categoryCore: categoryCore,
		contentCore:  contentCore,
	}
}

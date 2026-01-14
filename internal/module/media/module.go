package media

import (
	"chillhub/internal/module/media/handler"
	"chillhub/internal/module/media/repository"
	"chillhub/internal/module/media/service"
	minioshared "chillhub/internal/shared/minio"
	"log"
	"os"
)




type Module struct {
	Handler *handler.MediaHandler
}

func NewModule(repo repository.MediaRepository, minio *minioshared.Util) *Module {
	// Module chứa handler để đăng ký route
	var rawBucket = os.Getenv("BUCKET_MEDIA") // mỗi module quản lý bucket riêng

	log.Printf("Media Module sử dụng bucket: %s", rawBucket)

	transcoder := service.NewTranscodingService(minio, repo)


	// Service nhận minio và bucket riêng
	svc := service.NewMediaService(repo, transcoder, minio, rawBucket)

	// Handler nhận service
	h := handler.NewMediaHandler(svc)

	return &Module{Handler: h}
}

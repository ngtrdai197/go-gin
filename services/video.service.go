package services

import (
	"github.com/ngtrdai197/go-gin/entities"
)

type VideoService interface {
	Save(entities.Video) entities.Video
	FindAll() []entities.Video
}

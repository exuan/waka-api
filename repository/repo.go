package repository

import (
	"github.com/exuan/waka-api/validator"
	"gorm.io/gorm"
)

type (
	Repository interface {
		BatchHeartbeat(vhs *[]*validator.Heartbeat) (int64, error)
	}

	repository struct {
		db *gorm.DB
	}
)

func New(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

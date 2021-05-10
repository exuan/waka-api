package service

import (
	"github.com/exuan/waka-api/repository"
	"github.com/exuan/waka-api/validator"
)

type (
	Service interface {
		BatchHeartbeat(vhs *[]*validator.Heartbeat) (int64, error)
	}

	service struct {
		r repository.Repository
	}
)

func New(r repository.Repository) Service {
	return &service{r}
}

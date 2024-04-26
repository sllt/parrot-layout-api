package service

import (
	"github.com/sllt/parrot-layout-api/internal/repository"
	"github.com/sllt/parrot-layout-api/pkg/jwt"
	"github.com/sllt/parrot-layout-api/pkg/log"
	"github.com/sllt/parrot-layout-api/pkg/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(tm repository.Transaction, logger *log.Logger, sid *sid.Sid, jwt *jwt.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}

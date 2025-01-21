package service

import (
	pb "lottery/api/lottery/v1"
	"lottery/internal/usecase"
)

type LotteryService struct {
	pb.UnimplementedLotteryServer
	lotteryCase *usecase.LotteryCase
	limitCase   *usecase.LimitCase
	adminCase   *usecase.AdminCase
}

func NewLotteryService(loc *usecase.LotteryCase, lic *usecase.LimitCase, ac *usecase.AdminCase) *LotteryService {
	return &LotteryService{
		lotteryCase: loc,
		limitCase:   lic,
		adminCase:   ac,
	}
}

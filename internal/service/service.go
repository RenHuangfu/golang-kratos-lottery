package service

import (
	"github.com/google/wire"
	pb "lottery/api/lottery/v1"
	"lottery/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewLotteryService, NewAdminService)

type LotteryService struct {
	pb.UnimplementedLotteryServer
	lotteryCase *biz.LotteryCase
	limitCase   *biz.LimitCase
	adminCase   *biz.AdminCase
}

func NewLotteryService(loc *biz.LotteryCase, lic *biz.LimitCase, ac *biz.AdminCase) *LotteryService {
	return &LotteryService{
		lotteryCase: loc,
		limitCase:   lic,
		adminCase:   ac,
	}
}

// AdminService 奖品管理后台
type AdminService struct {
	adminCase *biz.AdminCase
}

func NewAdminService(ac *biz.AdminCase) *AdminService {
	return &AdminService{
		adminCase: ac,
	}
}

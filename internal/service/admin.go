package service

import (
	"context"
	"fmt"
	"github.com/BitofferHub/pkg/middlewares/log"
	pb "lottery/api/lottery/v1"
	"lottery/common/entity"
	"lottery/internal/usecase"
	"time"
)

// AdminService 奖品管理后台
type AdminService struct {
	pb.AdminHTTPServer
	adminCase *usecase.AdminCase
}

func NewAdminService(ac *usecase.AdminCase) *AdminService {
	return &AdminService{
		adminCase: ac,
	}
}

// AddPrize 添加奖品
func (a *AdminService) AddPrize(ctx context.Context, req *pb.AddPrizeRequest) (*pb.AddPrizeResponse, error) {
	if err := a.adminCase.AddPrize(ctx, &entity.ViewPrize{
		Id:           uint(req.Prize.Id),
		Title:        req.Prize.Title,
		Img:          req.Prize.Img,
		PrizeNum:     int(req.Prize.PrizeNum),
		PrizeCode:    req.Prize.PrizeCode,
		PrizeTime:    uint(req.Prize.PrizeTime),
		LeftNum:      int(req.Prize.LeftNum),
		PrizeType:    uint(req.Prize.PrizeType),
		PrizePlan:    req.Prize.PrizePlan,
		BeginTime:    time.Unix(req.Prize.BeginTime, 0),
		EndTime:      time.Unix(req.Prize.EndTime, 0),
		DisplayOrder: uint(req.Prize.DisplayOrder),
		SysStatus:    uint(req.Prize.SysStatus),
	}); err != nil {
		log.ErrorContextf(ctx, "adminService|AddPrize err:%v", err)
		return nil, fmt.Errorf("adminService|AddPrize:%v", err)
	}
	return &pb.AddPrizeResponse{}, nil
}

// AddPrizeList 添加奖品列表
func (a *AdminService) AddPrizeList(ctx context.Context, req *pb.AddPrizeListRequest) (*pb.AddPrizeListResponse, error) {
	var viewps []*entity.ViewPrize
	for _, prize := range req.Prizes {
		viewps = append(viewps, &entity.ViewPrize{
			Id:           uint(prize.Id),
			Title:        prize.Title,
			Img:          prize.Img,
			PrizeNum:     int(prize.PrizeNum),
			PrizeCode:    prize.PrizeCode,
			PrizeTime:    uint(prize.PrizeTime),
			LeftNum:      int(prize.LeftNum),
			PrizeType:    uint(prize.PrizeType),
			PrizePlan:    prize.PrizePlan,
			BeginTime:    time.Unix(prize.BeginTime, 0),
			EndTime:      time.Unix(prize.EndTime, 0),
			DisplayOrder: uint(prize.DisplayOrder),
			SysStatus:    uint(prize.SysStatus),
		})
	}
	if err := a.adminCase.AddPrizeList(ctx, viewps); err != nil {
		log.ErrorContextf(ctx, "adminService|AddPrizeList err:%v", err)
		return nil, fmt.Errorf("adminService|AddPrizeList:%v", err)
	}
	return &pb.AddPrizeListResponse{}, nil
}

// ClearPrize 清空奖品
func (a *AdminService) ClearPrize(ctx context.Context, req *pb.ClearPrizeRequest) (*pb.ClearPrizeResponse, error) {
	if err := a.adminCase.ClearPrize(ctx); err != nil {
		log.ErrorContextf(ctx, "adminService|ClearPrize err:%v", err)
		return nil, fmt.Errorf("adminService|ClearPrize:%v", err)
	}
	return &pb.ClearPrizeResponse{}, nil
}

func (a *AdminService) ImportCoupon(ctx context.Context, req *pb.ImportCouponRequest) (*pb.ImportCouponResponse, error) {
	successNum, failNum, err := a.adminCase.ImportCoupon(ctx, uint(req.PrizeId), req.Codes)
	if err != nil {
		return nil, fmt.Errorf("AdminService|ImportCoupon|%v", err)
	}
	log.Infof("ImportCoupon|successNum=%d|failNum=%d\n", successNum, failNum)
	return &pb.ImportCouponResponse{}, nil
}

func (a *AdminService) ClearCoupon(ctx context.Context, req *pb.ClearCouponRequest) (*pb.ClearCouponResponse, error) {
	if err := a.adminCase.ClearCoupon(ctx); err != nil {
		log.ErrorContextf(ctx, "adminService|ClearCoupon err:%v", err)
		return nil, fmt.Errorf("adminService|ClearCoupon:%v", err)
	}
	return &pb.ClearCouponResponse{}, nil
}

func (a *AdminService) ImportCouponWithCache(ctx context.Context, req *pb.ImportCouponWithCacheRequest) (*pb.ImportCouponWithCacheResponse, error) {
	successNum, failNum, err := a.adminCase.ImportCouponWithCache(ctx, uint(req.PrizeId), req.Codes)
	if err != nil {
		return nil, fmt.Errorf("AdminService|ImportCouponWithCache|%v", err)
	}
	log.Infof("ImportCouponWithCache|successNum=%d|failNum=%d\n", successNum, failNum)
	return &pb.ImportCouponWithCacheResponse{}, nil
}

func (a *AdminService) ClearLotteryTimes(ctx context.Context, req *pb.ClearLotteryTimesRequest) (*pb.ClearLotteryTimesResponse, error) {
	if err := a.adminCase.ClearLotteryTimes(ctx); err != nil {
		log.ErrorContextf(ctx, "adminService|ClearCoupon err:%v", err)
		return nil, fmt.Errorf("adminService|ClearCoupon:%v", err)
	}
	return &pb.ClearLotteryTimesResponse{}, nil
}

func (a *AdminService) ClearResult(ctx context.Context, req *pb.ClearResultRequest) (*pb.ClearResultResponse, error) {
	if err := a.adminCase.ClearResult(ctx); err != nil {
		log.ErrorContextf(ctx, "adminService|ClearCoupon err:%v", err)
		return nil, fmt.Errorf("adminService|ClearCoupon:%v", err)
	}
	return &pb.ClearResultResponse{}, nil
}

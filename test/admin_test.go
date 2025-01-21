package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	pb "lottery/api/lottery/v1"
	"lottery/common/constant"
	"strconv"
	"strings"
	"testing"
	"time"
)

func RestyPost(url string, req interface{}, resp interface{}) error {
	client := resty.New()

	bytesData, err := json.Marshal(req)
	if err != nil {
		return err
	}
	respBody, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(bytes.NewReader(bytesData)).
		SetResult(resp).
		Post(url)
	fmt.Println("respBody = ", respBody)
	return nil
}

func TestAddPrizeList(t *testing.T) {
	addPrize1 := pb.Prize{
		Title:     "iphone",
		Img:       "https://p0.ssl.qhmsg.com/t016ff98b934914aca6.png",
		PrizeNum:  10,
		PrizeCode: "1-10",
		EndTime:   time.Now().Add(time.Hour * 24 * 7).Unix(),
		BeginTime: time.Now().Unix(),
		PrizeType: constant.PrizeTypeEntityLarge,
	}

	addPrize2 := pb.Prize{
		Title:     "homepod",
		Img:       "https://imgservice.suning.cn/uimg1/b2c/image/t_QerWgoH9ergm0_NY4WhA.png_800w_800h_4e",
		PrizeNum:  50,
		PrizeCode: "100-150",
		EndTime:   time.Now().Add(time.Hour * 24 * 7).Unix(),
		BeginTime: time.Now().Unix(),
		PrizeType: constant.PrizeTypeEntityMiddle,
	}

	addPrize3 := pb.Prize{
		Title:     "充电器",
		Img:       "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS5Y7iNXpcthgJ6yvE3Os1bTwLARVnvwYXeKA&usqp=CAU",
		PrizeNum:  500,
		PrizeCode: "1500-2000",
		EndTime:   time.Now().Add(time.Hour * 24 * 7).Unix(),
		BeginTime: time.Now().Unix(),
		PrizeType: constant.PrizeTypeEntitySmall,
	}

	addPrize4 := pb.Prize{
		Title:     "优惠券",
		Img:       "https://static.699pic.com/images/diversion/d66d647c52cd66beb800ba09748ea080.jpgU",
		PrizeNum:  8000,
		PrizeCode: "2001-9999",
		EndTime:   time.Now().Add(time.Hour * 24 * 7).Unix(),
		BeginTime: time.Now().Unix(),
		PrizeType: constant.PrizeTypeCouponDiff,
	}
	prizeList := []*pb.Prize{
		&addPrize1, &addPrize2, &addPrize3, &addPrize4,
	}

	addPrizeListReq := pb.AddPrizeListRequest{
		Prizes: prizeList,
	}
	if err := RestyPost("http://localhost:10080/admin/add_prize_list", &addPrizeListReq, &pb.AddPrizeListResponse{}); err != nil {
		fmt.Errorf("add prize list err:%v", err)
	}
}

func TestClearPrize(t *testing.T) {
	clearReq := pb.ClearPrizeRequest{}
	if err := RestyPost("http://localhost:10080/admin/clear_prize", &clearReq, &pb.ClearPrizeResponse{}); err != nil {
		fmt.Errorf("clear prize err:%v", err)
	}
}

func TestImportCoupon(t *testing.T) {
	codeStr := ""
	for i := 1; i <= 50; i++ {
		code := "coupon_code00000"
		if i < 10 {
			code = code + "0" + strconv.Itoa(i) + "\n"
		} else {
			code = code + strconv.Itoa(i) + "\n"
		}
		codeStr = codeStr + code
	}
	couponCode := strings.Trim(codeStr, "\n")
	couponInfo := pb.ImportCouponRequest{
		PrizeId: 40,
		Codes:   couponCode,
	}
	if err := RestyPost("http://localhost:10080/admin/import_coupon", &couponInfo, &pb.ImportCouponResponse{}); err != nil {
		fmt.Errorf("import coupon err:%v", err)
	}
}

func TestImportCouponWithCache(t *testing.T) {
	codeStr := ""
	for i := 1; i <= 50; i++ {
		code := "coupon_code00000"
		if i < 10 {
			code = code + "0" + strconv.Itoa(i) + "\n"
		} else {
			code = code + strconv.Itoa(i) + "\n"
		}
		codeStr = codeStr + code
	}
	couponCode := strings.Trim(codeStr, "\n")
	couponInfo := pb.ImportCouponWithCacheRequest{
		PrizeId: 8,
		Codes:   couponCode,
	}
	if err := RestyPost("http://localhost:10080/admin/import_coupon_cache", &couponInfo, &pb.ImportCouponWithCacheResponse{}); err != nil {
		fmt.Errorf("import coupon err:%v", err)
	}
}

func TestClearCoupon(t *testing.T) {
	clearReq := pb.ClearCouponRequest{}
	if err := RestyPost("http://localhost:10080/admin/clear_coupon", &clearReq, &pb.ClearCouponResponse{}); err != nil {
		fmt.Errorf("clear coupon err:%v", err)
	}
}

func TestClearLotteryTimes(t *testing.T) {
	clearReq := pb.ClearLotteryTimesRequest{}
	if err := RestyPost("http://localhost:10080/admin/clear_lottery_times", &clearReq, &pb.ClearLotteryTimesResponse{}); err != nil {
		fmt.Errorf("clear lottery times err:%v", err)
	}
}

func TestClearResult(t *testing.T) {
	clearReq := pb.ClearResultRequest{}
	if err := RestyPost("http://localhost:10080/admin/clear_result", &clearReq, &pb.ClearResultResponse{}); err != nil {
		fmt.Errorf("clear result err:%v", err)
	}
}

func TestLotteryV1(t *testing.T) {
	lotteryReq := pb.LotteryReq{
		UserName: "zhangsan",
		Ip:       "192.168.9.9",
	}
	if err := RestyPost("http://localhost:10080/v1/lottery", &lotteryReq, &pb.LotteryRsp{}); err != nil {
		fmt.Errorf("LotteryV1 err:%v", err)
	}
}

func TestLotteryV2(t *testing.T) {
	lotteryReq := pb.LotteryReq{
		UserName: "lisi",
		Ip:       "192.168.9.10",
	}
	if err := RestyPost("http://localhost:10080/v2/lottery", &lotteryReq, &pb.LotteryRsp{}); err != nil {
		fmt.Errorf("LotteryV1 err:%v", err)
	}
}

func TestLotteryV3(t *testing.T) {
	lotteryReq := pb.LotteryReq{
		UserName: "lisi",
		Ip:       "192.168.9.10",
	}
	if err := RestyPost("http://localhost:10080/v3/lottery", &lotteryReq, &pb.LotteryRsp{}); err != nil {
		fmt.Errorf("LotteryV1 err:%v", err)
	}
}

func TestClear(t *testing.T) {
	TestClearPrize(t)
	TestClearCoupon(t)
	TestClearLotteryTimes(t)
	TestClearResult(t)
}

func TestImport(t *testing.T) {
	TestAddPrizeList(t)
	TestImportCoupon(t)
}

func TestImportWithCache(t *testing.T) {
	TestAddPrizeList(t)
	TestImportCouponWithCache(t)
}

package go_zpay

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ZPayInitParams{MerchantInfo{MERCHANT_CODE, MERCHANT_KEY}, DEPOSIT_URL, DEPOSIT_RESPONSE_URL, WITHDRAW_URL, WITHDRAW_RESPONSE_URL})

	//发请求
	resp, err := cli.WithdrawReq(GenWithdrawRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

// MYR
// INR
// IDR_NATIVE_FOREX
// THB_FOREX（暂无）
// VND

func GenWithdrawRequestDemo() ZPayWithdrawReq {
	// THB时additionalParams必填
	return ZPayWithdrawReq{
		Currency:      "MYR",
		MerchantRefNo: "2026_myyr_wd_01",
		Amount:        "1000.00",
		BankName:      "ABBANK",
		IfscCode:      "IBK",
		AccountNumber: "1234567890",
		AccountName:   "zf",
		AdditionalParams: AdditionalParamsObj{
			CustomerUserId:   "484799",
			CustomerUsername: "user_123",
		},
	}
}
